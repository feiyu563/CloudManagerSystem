package routers

import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/kubeCluster/datagrid", &controllers.KubeClusterController{}, "GET,POST:DataGrid")
	beego.Router("/kubeCluster/", &controllers.KubeClusterController{}, "POST:Save")
	beego.Router("/kubeCluster/relation", &controllers.KubeClusterController{}, "POST:Relation")
	//beego.Router("/kubeCluster/relation", &controllers.KubeClusterController{}, "GET:GetRelByCluster")
	beego.Router("/kubeCluster/authrelation", &controllers.KubeClusterController{}, "DELETE:DeleteAuthRelation")
	beego.Router("/kubeCluster/noderelation", &controllers.KubeClusterController{}, "DELETE:DeleteNodeRelation")

	beego.Router("/kubeCluster/", &controllers.KubeClusterController{}, "DELETE:Delete")
}
