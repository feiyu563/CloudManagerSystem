package routers

import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/base/clusterrole/dataGrid", &controllers.ClusterRoleController{}, "GET,POST:GetALL")
	beego.Router("/base/clusterrole/", &controllers.ClusterRoleController{}, "GET:Get")
	//beego.Router("/base/clusterrole", &controllers.ClusterRoleController{}, "POST:Search")
	//beego.Router("/base/clusterrole", &controllers.ClusterRoleController{}, "POST:ClusterRoleCreate")
	beego.Router("/base/clusterrole", &controllers.ClusterRoleController{}, "POST,PUT:Save")
	beego.Router("/base/clusterrole", &controllers.ClusterRoleController{}, "DELETE:Delete")

	//ns := beego.NewNamespace("/base",
	//	beego.NSNamespace("/clusterrole", beego.NSInclude(&controllers.ClusterRoleController{},),),
	//	//beego.NSNamespace("/clusterresource", beego.NSInclude(&controllers.AirAdController{},),),
	//)
	//beego.AddNamespace(ns)
}






