package controllers

import (
	"CloudManagerSystem/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)
type KubeController struct {
	BaseController
}

func (c *KubeController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("Namespaces", "Nodes")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}
//deploy
func (c *KubeController) Deploy() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="deploy"
	//页面模板设置
	c.setTpl("kube/deploy.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/deploy_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/deploy_footerjs.html"
}

//editdeploy
func (c *KubeController) EditDeploy() {
	Id:= c.Input().Get("id")
	if Id !="0" {
		c.Data["Id"] = Id
		service:=models.KubeService{Id:Id}
		o := orm.NewOrm()
		o.Read(&service)
		c.Data["m"]=service
		c.Data["url"]=c.URLFor("KubeServicePortController.ServciePort")+"/?ServiceId="+Id
	}else {
		c.Data["url"]=""
	}
	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	c.Data["clusterId"]=envUserCluster.ClusterId
	c.setTpl("kube/kube_deploy/deploy/kube_deploy.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_deploy/deploy/kube_deploy_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_deploy/deploy/kube_deploy_footerjs.html"
}

//publish
func (c *KubeController) Publish() {
	id:= c.Input().Get("id")
	c.Data["Id"] = id
	c.setTpl("kube/kube_deploy/publish/kube_publish.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_deploy/publish/kube_publish_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_deploy/publish/kube_publish_footerjs.html"
}

//history
func (c *KubeController) History() {
	id:= c.Input().Get("id")
	c.Data["Id"] = id
	c.setTpl("kube/kube_deploy/history/kube_history.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_deploy/history/kube_history_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_deploy/history/kube_history_footerjs.html"
}

//Manage
func (c *KubeController) Manage() {
	id:= c.Input().Get("id")
	namespace:=c.Input().Get("namespace")
	c.Data["Id"] = id
	c.Data["Namespace"] = namespace
	c.setTpl("kube/kube_deploy/manage/kube_manage.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_deploy/manage/kube_manage_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_deploy/manage/kube_manage_footerjs.html"
}

//scale
func (c *KubeController) Scale() {
	Id:= c.Input().Get("id")
	service:=models.KubeService{Id:Id}
	o := orm.NewOrm()
	o.Read(&service)
	c.Data["m"]=service
	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	c.Data["clusterId"]=envUserCluster.ClusterId
	c.setTpl("kube/kube_deploy/scale/kube_scale.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_deploy/scale/kube_scale_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_deploy/scale/kube_scale_footerjs.html"
}

//deployservice
func (c *KubeController) DeployProxy() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="deploy"
	//页面模板设置
	c.setTpl("kube/deploy_proxy.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/deploy_proxy_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/deploy_proxy_footerjs.html"
}

//deployservice
func (c *KubeController) DeployIngress() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="deploy"
	//页面模板设置
	c.setTpl("kube/deploy_ingress.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/deploy_ingress_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/deploy_ingress_footerjs.html"
}

//editdeployservice
func (c *KubeController) EditDeployIngress() {
	Id:= c.Input().Get("id")
	if Id !="0" {
		c.Data["Id"] = Id
		service:=models.KubePublishService{Id:Id}
		o := orm.NewOrm()
		o.Read(&service)
		c.Data["m"]=service
		c.Data["url"]=c.URLFor("KubePublishServicePathController.DataGrid")
	}else {
		c.Data["url"]=""
	}
	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	c.Data["clusterId"]=envUserCluster.ClusterId
	c.setTpl("kube/kube_deploy_ingress_service/kube_deploy_ingress_service.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_deploy_ingress_service/kube_deploy_ingress_service_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_deploy_ingress_service/kube_deploy_ingress_service_footerjs.html"
}

//editdeployservice
func (c *KubeController) EditDeployProxy() {
	Id:= c.Input().Get("id")
	if Id !="0" {
		c.Data["Id"] = Id
		service:=models.KubePublishService{Id:Id}
		o := orm.NewOrm()
		o.Read(&service)
		c.Data["m"]=service
		c.Data["url"]=c.URLFor("KubePublishServicePathController.DataGrid")
	}else {
		c.Data["url"]=""
	}
	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	c.Data["clusterId"]=envUserCluster.ClusterId
	c.setTpl("kube/kube_deploy_proxy_service/kube_deploy_proxy_service.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/kube_deploy_proxy_service/kube_deploy_proxy_service_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/kube_deploy_proxy_service/kube_deploy_proxy_service_footerjs.html"
}

//images
func (c *KubeController) Images() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="Images"
	//页面模板设置
	c.setTpl("kube/images.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/images_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/images_footerjs.html"
}

//Repositories
func (c *KubeController) Repositories() {
	//是否显示更多查询条件的按钮K
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor("KubeController.Images")
	//页面title设置
	c.Data["pageTitle"]="Images"
	c.Data["id"]=c.Input().Get("id")
	//页面模板设置
	c.setTpl("kube/images_repositories.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/images_repositories_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/images_repositories_footerjs.html"
}
//Tags
func (c *KubeController) Tags() {
	c.Data["name"]=c.Input().Get("name")
	c.Data["imageurl"]=beego.AppConfig.String("harbor::harbor_ip")
	c.setTpl("kube/images_tags.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "kube/images_tags_headcssjs.html"
	c.LayoutSections["footerjs"] = "kube/images_tags_footerjs.html"
}

//editdeployservice
func (c *KubeController) EditImages() {
	Id:= c.Input().Get("id")
	if Id !="0" {
		c.Data["Id"] = Id
		Images:=models.KubeImage{Id:Id}
		o := orm.NewOrm()
		o.Read(&Images)
		c.Data["m"]=Images
		c.Data["url"]=c.URLFor("KubePublishServicePathController.DataGrid")
	}else {
		c.Data["url"]=""
	}
	//u:=c.GetSessionUser()
	//envUserCluster,_ :=models.EnvUserCluster(u.Id)
	//c.Data["clusterId"]=envUserCluster.ClusterId
	c.setTpl("kube/images/edit_images.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "kube/images/edit_images_footerjs.html"
}