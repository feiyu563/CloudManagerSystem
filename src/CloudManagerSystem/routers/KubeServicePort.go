package routers

import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//beego.Router("/kubeService/datagrid", &controllers.KubeServicePortController{}, "GET,POST:DataGrid")
	beego.Router("/kubeServicePort/", &controllers.KubeServicePortController{}, "Get:ServciePort")

	beego.Router("/bkubeServicePort/", &controllers.KubeServicePortController{}, "DELETE:Delete")
}
