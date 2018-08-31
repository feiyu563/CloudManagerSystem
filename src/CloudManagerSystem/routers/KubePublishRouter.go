package routers

import (
	"github.com/astaxie/beego"
	"CloudManagerSystem/controllers"
)


func init(){
	//publish ingress
	beego.Router("/KubePublishService/dataGrid",&controllers.KubePublishServiceController{},"Get,Post:DataGrid")
	beego.Router("/KubePublishService",&controllers.KubePublishServiceController{})
	beego.Router("/KubePublishServicePath/dataGrid",&controllers.KubePublishServicePathController{},"Get,Post:DataGrid")
	//publish proxy
	/*beego.Router("/KubePublishProxy/dataGrid",&controllers.KubePublishProxyController{},"Get,Post:DataGrid")
	beego.Router("/KubePublishProxy",&controllers.KubePublishProxyController{})*/
}