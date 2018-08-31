package controllers

import (
	"CloudManagerSystem/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/satori/go.uuid"
	"encoding/json"
	"CloudManagerSystem/enums"
)

type KubePublishProxyController struct {
	BaseController
}


//分页展示NameSpace
func (c *KubePublishProxyController) DataGrid(){
	KubePublishProxyParam :=models.KubePublishProxyQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &KubePublishProxyParam)

	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)

	if(len(envUserCluster.ClusterId) ==0){
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		return
	}else {
		KubePublishProxyParam.ClusterId = envUserCluster.ClusterId
	}

	data, total := models.KubePublishProxyPageList(&KubePublishProxyParam)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//获取一个
func(c *KubePublishProxyController) Get(){
	strs := c.GetString("id")
	if(len(strs) == 0){
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("id 不能为空"), 0)
	}
	c.Ctx.Input.RequestBody =[]byte("{\"Id\":\""+strs+"\"}")
	c.DataGrid()
}

//添加
func(c *KubePublishProxyController)Post(){
	c.Save()
}

//修改
func(c *KubePublishProxyController)Put(){
	c.Save()
}

//删除
func(c *KubePublishProxyController)Delete(){
	strs := c.GetString("ids")
	ids := make([]string, 0, len(strs))
	for _, id := range strings.Split(strs, ",") {
		ids = append(ids, id)
	}
	if num, err := models.KubePublishProxyDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

//保存方法 (实现 添加与修改功能)
func (c *KubePublishProxyController) Save(){
	var err error
	m := models.KubePublishProxy{}
	/*if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
		return
	}*/
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", "")
	}

	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	if(len(envUserCluster.ClusterId) ==0){
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		return
	}

	o := orm.NewOrm()
	var title string
	m.ClusterId = envUserCluster.ClusterId
	m.CreateUser = u.Id
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