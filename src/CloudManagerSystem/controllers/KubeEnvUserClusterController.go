package controllers

import (
	"CloudManagerSystem/enums"
	"fmt"
	"CloudManagerSystem/models"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/satori/go.uuid"
	"encoding/json"
)

type KubeEnvUserClusterController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *KubeEnvUserClusterController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid", "DataList", "UpdateSeq")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}


func (c *KubeEnvUserClusterController) Get() {

	//u:= models.BackendUser{Id:"1"}
	u:=c.GetSessionUser()
	if (  len(u.Id) == 0) {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("id 不能为空"), 0)
	}

	result := make(map[string]interface{})
	data := models.SeleClusterEnvUser(u.Id)
	//定义返回的数据结构
	result["total"] = len(data)
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
}

func (c *KubeEnvUserClusterController) Post() {
	c.Save()
}


func (c *KubeEnvUserClusterController) Put() {
	c.Save()
}


func (c *KubeEnvUserClusterController) Save() {
	var err error
	m := models.KubeEnvUserCluster{}

	//if err = c.ParseForm(&m); err != nil {
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", "")
	}



	//m.UserId = "2"
	m.UserId = c.GetSessionUser().Id

	models.KubeEnvUserClusterDeleteByUserId(m.UserId)
	o := orm.NewOrm()
	//o.Delete(&models.KubeEnvUserCluster{UserId:m.UserId})
	var title string
	if(strings.Compare(m.Id,"")==0){
		u4,_:= uuid.NewV4()
		m.Id=u4.String()
		title="添加"
		_, err = o.Insert(&m)
	}else{
		title="编辑"
		_, err = o.Update(&m)
	}
	if err == nil {
		c.jsonResult(enums.JRCodeSucc, title+"成功", m.Id)
	} else {
		c.jsonResult(enums.JRCodeFailed, title+"失败", m.Id)
	}
}


