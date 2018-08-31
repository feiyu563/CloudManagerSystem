package controllers

import (
	"CloudManagerSystem/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"os"
	"CloudManagerSystem/enums"
)

type BaseclusterController struct {
	BaseController
}

func (c *BaseclusterController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("Namespaces", "Nodes")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}

//Hosts
func (c *BaseclusterController) Hosts() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="智享云主机管理"
	c.setTpl("basecluster/index_hosts.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basecluster/index_hosts_headcssjs.html"
	c.LayoutSections["footerjs"] = "basecluster/index_hosts_footerjs.html"
}
//Edit 添加、编辑Host界面
func (c *BaseclusterController) Edit() {
	Id:= c.Input().Get("id")
	m := models.KubeHost{Id: Id}
	if Id !="0" {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
		c.Data["m"] = m
	}

	c.setTpl("basecluster/edit_hosts.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "basecluster/edit_hosts_footerjs.html"
}

//Cluster
func (c *BaseclusterController) Cluster() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="智享云集群管理"
	c.setTpl("basecluster/index_cluster.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basecluster/index_cluster_headcssjs.html"
	c.LayoutSections["footerjs"] = "basecluster/index_cluster_footerjs.html"
}
//Edit 添加、编辑Cluster界面
func (c *BaseclusterController) EditCluster() {
	Id:= c.Input().Get("id")
	m := models.KubeCluster{Id: Id}
	if Id !="0" {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
		c.Data["m"] = m
	}
	c.setTpl("basecluster/edit_cluster.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "basecluster/edit_cluster_footerjs.html"
}

//kubeconfig文件上传
func (c *BaseclusterController) Saveconfig()  {
	Id:= c.Input().Get("id")
	m := models.KubeCluster{Id: Id}
	o := orm.NewOrm()
	err := o.Read(&m)
	if err != nil {
		c.pageError("数据无效，请刷新后重试")
	}
	c.Data["m"] = m
	dr:="conf/kubeconfig/"+Id+"/config"
	if c.CheckD(dr,false){
		c.Data["Filename"] = "config"
	}
	c.setTpl("basecluster/edit_cluster_config.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "basecluster/edit_cluster_config_footerjs.html"
}
//Save文件上传
func (c *BaseclusterController) Savecfg()  {
	Id:= c.Input().Get("Id")
	//获取上传的附件
	_,Kubeconfig,file_err:=c.GetFile("Kubeconfig")
	if file_err!=nil {
		beego.Error(file_err)
	}
	if Kubeconfig!=nil {
		//保存临时文件
		dr:="conf/kubeconfig/"+Id+"/"
		c.CheckD(dr,true)
		file_err=c.SaveToFile("Kubeconfig",dr+Id)
		if file_err!=nil {
			beego.Error(file_err)
			c.jsonResult(enums.JRCodeFailed, "上传失败"+file_err.Error(), file_err)
		}
		//开始判断是否可以正常用于连接
		_,err:=models.GetApiServerHandle(Id,true)
		if err!=nil {
			os.Remove("conf/kubeconfig/"+Id+"/"+Id)
			c.jsonResult(enums.JRCodeFailed, "绑定失败"+err.Error(), err)
		}else {
			os.Remove("conf/kubeconfig/"+Id+"/"+Id)
			c.SaveToFile("Kubeconfig","conf/kubeconfig/"+Id+"/config")
		}
		c.jsonResult(enums.JRCodeSucc, "上传成功", "")
	}else {
		c.jsonResult(enums.JRCodeSucc, "", "")
	}
}
func (c *BaseclusterController)CheckD(d string,ifcreate bool)(bool)  {
	_, err := os.Stat(d)
	if err != nil {
		if ifcreate{os.Mkdir(d,0777)}
		return false
	}else {
		return true
	}
}

//AllocationNode 分配节点界面
func (c *BaseclusterController) AllocationNode() {
	c.Data["clusterid"] = c.Input().Get("clusterid")
	c.setTpl("basecluster/allocation_node.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basecluster/allocation_node_headcssjs.html"
	c.LayoutSections["footerjs"] = "basecluster/allocation_node_footerjs.html"
}

//AllocationRole 分配权限界面
func (c *BaseclusterController) AllocationRole() {
	c.Data["clusterid"] = c.Input().Get("clusterid")
	c.setTpl("basecluster/allocation_role.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basecluster/allocation_role_headcssjs.html"
	c.LayoutSections["footerjs"] = "basecluster/allocation_role_footerjs.html"
}

//ClusterSetup
func (c *BaseclusterController) ClusterSetup() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面title设置
	c.Data["pageTitle"]="智享云主机管理"
	c.setTpl("basecluster/index_cluster_setup.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "basecluster/index_cluster_setup_headcssjs.html"
	c.LayoutSections["footerjs"] = "basecluster/index_cluster_setup_footerjs.html"
}
