package controllers

import (
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/satori/go.uuid"
	"encoding/json"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
	"CloudManagerSystem/models/k8s"
	"k8s.io/api/extensions/v1beta1"
)

type KubePublishServiceController struct {
	BaseController
}

//分页展示NameSpace
func (c *KubePublishServiceController) DataGrid() {
	KubePublishServiceParam := models.KubePublishServiceQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &KubePublishServiceParam)

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)

	if (len(envUserCluster.ClusterId) == 0) {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
		return
	} else {
		KubePublishServiceParam.ClusterId = envUserCluster.ClusterId
	}

	data, total := models.KubePublishServicePageList(&KubePublishServiceParam)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//获取一个
func (c *KubePublishServiceController) Get() {
	strs := c.GetString("id")
	if (len(strs) == 0) {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("id 不能为空"), 0)
	}
	c.Ctx.Input.RequestBody = []byte("{\"Id\":\"" + strs + "\"}")
	c.DataGrid()
}

//添加
func (c *KubePublishServiceController) Post() {
	c.Save()
}

//修改
func (c *KubePublishServiceController) Put() {
	c.Save()
}

//删除
func (c *KubePublishServiceController) Delete() {
	strs := c.GetString("ids")
	ids := make([]string, 0, len(strs))
	for _, id := range strings.Split(strs, ",") {
		ids = append(ids, id)
	}
	if num, err := models.KubePublishServiceDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

//保存方法 (实现 添加与修改功能)
func (c *KubePublishServiceController) Save() {
	var err error
	m := models.KubePublishService{}
	/*if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
		return
	}*/
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", "")
	}

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)
	if len(envUserCluster.ClusterId) == 0 {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
		return
	}

	o := orm.NewOrm()
	var title string
	m.ClusterId = envUserCluster.ClusterId
	m.CreateUser = u.Id
	if strings.Compare(m.Id, "") == 0 {
		u4, _ := uuid.NewV4()
		m.Id = u4.String()
		title = "添加"
		_, err = o.Insert(&m)
	} else {
		title = "编辑"
		_, err = o.Update(&m)
	}

	kubeService := &models.KubeService{}
	o.QueryTable(models.KubeServiceTBName()).Filter("id", m.ServiceId).One(kubeService)

	kubeNamespace := &models.KubeNameSpace{}
	o.QueryTable(models.KubeNameSpaceTBName()).Filter("id", m.NamespaceId).One(kubeNamespace)
	//query := orm.NewOrm().QueryTable(KubeServicePortTBName())
	//query.Filter("service_id", ServiceId).All(&data)
	//serviceType := ""

	servic := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: m.Name, // "healthcheck-daolin-service",
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{"app": kubeService.Name}, //"healthcheckweb-deployment-daolin"},
			Ports: []apiv1.ServicePort{
			//{
			//	Protocol:apiv1.ProtocolTCP,
			//	Port:3004,
			//	TargetPort:intstr.FromInt(3040),
			//	NodePort:31182,
			//
			//},
			},
			//Type:,
		},
	}
	if kubeService.HostIp {
		servic.Spec.Type = apiv1.ServiceTypeNodePort
	} else {
		servic.Spec.Type = apiv1.ServiceTypeClusterIP
	}

	//Ingress
	paths := []v1beta1.HTTPIngressPath{}

	//删除M.ID记录
	models.KubePublishServicePathByIngressId(m.Id)
	//循环IngressPath
	for _, v := range m.Paths {

		v.PserviceId = m.Id
		u4, _ := uuid.NewV4()
		v.Id = u4.String()
		o.Insert(v)

		//获取ServicePort
		kubeServicePort := models.KubeServicePort{}
		orm.NewOrm().QueryTable(models.KubeServicePortTBName()).Filter("id", v.PortId).One(&kubeServicePort)

		if strings.ToLower(m.Stype) == "tcp" {

			//query.Filter("service_id", ServiceId).All(&data)
			getPort, _ := strconv.ParseInt(kubeServicePort.ServicePort, 10, 32)
			getTargetPort, _ := strconv.ParseInt(kubeServicePort.ContainerPort, 10, 32)

			servicePort := apiv1.ServicePort{
				Protocol:   apiv1.ProtocolTCP,
				Port:       int32(getPort),
				TargetPort: intstr.FromInt(int(getTargetPort)),
				//NodePort:nodePort,
				Name: kubeServicePort.Name,
			}

			if len(v.HostPort) > 0 {
				nodePort, _ := strconv.ParseInt(v.HostPort, 10, 32)
				servicePort.NodePort = int32(nodePort)
			}

			servic.Spec.Ports = append(servic.Spec.Ports, servicePort)
		} else {

			o.QueryTable(models.KubeServiceTBName()).Filter("id", v.ServiceId).One(kubeService)

			//创建Ingress path
			path := v1beta1.HTTPIngressPath{
				Path: v.Path,
				Backend: v1beta1.IngressBackend{
					ServiceName: kubeService.Name,
					ServicePort: intstr.FromString(kubeServicePort.Name),
				},
			}
			paths = append(paths, path)

		}
	}

	//k8s.ServiceCreateChannel
	//k8s.GetServiceCreateChannel(envUserCluster.ClusterId,)

	if strings.ToLower(m.Stype) == "tcp" {
		paras := []*apiv1.Service{servic}
		//调用k8s Clinet-go 创建service
		_ , err = k8s.CreateService(envUserCluster.ClusterId, kubeNamespace.Name, paras)

		s, _ := json.Marshal(servic)
		fmt.Println(string(s))
		//fmt.Println(result)
	} else {
		// 调用 k8s Ingress 创建方法
		ingress := &v1beta1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name: m.Name,
			},
			Spec: v1beta1.IngressSpec{
				Rules: []v1beta1.IngressRule{
					{
						Host: m.DomainName,
						IngressRuleValue: v1beta1.IngressRuleValue{
							HTTP: &v1beta1.HTTPIngressRuleValue{
								Paths: []v1beta1.HTTPIngressPath{
									//{
									//	Path: "/",
									//	Backend: v1beta1.IngressBackend{
									//		ServiceName: kubeService.Name,
									//		ServicePort: intstr.FromInt(3004),
									//	},
									//},
								},
							},
						},
					},
				},
			},
		}

		ingress.Spec.Rules[0].IngressRuleValue.HTTP.Paths =paths

		_ , err =k8s.CreateIngress(envUserCluster.ClusterId, kubeNamespace.Name, ingress)

	}

	if err == nil {
		c.jsonResult(enums.JRCodeSucc, title+"成功", m.Id)
	} else {
		fmt.Println(err.Error())
		c.jsonResult(enums.JRCodeFailed, title+"失败", err.Error())
	}
}
