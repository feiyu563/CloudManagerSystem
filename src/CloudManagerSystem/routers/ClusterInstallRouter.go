package routers
import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/k8s/installcluster", &controllers.InstallClusterController{}, "GET,POST:CreateBaseCluster")
	beego.Router("/k8s/installstorageclass", &controllers.InstallClusterController{}, "GET,POST:CreateStorageClass")
}