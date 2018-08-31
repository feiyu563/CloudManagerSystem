package controllers

import (
	"fmt"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/satori/go.uuid"
	"encoding/json"
)

//KubeHostController host管理
type KubeHostController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *KubeHostController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList", "UpdateSeq")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *KubeHostController) Index() {
	//需要权限控制
	c.checkAuthor()
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backenduser/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "backenduser/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("BackendUserController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("BackendUserController", "Delete")
}

func (c *KubeHostController) DataGrid() {
	kubeHost :=models.KubeHostQueryParam{}
	//KubeHostQueryParam
	//if err := c.ParseForm(&kubeHost); err != nil {
	//	fmt.Println("handle error")
	//}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeHost)
	data, total := models.KubeHostPageList(&kubeHost)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *KubeHostController) GetALL() {
	var ClusterId string
	var data []*models.KubeHost
	var total int64
	c.Ctx.Input.Bind(&ClusterId, "ClusterId")

	if strings.Compare(ClusterId, "") == 0 {
		data, total = models.GetAllKubeHostWithNOCluster("")
	} else {
		data, total = models.GetAllKubeHostWithNOCluster(ClusterId)
	}

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}


func (c *KubeHostController) Save() {
	var err error
	m := models.KubeHost{}
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	o := orm.NewOrm()
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

func (c *KubeHostController) Delete() {
	strs := c.GetString("ids")
	ids := strings.Split(strs, ",")
	if num, err := models.KubeHostDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
