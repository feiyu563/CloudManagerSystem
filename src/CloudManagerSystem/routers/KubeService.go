package routers

import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/kubeService/datagrid", &controllers.KubeServiceController{}, "GET,POST:DataGrid")
	beego.Router("/kubeService/", &controllers.KubeServiceController{}, "POST:Save")
	//beego.Router("/kubeService/version", &controllers.KubeServiceController{}, "POST:VersionPublish")
	//beego.Router("/kubeService/version", &controllers.KubeServiceController{}, "GET:VersionGetALL")
	beego.Router("/kubeService/scale", &controllers.KubeServiceController{}, "POST:Scale")

	beego.Router("/kubeService/", &controllers.KubeServiceController{}, "DELETE:Delete")

	beego.Router("/kubeService/pubroll", &controllers.KubeServiceController{}, "POST:PublishORRollback")

}

