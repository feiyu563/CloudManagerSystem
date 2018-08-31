package controllers

import (
	"encoding/json"
	"fmt"
	//"strconv"
	"strings"

	"CloudManagerSystem/enums"
	//"CloudManagerSystem/utils"
	//
	//"github.com/astaxie/beego/orm"
	"CloudManagerSystem/models"
	"CloudManagerSystem/models/k8s"
)

type ClusterRoleController struct {
	BaseController
}

func (c *ClusterRoleController) GetALL() {
	kubeCluster := models.ClusterRoleQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeCluster)
	var ClusterId string

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)

	if strings.Compare(envUserCluster.ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
	} else {
		ClusterId = envUserCluster.ClusterId
	}

	data, total := models.KubeClusterPageList(&kubeCluster, ClusterId)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
	//ret:= models.ClusterRoleGetALL()
	//c.jsonResult(enums.JRCodeSucc,"ok",ret)
}

func (c *ClusterRoleController) Get() {
	var ClusterId string
	result := make(map[string]interface{})

	c.Ctx.Input.Bind(&ClusterId, "ClusterId")

	if ClusterId == "" {
		c.jsonResult(enums.JRCodeFailed, "NO ClusterId失败", "")
	}

	kubeCluster := models.ClusterResourceQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeCluster)

	data, total := models.KubeClusterRoleGet(ClusterId)
	//定义返回的数据结构
	result["total"] = total
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
}

//
func (c *ClusterRoleController) Search() {
	fmt.Println("------Search-------")
	fmt.Println("")
	var s models.ClusterSearchResource
	json.Unmarshal(c.Ctx.Input.RequestBody, &s)
	fmt.Printf("%+v", s)
	fmt.Println("")

	if strings.Compare(s.Rolename, "") == 0 {
		c.jsonResult(enums.JRCodeFailed, s.Rolename+"Search失败", s)
	}

	ret := models.ResourceSearch(s)
	c.jsonResult(enums.JRCodeSucc, "ok", ret)
}

func (c *ClusterRoleController) Create() {
	//c.Ctx.Input.Params()
	fmt.Println("ClusterRoleCreate")
	fmt.Println(c.Ctx.Input.Params())
	c.jsonResult(enums.JRCodeSucc, "ok", "")

}

func (c *ClusterRoleController) Save() {
	kubeResource := models.JSONResourceOperQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeResource)
	fmt.Println(kubeResource)

	var tt bool
	var err error
	var title, roleid string

	var ClusterId string

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)

	if strings.Compare(envUserCluster.ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
	} else {
		ClusterId = envUserCluster.ClusterId
	}

	//Resource:=kubeResource["ResourceOperQueryParam"]
	if strings.Compare(kubeResource.RoleName, "") == 0 {
		c.jsonResult(enums.JRCodeFailed, "NO RoleName失败", "")
	} else {
		roleid, err = models.ClusterRoleInsert(kubeResource.RoleName, ClusterId, u.UserName)
		//fmt.Println(roleid)
		//fmt.Println("-----------------------------")
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, kubeResource.RoleName+"添加失败"+err.Error(), err)
		}
	}
	clienthandle, err := models.GetApiServerHandle(envUserCluster.ClusterId, false)

	k8s.K8sCreateClusterRole(clienthandle, kubeResource.RoleName)
	var datares int
	Resource := kubeResource.ResourceOperQueryParam
	datares = len(Resource)
	if datares == 0 {
		c.jsonResult(enums.JRCodeSucc, kubeResource.RoleName+"添加成功", kubeResource.RoleName)
	}
	//add
	for k, _ := range Resource {

		if strings.Compare(Resource[k].Id, "") == 0 {
			title = "添加"
			tt, err = models.InsertResourceOper(*Resource[k], roleid)
		} else {
			//update
			title = "编辑"
			tt, err = models.UpdateResourceOper(*Resource[k])
		}
		if !tt {
			c.jsonResult(enums.JRCodeFailed, title+"失败", err)
		}
	}

	c.jsonResult(enums.JRCodeSucc, kubeResource.RoleName+"添加成功", kubeResource.RoleName)
}

func (c *ClusterRoleController) Delete() {
	strs := c.GetString("ids")
	ids := strings.Split(strs, ",")

	if num, err := models.KubeClusterRoleDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}

}
