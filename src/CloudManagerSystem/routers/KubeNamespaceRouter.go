package routers

import (
	"github.com/astaxie/beego"
	"CloudManagerSystem/controllers"
)

func init(){

	//命名空间
	beego.Router("/kubeNameSpace/dataGrid",&controllers.KubeNameSpaceController{},"Get,Post:DataGrid")
	beego.Router("/kubeNameSpace",&controllers.KubeNameSpaceController{})

	//命名空间授权
	beego.Router("/kubeAuth/userNamespace/dataGrid",&controllers.KubeAuthUserNameSpaceController{},"Get,Post:DataGrid")
	beego.Router("/kubeAuth/userNamespace",&controllers.KubeAuthUserNameSpaceController{})

}