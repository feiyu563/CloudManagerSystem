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
)

//KubeHostController host管理
type KubeServiceAccountsController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *KubeServiceAccountsController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *KubeServiceAccountsController) DataGrid() {
	vo := models.KubeServiceAccountsQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &vo)
	data, total := models.KubeServiceAccountsPageList(&vo)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *KubeServiceAccountsController) Save() {
	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	var err error
	m := models.KubeServiceAccountsVO{KubeServiceAccounts:&models.KubeServiceAccounts{}}
	json_iterator.Unmarshal( c.Ctx.Input.RequestBody, &m)
	o := orm.NewOrm()
	var title string
	if strings.Compare(m.KubeServiceAccounts.Id, "") == 0 {
		title = "添加"
		tem_uuid_t, _ := uuid.NewV4()
		m.KubeServiceAccounts.Id = tem_uuid_t.String()
		m.KubeServiceAccounts.CreateUser=u.Id
		//当前创建
		m.KubeServiceAccounts.ClusterId=envUserCluster.ClusterId
		_, err = o.Insert(m.KubeServiceAccounts)
	} else {
		title = "编辑"
		_, err = o.Update(m.KubeServiceAccounts)
		//删除Bind
		models.KubeBindDeleteByUserId(m.KubeServiceAccounts.Id)
	}
	for _,kubeBind :=range m.KubeBinds{

		tem_uuid_t, _ := uuid.NewV4()
		kubeBind.Id = tem_uuid_t.String()
		kubeBind.UserId = m.KubeServiceAccounts.Id
		kubeBind.UserType = 1
		kubeBind.ClusterId = envUserCluster.ClusterId
		o.Insert(kubeBind)
	}
	if err == nil {
		c.jsonResult(enums.JRCodeSucc, title+"成功", m.KubeServiceAccounts.Name)
	} else {
		c.jsonResult(enums.JRCodeFailed, title+"失败", m.KubeServiceAccounts.Name)
	}
}

func (c *KubeServiceAccountsController) Delete() {
	strs := c.GetString("ids")
	ids:= strings.Split(strs, ",")
	//删除用户组和角色关联
	_,err:=models.KubeBindDeleteByUserIds(ids)
	if(err==nil){
		//删除用户组
		if num, err := models.KubeServiceAccountsDelete(ids); err == nil {
			c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
		} else {
			c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
		}
	}else{
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
