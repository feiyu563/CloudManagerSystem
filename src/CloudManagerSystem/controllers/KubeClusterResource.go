package controllers

import (
	"encoding/json"
	"CloudManagerSystem/models"
	"strings"
	"CloudManagerSystem/enums"

	"fmt"
)

type ClusterResourceController struct {
	BaseController
}

func (c *ClusterResourceController) GetALL() {
	var rolename string
	c.Ctx.Input.Bind(&rolename, "name")
	result := make(map[string]interface{})

	if strings.Compare(rolename, "") == 0 {
		c.jsonResult(enums.JRCodeFailed, "NO RoleName失败", "")
	}

	kubeCluster := models.ClusterResourceQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeCluster)

	data, total := models.KubeResourceOperPageList(&kubeCluster, rolename)
	//data, total := models.KubeClusterResourcePageList(&kubeCluster)
	//定义返回的数据结构
	result["total"] = total
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
	//ret:= models.ClusterRoleGetALL()
	//c.jsonResult(enums.JRCodeSucc,"ok",ret)
}

func (c *ClusterResourceController) Update() {
	//var kubeResource map[string][]models.ResourceOperQueryParam
	kubeResource := models.JSONResourceOperQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeResource)
	fmt.Println(kubeResource)

	var tt bool
	var err error
	//Resource:=kubeResource["ResourceOperQueryParam"]
	Resource := kubeResource.ResourceOperQueryParam
	if strings.Compare(kubeResource.RoleID, "") == 0 {
		c.jsonResult(enums.JRCodeFailed, "NO RoleName失败", "")
	}
	//fmt.Println(Resource)
	//fmt.Printf("%+v", Resource)
	//fmt.Println(kubeResource)
	var title string
	for k, _ := range Resource {
		fmt.Println(Resource[k])
		//fmt.Println(Resource[k].Id)

		//add
		if strings.Compare(Resource[k].Id, "") == 0 {
			title = "添加"
			tt, err = models.InsertResourceOper(*Resource[k], kubeResource.RoleID)
		} else {
			//update
			title = "编辑"
			tt, err = models.UpdateResourceOper(*Resource[k])
		}
		if !tt {
			c.jsonResult(enums.JRCodeFailed, title+"失败", err.Error())
		}
	}
	c.jsonResult(enums.JRCodeSucc, title+"成功", "")
}

func (c *ClusterResourceController) RsGetALL() {
	result := make(map[string]interface{})
	data, total := models.GetResourceALL()
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ClusterResourceController) OpGetALL() {
	result := make(map[string]interface{})
	data, total := models.GetOperALL()
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *ClusterResourceController) Delete() {
	strs := c.GetString("ids")
	ids := make([]string, 0, len(strs))
	for _, id := range strings.Split(strs, ",") {
		ids = append(ids, id)
	}
	if num, err := models.ResourceDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
