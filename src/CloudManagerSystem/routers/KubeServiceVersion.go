package routers

import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/kubeService/version", &controllers.KubeServiceVersionController{}, "POST:Publish")
	beego.Router("/kubeService/version/datagrid", &controllers.KubeServiceVersionController{}, "GET,POST:DataGrid")
	//beego.Router("/kubeService/version", &controllers.KubeServiceVersionController{}, "DELETE:Delete")
}
