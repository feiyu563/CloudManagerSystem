package controllers

import (
	"CloudManagerSystem/models"
	"encoding/json"
	"CloudManagerSystem/enums"
)


type KubePublishServicePathController struct {
	BaseController
}

//分页展示NameSpace
func (c *KubePublishServicePathController) DataGrid(){
	KubePublishServicePathParam :=models.KubePublishServicePathQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &KubePublishServicePathParam)

	if(len(KubePublishServicePathParam.PserviceId) ==0){
		c.jsonResult(enums.JRCodeSucc, "PserviceId不能为空", 0)
		return
	}

	data, total := models.KubePublishServicePathPageList(&KubePublishServicePathParam)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}