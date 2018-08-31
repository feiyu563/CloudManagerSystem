package controllers

import (
	"CloudManagerSystem/models"
)

type KubeBindController struct {
	BaseController
}
func (c *KubeBindController) DataGrid() {
	userId:=c.GetString("userId")
	data, _ := models.FindKubeBindsByUserId(userId)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}