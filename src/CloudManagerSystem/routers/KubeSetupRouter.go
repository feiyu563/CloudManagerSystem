package routers

import (
	"github.com/astaxie/beego"
	"CloudManagerSystem/controllers"
)

func init()  {
	//BaseclusterSetupController.
	beego.Router("/baseclustersetup/namespacesetup", &controllers.BaseclusterSetupController{}, "*:NameSpaceSetup")

	beego.Router("/baseclustersetup/editnamespace", &controllers.BaseclusterSetupController{}, "*:EditNamespace")
	beego.Router("/baseclustersetup/grantnamespace", &controllers.BaseclusterSetupController{}, "*:GrantNamespace")

	beego.Router("/baseclustersetup/clusterrole", &controllers.BaseclusterSetupController{}, "*:ClusterRole")
	beego.Router("/basecluster/editclusterrole", &controllers.BaseclusterSetupController{}, "*:EditClusterRole")

	beego.Router("/baseclustersetup/grantusergroup", &controllers.BaseclusterSetupController{}, "*:GrantUserGroup")
	beego.Router("/baseclustersetup/editclustergrantusergroup", &controllers.BaseclusterSetupController{}, "*:EditClusterGrantUserGroup")


	beego.Router("/baseclustersetup/sa", &controllers.BaseclusterSetupController{}, "*:Sa")
	beego.Router("/baseclustersetup/editsa", &controllers.BaseclusterSetupController{}, "*:EditSa")
	beego.Router("/baseclustersetup/new", &controllers.BaseclusterSetupController{}, "*:New")

}