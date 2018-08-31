package controllers

type AppshopController struct {
	BaseController
}

func (c *AppshopController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("Namespaces", "Nodes")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}

//Namespaces
func (c *AppshopController) Apps() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="Apps"
	//页面模板设置
	c.setTpl("appshop/apps.html", "shared/layout_page.html")
	//c.LayoutSections = make(map[string]string)
	//c.LayoutSections["footerjs"] = "appshop/tables_js.html"

}

