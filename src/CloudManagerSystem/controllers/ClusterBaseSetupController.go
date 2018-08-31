package controllers

import (
	"CloudManagerSystem/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"CloudManagerSystem/enums"
	"strings"
	"github.com/satori/go.uuid"
)

type BaseclusterSetupController struct {
	BaseController
}

func (c *BaseclusterSetupController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("Namespaces", "Nodes")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}

//NameSpaceSetup页面
func (c *BaseclusterSetupController) NameSpaceSetup() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="智享云基础信息管理"
	c.setTpl("basesetup/index_cluster_setup_namespace.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basesetup/index_cluster_setup_namespace_headcssjs.html"
	c.LayoutSections["footerjs"] = "basesetup/index_cluster_setup_namespace_footerjs.html"
}

//EditNamespace 弹窗页面
func (c *BaseclusterSetupController) EditNamespace() {
	Id:= c.Input().Get("id")
	kubeNameSpace := models.KubeNameSpaceQueryParam{Id: Id}
	if Id !="0" {
		data, _ := models.KubeNameSpacePageList(&kubeNameSpace)
		beego.Error(data[0])
		c.Data["m"] = data[0]
		} else {
		u:=c.GetSessionUser()
		envUserCluster,_ :=models.EnvUserCluster(u.Id)
		//envUserCluster.Id
		m := models.KubeCluster{Id:envUserCluster.ClusterId}
		o := orm.NewOrm()
		o.Read(&m)
		c.Data["ClusterName"]=m.Name
	}
	//KubeNameSpacePageList
	c.setTpl("basesetup/edit_cluster_setup_namespace.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "basesetup/edit_cluster_setup_namespace_footerjs.html"
}

//GrantNamespace 命名空间授权界面 弹窗页面
func (c *BaseclusterSetupController) GrantNamespace() {
	id:=c.Input().Get("id")
	kubeNameSpace :=models.KubeNameSpaceQueryParam{Id:id}
	data, _ := models.KubeNameSpacePageList(&kubeNameSpace)
	c.Data["m"] = data[0]
	c.setTpl("basesetup/allocation_namespace/allocation_namespace.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basesetup/allocation_namespace/allocation_namespace_headcssjs.html"
	c.LayoutSections["footerjs"] = "basesetup/allocation_namespace/allocation_namespace_footerjs.html"
}

//ClusterRole集群角色-集群页面
func (c *BaseclusterSetupController) ClusterRole() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="智享云基础信息管理"
	c.setTpl("basesetup/index_cluster_role.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basesetup/index_cluster_role_headcssjs.html"
	c.LayoutSections["footerjs"] = "basesetup/index_cluster_role_footerjs.html"
}

//EditClusterRole 弹窗页面
func (c *BaseclusterSetupController) EditClusterRole() {
	id:=c.Input().Get("id")
	if id!="0" {
		m := models.ClusterRole{Id:id}
		o := orm.NewOrm()
		o.Read(&m)
		c.Data["url"]=c.URLFor("ClusterResourceController.GetALL")+"/?name="
		c.Data["m"] = m
	}else {
		c.Data["url"]=""
	}
	c.setTpl("basesetup/edit_cluster_role/edit_cluster_role.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basesetup/edit_cluster_role/edit_cluster_role_headcssjs.html"
	c.LayoutSections["footerjs"] = "basesetup/edit_cluster_role/edit_cluster_role_footerjs.html"
}

//GrantUserGroup 用户组授权-集群授权界面
func (c *BaseclusterSetupController) GrantUserGroup() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="智享云基础信息管理"
	c.setTpl("basesetup/index_cluster_grant_usergroup.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basesetup/index_cluster_grant_usergroup_headcssjs.html"
	c.LayoutSections["footerjs"] = "basesetup/index_cluster_grant_usergroup_footerjs.html"
}

//EditClusterGrantUserGroup 弹窗页面
func (c *BaseclusterSetupController) EditClusterGrantUserGroup() {
	id:=c.Input().Get("id")
	vo:=&models.KubeUserGroupVO{}
	if id!="0" {
		o := orm.NewOrm()
		m := models.KubeUserGroup{Id:id}
		o.Read(&m)
		cluter:=&models.KubeCluster{Id:m.ClusterId}
		o.Read(cluter)
		m.ClusterName=cluter.Name
		vo.KubeUserGroup=&m
		c.Data["url"]=c.URLFor("KubeBindController.DataGrid")+"/?userId="+id
	}else {
		u:=c.GetSessionUser()
		o := orm.NewOrm()
		envUserCluster,_ :=models.EnvUserCluster(u.Id)
		vo.KubeUserGroup=&models.KubeUserGroup{ClusterName:envUserCluster.ClusterName,ClusterId:envUserCluster.ClusterId}
		cluter:=&models.KubeCluster{Id:vo.KubeUserGroup.ClusterId}
		o.Read(cluter)
		vo.KubeUserGroup.ClusterName=cluter.Name
		c.Data["url"]=""
	}
	c.Data["m"] = vo
	c.setTpl("basesetup/edit_cluster_grant_usergroup/edit_cluster_grant_usergroup.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basesetup/edit_cluster_grant_usergroup/edit_cluster_grant_usergroup_headcssjs.html"
	c.LayoutSections["footerjs"] = "basesetup/edit_cluster_grant_usergroup/edit_cluster_grant_usergroup_footerjs.html"
}

//Sa 界面
func (c *BaseclusterSetupController) Sa() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="智享云基础信息管理"
	c.setTpl("basesetup/index_cluster_sa.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basesetup/index_cluster_sa_headcssjs.html"
	c.LayoutSections["footerjs"] = "basesetup/index_cluster_sa_footerjs.html"
}

//EditSa 弹窗页面
func (c *BaseclusterSetupController) EditSa() {
	vo:=&models.KubeServiceAccountsVO{}
	id:=c.Input().Get("id")
	if id!="0" {
		m := models.KubeServiceAccounts{Id: id}
		o := orm.NewOrm()
		o.Read(&m)

		cluter:=&models.KubeCluster{Id:m.ClusterId}
		o.Read(cluter)
		m.ClusterName=cluter.Name
		vo.KubeServiceAccounts=&m
		c.Data["url"]=c.URLFor("KubeBindController.DataGrid")+"/?userId="+id
	}else {
		u:=c.GetSessionUser()
		o := orm.NewOrm()
		envUserCluster,_ :=models.EnvUserCluster(u.Id)
		vo.KubeServiceAccounts=&models.KubeServiceAccounts{ClusterName:envUserCluster.ClusterName,ClusterId:envUserCluster.ClusterId}
		cluter:=&models.KubeCluster{Id:vo.KubeServiceAccounts.ClusterId}
		o.Read(cluter)
		vo.KubeServiceAccounts.ClusterName=cluter.Name
		c.Data["url"]=""
	}
	c.Data["m"] = vo
	c.setTpl("basesetup/edit_cluster_sa/edit_cluster_sa.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basesetup/edit_cluster_sa/edit_cluster_sa_headcssjs.html"
	c.LayoutSections["footerjs"] = "basesetup/edit_cluster_sa/edit_cluster_sa_footerjs.html"
}

func (c *BaseclusterSetupController) New()  {
	beego.Error(string(c.Ctx.Input.RequestBody))
	var err error
	m := models.KubeUserGroupVO{}

	if err = c.ParseForm(&m); err != nil {
	//if err = json.Unmarshal(c.Ctx.Input.RequestBody, &m); err != nil {
		beego.Error("\n\n\n\n\n\n\n\n\n\n\n4")
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.KubeUserGroup.GroupName)
	}
	beego.Error("\n\n\n\n\n\n\n\n\n\n\n3")
	beego.Error(m)
	o := orm.NewOrm()
	var title string
	if strings.Compare(m.KubeUserGroup.Id, "") == 0 {
		title = "添加"
		tem_uuid_t, _ := uuid.NewV4()
		m.KubeUserGroup.Id = tem_uuid_t.String()
		_, err = o.Insert(&m.KubeUserGroup)
	} else {
		title = "编辑"
		_, err = o.Update(&m.KubeUserGroup)
		//删除Bind
		models.KubeBindDeleteByUserId(m.KubeUserGroup.Id)
	}
	for _,kubeBind :=range m.KubeBinds{
		tem_uuid_t, _ := uuid.NewV4()
		kubeBind.Id = tem_uuid_t.String()
		kubeBind.UserId = m.KubeUserGroup.Id
		kubeBind.UserType = 1
		kubeBind.ClusterId = m.ClusterId
		o.Insert(kubeBind)
	}
	if err == nil {
		c.jsonResult(enums.JRCodeSucc, title+"成功", m.KubeUserGroup.GroupName)
	} else {
		c.jsonResult(enums.JRCodeFailed, title+"失败", m.KubeUserGroup.GroupName)
	}
	c.jsonResult(enums.JRCodeFailed, "成功", 0)

}