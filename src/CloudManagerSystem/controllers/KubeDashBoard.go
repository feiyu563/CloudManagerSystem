package controllers

type KubeDashBoardController struct {
	BaseController
}

func (c *KubeDashBoardController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("Namespaces", "Nodes")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}

//Nodes
func (c *KubeDashBoardController) Nodes() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="Nodes"
	//页面模板设置
	c.setTpl("kube/kube_nodes.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_nodes_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_nodes_footerjs.html"
}
//Deployment
func (c *KubeDashBoardController) Deployment() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="Deployment"
	//页面模板设置
	c.setTpl("kube/kube_deployment.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_deployment_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_deployment_footerjs.html"
}
//Edit 编辑Deployment界面
func (c *KubeDashBoardController) Edit_Deployment() {
	c.Data["name"]= c.Input().Get("name")
	c.Data["namespace"]=c.Input().Get("namespace")
	c.Data["kind"]=c.Input().Get("kind")
	c.setTpl("kube/edit_deployment.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "kube/edit_deployment_footerjs.html"
}
//Deployment_detail
func (c *KubeDashBoardController) Deployment_detail() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor("KubeDashBoardController.Deployment")
	//页面title设置
	c.Data["pageTitle"]="Deployment_detail"
	//页面模板设置
	c.Data["Namespace"]=c.Input().Get("namespace")
	c.Data["Name"]=c.Input().Get("name")
	c.setTpl("kube/kube_deployment_detail.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_deployment_detail_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_deployment_detail_footerjs.html"
}
//Pod
func (c *KubeDashBoardController) Pod() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="Pod"
	//页面模板设置
	c.setTpl("kube/kube_pod.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_pod_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_pod_footerjs.html"
}
//Services
func (c *KubeDashBoardController) Services() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="Services"
	//页面模板设置
	c.setTpl("kube/kube_services.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_services_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_services_footerjs.html"
}
//Deployment_detail
func (c *KubeDashBoardController) Services_detail() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor("KubeDashBoardController.Services")
	//页面title设置
	c.Data["pageTitle"]="Services_detail"
	//页面模板设置
	c.Data["Namespace"]=c.Input().Get("namespace")
	c.Data["Name"]=c.Input().Get("name")
	c.setTpl("kube/kube_services_detail.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_services_detail_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_services_detail_footerjs.html"
}
//Namespaces
func (c *KubeDashBoardController) Namespaces() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="Namespaces"
	//页面模板设置
	c.setTpl("kube/kube_namespaces.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_namespaces_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_namespaces_footerjs.html"
}
//Pvs
func (c *KubeDashBoardController) Pvs() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="PersistentVolume"
	//页面模板设置
	c.setTpl("kube/kube_pvs.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_pvs_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_pvs_footerjs.html"
}
//Pvs_detail
func (c *KubeDashBoardController) Pvs_detail() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor("KubeDashBoardController.Pvs")
	//页面title设置
	c.Data["pageTitle"]="PersistentVolume_detail"
	//页面模板设置
	c.Data["Namespace"]=c.Input().Get("namespace")
	c.Data["Name"]=c.Input().Get("name")
	c.setTpl("kube/kube_pvs_detail.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_pvs_detail_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_pvs_detail_footerjs.html"
}
//StorageClass
func (c *KubeDashBoardController) StorageClass() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="StorageClass"
	//页面模板设置
	c.setTpl("kube/kube_storageclass.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_storageclass_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_storageclass_footerjs.html"
}
//StorageClass_detail
func (c *KubeDashBoardController) StorageClass_detail() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor("KubeDashBoardController.StorageClass")
	//页面title设置
	c.Data["pageTitle"]="StorageClass_detail"
	//页面模板设置
	c.Data["Namespace"]=c.Input().Get("namespace")
	c.Data["Name"]=c.Input().Get("name")
	c.setTpl("kube/kube_storageclass_detail.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_storageclass_detail_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_storageclass_detail_footerjs.html"
}

//创建sc界面
func (c *KubeDashBoardController) CreateCephStorageClass() {
	c.Data["namespce"]=c.Input().Get("namespace")
	c.setTpl("kube/kube_storageclass_create.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "kube/kube_storageclass_create_footerjs.html"
}

//Statefulset
func (c *KubeDashBoardController) Statefulset() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="Statefulset"
	//页面模板设置
	c.setTpl("kube/kube_statefulset.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_statefulset_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_statefulset_footerjs.html"
}
//DaemonSet
func (c *KubeDashBoardController) DaemonSet() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="DaemonSet"
	//页面模板设置
	c.setTpl("kube/kube_daemonset.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_daemonset_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_daemonset_footerjs.html"
}

//CronJob
func (c *KubeDashBoardController) CronJob() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="CronJob"
	//页面模板设置
	c.setTpl("kube/kube_cronjob.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_cronjob_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_cronjob_footerjs.html"
}