package routers

import (
	"github.com/astaxie/beego"
	"CloudManagerSystem/controllers"
)

func init()  {
	//Basecluster
	beego.Router("/basecluster/hosts", &controllers.BaseclusterController{}, "*:Hosts")
	beego.Router("/basecluster/edit?:id", &controllers.BaseclusterController{}, "Get,Post:Edit")
	beego.Router("/basecluster/cluster", &controllers.BaseclusterController{}, "*:Cluster")
	beego.Router("/basecluster/save", &controllers.BaseclusterController{}, "*:Saveconfig")
	beego.Router("/basecluster/savecfg", &controllers.BaseclusterController{}, "*:Savecfg")
	beego.Router("/basecluster/editcluster?:id", &controllers.BaseclusterController{}, "Get,Post:EditCluster")
	beego.Router("/basecluster/allocationnode", &controllers.BaseclusterController{}, "*:AllocationNode")
	beego.Router("/basecluster/allocationrole", &controllers.BaseclusterController{}, "*:AllocationRole")
	beego.Router("/basecluster/clustersetup", &controllers.BaseclusterController{}, "*:ClusterSetup")
}