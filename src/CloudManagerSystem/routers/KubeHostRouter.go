package routers

import (
	"github.com/astaxie/beego"
	"CloudManagerSystem/controllers"
)

func init()  {
	//Host管理
	beego.Router("/kubeHost/datagrid", &controllers.KubeHostController{}, "Get,POST:DataGrid")
	beego.Router("/kubeHost/save", &controllers.KubeHostController{}, "Post:Save")
	beego.Router("/kubeHost/delete", &controllers.KubeHostController{}, "Post:Delete")
	beego.Router("/kubeHost/getall", &controllers.KubeHostController{}, "Get,POST:GetALL")

}
