package controllers

import (
	"encoding/json"
	//"fmt"
	"errors"
	"strings"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"fmt"
	"CloudManagerSystem/models/k8s"
)

type KubeServiceController struct {
	BaseController
}

func (c *KubeServiceController) DataGrid() {
	kubeservice := models.KubeServiceQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeservice)
	var ClusterId string
	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)
	//fmt.Println(u.UserName)
	if strings.Compare(envUserCluster.ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
	} else {
		ClusterId = envUserCluster.ClusterId
	}

	data, total := models.KubeServicePageList(&kubeservice, ClusterId)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func kubeServicePortSave(serviceport []*models.KubeServicePort, ServiceId string) error {
	var err error
	servicePort := models.KubeServicePort{}
	o := orm.NewOrm()

	for i, _ := range serviceport {
		servicePort = *(serviceport[i])
		//fmt.Println("--------------------",servicePort)

		if strings.Compare(servicePort.Id, "") == 0 {
			uu, _ := uuid.NewV4()
			servicePort.Id = uu.String()
			servicePort.ServiceId = ServiceId
			_, err = o.Insert(&servicePort)
			if err != nil {
				err = errors.New("添加" + servicePort.Name + "失败")
				return err
			}
		} else {
			_, err = o.Update(&servicePort)
			if err != nil {
				err = errors.New("编辑" + servicePort.Name + "失败")
				return err
			}
		}

	}
	return nil
}

func (c *KubeServiceController) Save() {
	var err error
	var title string

	kubeservice := models.KubeService{}
	//serviceport := models.KubeServicePort{}

	jsonkubeservice := models.JSONKubeService{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &jsonkubeservice)

	var ClusterId string

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)

	if strings.Compare(envUserCluster.ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
	} else {
		ClusterId = envUserCluster.ClusterId
	}

	//fmt.Println("-------", jsonkubeservice)
	o := orm.NewOrm()
	kubeservice = *(jsonkubeservice.Service)
	//fmt.Println("---kubeservice----", kubeservice)

	if strings.Compare(kubeservice.Id, "") == 0 {
		title = "添加"
		var flag bool

		flag, err = models.KubeServiceIsExist(kubeservice.Name)

		if !flag {
			c.jsonResult(enums.JRCodeFailed, title+"失败,"+err.Error(), err)
		}
		u4, _ := uuid.NewV4()
		kubeservice.Id = u4.String()
		kubeservice.ClusterId = ClusterId
		kubeservice.CreateUser = u.UserName
		err = kubeServicePortSave(jsonkubeservice.ServicePort, kubeservice.Id)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, err.Error(), kubeservice.Id)
		}

		_, err = o.Insert(&kubeservice)
	} else {
		title = "编辑"
		err = kubeServicePortSave(jsonkubeservice.ServicePort, kubeservice.Id)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, err.Error(), kubeservice.Id)
		}
		kubeservice.ClusterId = ClusterId

		aa := &models.KubeService{Id: kubeservice.Id}
		o.Read(aa)
		//not update
		kubeservice.CreateUser = aa.CreateUser
		kubeservice.CreateTime = aa.CreateTime

		_, err = o.Update(&kubeservice)

	}
	if err == nil {
		c.jsonResult(enums.JRCodeSucc, title+kubeservice.Name+"成功", kubeservice.Id)
	} else {
		c.jsonResult(enums.JRCodeFailed, title+kubeservice.Name+"失败", kubeservice.Id)
	}

}

func (c *KubeServiceController) Delete() {
	strs := c.GetString("ids")
	ids := strings.Split(strs, ",")

	if num, err := models.KubeServiceDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, err.Error(), 0)
	}

}

func (c *KubeServiceController) PublishORRollback() {
	prkubeservice := models.KubeServicePubORRollback{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &prkubeservice)
	var title string
	var ClusterId string

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)

	if strings.Compare(envUserCluster.ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
	} else {
		ClusterId = envUserCluster.ClusterId
	}

	if strings.Compare(prkubeservice.VersionId, "") == 0 || strings.Compare(prkubeservice.ServiceId, "") == 0 {
		title = "pub"
		c.jsonResult(enums.JRCodeFailed, title+"失败", prkubeservice)
	}

	clienthandle, err := models.GetApiServerHandle(ClusterId, false)

	if err != nil {
		fmt.Println(err)
		c.jsonResult(enums.JRCodeFailed, err.Error(), "")
	}

	if strings.Compare(prkubeservice.Type, "17") == 0 {
		//appDeploymentSpec := new(k8s.AppDeploymentSpec)
		err = k8s.DeployApp(clienthandle, &prkubeservice, ClusterId, true)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "pub失败"+err.Error(), "")
		}
		c.jsonResult(enums.JRCodeSucc, "pub成功", "")
	} else if strings.Compare(prkubeservice.Type, "18") == 0 {
		err = k8s.DeployApp(clienthandle, &prkubeservice, ClusterId, false)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "rollback失败"+err.Error(), "")
		}
		c.jsonResult(enums.JRCodeSucc, "rollback成功", "")
	} else {
		c.jsonResult(enums.JRCodeFailed, "TypeError", prkubeservice.Type)
	}

}

func (c *KubeServiceController) Scale() {
	prkubeservice := models.KubeServiceScale{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &prkubeservice)

	var ClusterId string

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)

	if strings.Compare(envUserCluster.ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
	} else {
		ClusterId = envUserCluster.ClusterId
	}

	clienthandle, err := models.GetApiServerHandle(ClusterId, false)


	replicaCountSpec, err := k8s.ScaleResource(clienthandle, prkubeservice.Kind, prkubeservice.Namespace, prkubeservice.Name, prkubeservice.Count)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "Scale error", "")

	}
	c.jsonResult(enums.JRCodeSucc, "Scale success", replicaCountSpec)

}
