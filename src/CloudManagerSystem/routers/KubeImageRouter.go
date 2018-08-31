package routers

import (
	"CloudManagerSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/kubeImage", &controllers.KubeImageController{})
	beego.Router("/kubeImage/datagrid", &controllers.KubeImageController{}, "GET,POST:DataGrid")
	beego.Router("/kubeImage/getall", &controllers.KubeImageController{}, "GET,POST:GetALL")
	beego.Router("/kubeImage/projects", &controllers.KubeImagesController{}, "GET,POST:Projects")
	beego.Router("/kubeImage/repositories", &controllers.KubeImagesController{}, "GET,POST:Repositories")
	beego.Router("/kubeImage/tags", &controllers.KubeImagesController{}, "GET,POST:Tags")
}
