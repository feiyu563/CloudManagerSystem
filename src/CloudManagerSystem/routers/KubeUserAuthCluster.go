package routers

import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/kubeauth/datagrid", &controllers.KubeUserAuthClusterController{}, "GET,POST:DataGrid")
}
