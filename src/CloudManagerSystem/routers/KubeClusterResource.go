package routers

import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/base/clusterresource/dataGrid", &controllers.ClusterResourceController{}, "GET,POST:GetALL")
	//beego.Router("/base/clusterresource/:rolename([\\w]+)", &controllers.ClusterResourceController{}, "GET:GetByRoleName")
	beego.Router("/base/clusterresource/rsgetall", &controllers.ClusterResourceController{}, "GET:RsGetALL")
	beego.Router("/base/clusterresource/opgetall", &controllers.ClusterResourceController{}, "GET:OpGetALL")

	////beego.Router("/base/clusterrole", &controllers.ClusterRoleController{}, "POST:ClusterRoleCreate")
	beego.Router("/base/clusterresource", &controllers.ClusterResourceController{}, "POST,PUT:Update")
	beego.Router("/base/clusterresource", &controllers.ClusterResourceController{}, "DELETE:Delete")

	//ns := beego.NewNamespace("/base",
	//	beego.NSNamespace("/clusterrole", beego.NSInclude(&controllers.ClusterRoleController{},),),
	//	//beego.NSNamespace("/clusterresource", beego.NSInclude(&controllers.AirAdController{},),),
	//)
	//beego.AddNamespace(ns)
}
