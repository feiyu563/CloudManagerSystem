package routers

import (
	"github.com/astaxie/beego"
	"CloudManagerSystem/controllers"
)

func init()  {
	beego.Router("/kubeUserGroup/datagrid", &controllers.KubeUserGroupController{}, "Get,POST:DataGrid")
	beego.Router("/kubeUserGroup/save", &controllers.KubeUserGroupController{}, "Post:Save")
	beego.Router("/kubeUserGroup/delete", &controllers.KubeUserGroupController{}, "Post:Delete")
	beego.Router("/kubeUserGroup/allList", &controllers.KubeUserGroupController{}, "Get,POST:AllList")
}