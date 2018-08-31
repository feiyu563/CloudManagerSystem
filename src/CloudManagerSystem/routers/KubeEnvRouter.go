package routers

import (
	"github.com/astaxie/beego"
	"CloudManagerSystem/controllers"
)

func init(){
	//环境变量
	beego.Router("/KubeEnvUserCluster",&controllers.KubeEnvUserClusterController{})
	beego.Router("/KubeEnvUserNameSpace",&controllers.KubeEnvUserNamespaceController{})

}