package controllers

import (
	//"encoding/json"
	"fmt"
	"strings"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	//	"github.com/astaxie/beego/orm"
	//	"github.com/satori/go.uuid"
	//	"fmt"
)

type KubeServicePortController struct {
	BaseController
}

func (c *KubeServicePortController) ServciePort() {
	var ServiceId string
	result := make(map[string]interface{})

	c.Ctx.Input.Bind(&ServiceId, "ServiceId")

	if ServiceId == "" {
		c.jsonResult(enums.JRCodeFailed, "ServiceId失败", "")
	}

	data, total := models.KubeServicePortGet(ServiceId)
	//定义返回的数据结构
	result["total"] = total
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
}


func (c *KubeServicePortController) Delete() {
	strs := c.GetString("ids")
	ids := make([]string, 0, len(strs))
	for _, id := range strings.Split(strs, ",") {
		ids = append(ids, id)
	}
	if num, err := models.ServicePortDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
