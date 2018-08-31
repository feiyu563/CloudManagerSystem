package routers

import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/kubeBind/dataGrid", &controllers.KubeBindController{}, "GET,POST:DataGrid")
}