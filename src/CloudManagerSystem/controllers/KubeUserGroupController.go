package controllers

import (
	"CloudManagerSystem/enums"
	"CloudManagerSystem/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"strings"
	"fmt"
	"github.com/json-iterator/go"
	"CloudManagerSystem/models/k8s"
)

//KubeHostController host管理
type KubeUserGroupController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *KubeUserGroupController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList", "UpdateSeq")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *KubeUserGroupController) DataGrid() {
	vo := models.KubeUserGroupQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &vo)
	data, total := models.KubeUserGroupPageList(&vo)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *KubeUserGroupController) AllList() {
	vo := models.KubeUserGroupQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &vo)
	data, _ := models.KubeUserGroupAllList(&vo)
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *KubeUserGroupController) Save() {
	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	var err error

	clienthandle, err := models.GetApiServerHandle(envUserCluster.ClusterId, false)

	m := models.KubeUserGroupVO{KubeUserGroup: &models.KubeUserGroup{}}
	json_iterator.Unmarshal(c.Ctx.Input.RequestBody, &m)

	if len(m.KubeBinds) == 0 {
		c.jsonResult(enums.JRCodeFailed, "no clusterrole bind!", m.KubeUserGroup.GroupName)
	}

	o := orm.NewOrm()
	var title string
	if strings.Compare(m.KubeUserGroup.Id, "") == 0 {
		title = "添加"
		tem_uuid_t, _ := uuid.NewV4()
		m.KubeUserGroup.Id = tem_uuid_t.String()
		m.KubeUserGroup.CreateUser = u.Id
		//当前创建
		m.KubeUserGroup.ClusterId = envUserCluster.ClusterId
		_, err = o.Insert(m.KubeUserGroup)
		//err = k8s.K8sCreateClusterRoleBinding(clienthandle, "", m.KubeUserGroup.GroupName, "") //for no resource test
		if err != nil {
			fmt.Println(err)
		}
	} else {
		title = "编辑"
		m_old := &models.KubeUserGroup{Id: m.KubeUserGroup.Id}
		o.Read(m_old)
		k8s.K8sDeleteClusterRoleBinding(clienthandle, m_old.GroupName) //for no resource test

		_, err = o.Update(m.KubeUserGroup)
		//err = k8s.K8sCreateClusterRoleBinding(clienthandle, "", m.KubeUserGroup.GroupName, "") //for no resource test
		if err != nil {
			fmt.Println(err)

		}
		//删除Bind
		models.KubeBindDeleteByUserId(m.KubeUserGroup.Id)
	}
	for _, kubeBind := range m.KubeBinds {
		tem_uuid_t, _ := uuid.NewV4()
		kubeBind.Id = tem_uuid_t.String()
		kubeBind.UserId = m.KubeUserGroup.Id
		kubeBind.UserType = 1
		kubeBind.ClusterId = envUserCluster.ClusterId
		o.Insert(kubeBind)
		ClusterRoleName := models.KubeClusterGetOne(kubeBind.RoleId).Name
		namespace := models.GetEnvNameSpaceByUserIdClusterId(u.Id,envUserCluster.ClusterId).Name
		err = k8s.K8sCreateClusterRoleBinding(clienthandle, ClusterRoleName, m.KubeUserGroup.GroupName, namespace) //for no resource test
	}
	if err == nil {
		c.jsonResult(enums.JRCodeSucc, title+"成功", m.KubeUserGroup.GroupName)
	} else {
		c.jsonResult(enums.JRCodeFailed, title+"失败", m.KubeUserGroup.GroupName)
	}
}

func (c *KubeUserGroupController) Delete() {
	strs := c.GetString("ids")
	ids := strings.Split(strs, ",")
	//删除用户组和角色关联
	_, err := models.KubeBindDeleteByUserIds(ids)
	if (err == nil) {
		//删除用户组
		if num, err := models.KubeUserGroupDelete(ids); err == nil {
			c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
		} else {
			c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
		}
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
