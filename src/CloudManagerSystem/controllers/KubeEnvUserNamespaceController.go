package controllers

import (
	"fmt"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/satori/go.uuid"
)

type KubeEnvUserNamespaceController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *KubeEnvUserNamespaceController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid", "DataList", "UpdateSeq")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *KubeEnvUserNamespaceController) Get() {

	clusterId := c.GetString("clusterId")

	//u:= models.BackendUser{Id:"1"}
	u := c.GetSessionUser()
	if (len(clusterId) == 0) {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("clusterId 不能为空"), 0)
	}

	result := make(map[string]interface{})
	data := models.SeleNameSpaceEnvUser(u.Id,clusterId)
	//定义返回的数据结构
	result["total"] = len(data)
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
}

func (c *KubeEnvUserNamespaceController) Post() {
	c.Save()
}

func (c *KubeEnvUserNamespaceController) Put() {
	c.Save()
}

func (c *KubeEnvUserNamespaceController) Save() {
	var err error
	ms := models.KubeEnvUserNamespaceMore{}

	//if err = c.ParseForm(&m); err != nil {
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &ms); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", "")
	}

	//userId := "2"
	userId := c.GetSessionUser().Id

	models.KubeEnvUserNameSpaceDeleteByUserId(userId)
	for _, m := range ms.Eunsm {

		m.UserId = userId
		m.ClusterId = ms.ClusterId

		o := orm.NewOrm()
		o.Delete(&models.KubeEnvUserNamespace{UserId:m.UserId})
		if (strings.Compare(m.Id, "") == 0) {
			u4, _ := uuid.NewV4()
			m.Id = u4.String()
			_, err = o.Insert(m)
		} else {
			_, err = o.Update(m)
		}

	}
	if err == nil {
		c.jsonResult(enums.JRCodeSucc, "成功", "")
	} else {
		c.jsonResult(enums.JRCodeFailed, "失败", "")
	}
}
