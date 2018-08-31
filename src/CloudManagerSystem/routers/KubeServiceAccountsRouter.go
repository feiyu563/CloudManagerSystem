package routers

import (
	"github.com/astaxie/beego"
	"CloudManagerSystem/controllers"
)

func init()  {
	beego.Router("/kubeServiceAccountsRouter/datagrid", &controllers.KubeServiceAccountsController{}, "Get,POST:DataGrid")
	beego.Router("/kubeServiceAccountsRouter/save", &controllers.KubeServiceAccountsController{}, "Post:Save")
	beego.Router("/kubeServiceAccountsRouter/delete", &controllers.KubeServiceAccountsController{}, "Post:Delete")
}