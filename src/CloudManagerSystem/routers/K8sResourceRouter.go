package routers

import (
	"CloudManagerSystem/controllers/k8s"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/k8s/deployment", &k8s.DeploymentController{}, "GET,POST:GetALL")
	beego.Router("/k8s/deploymentdetail", &k8s.DeploymentController{}, "GET,POST:GetDetail")

	beego.Router("/k8s/raw", &k8s.RawResourceController{}, "GET,POST:GetRaw")
	beego.Router("/k8s/raw/update", &k8s.RawResourceController{}, "POST:PutRaw")
	beego.Router("/k8s/raw", &k8s.RawResourceController{}, "DELETE:DeleteRaw")


	beego.Router("/k8s/namespaces", &k8s.NamespaceController{}, "GET,POST:GetALL")
	beego.Router("/k8s/nodes", &k8s.NodeController{}, "GET,POST:GetALL")
	beego.Router("/k8s/pods", &k8s.PodController{}, "GET,POST:GetALL")
	beego.Router("/k8s/services", &k8s.ServiceController{}, "GET,POST:GetALL")
	beego.Router("/k8s/service", &k8s.ServiceController{})

	beego.Router("/k8s/persistentvolumes", &k8s.PersistentVolumeController{}, "GET,POST:GetALL")
	beego.Router("/k8s/persistentvolume", &k8s.PersistentVolumeController{})

	beego.Router("/k8s/storageclasss", &k8s.StorageClassController{}, "GET,POST:GetALL")
	beego.Router("/k8s/storageclass", &k8s.StorageClassController{})

	beego.Router("/k8s/statefulsets", &k8s.StatefulSetController{}, "GET,POST:GetALL")
	//beego.Router("/k8s/statefulset", &k8s.StatefulSetController{})

	beego.Router("/k8s/daemonsets", &k8s.DaemonSetController{}, "GET,POST:GetALL")
	//beego.Router("/k8s/statefulset", &k8s.StatefulSetController{})

	beego.Router("/k8s/cronjobs", &k8s.CronJobController{}, "GET,POST:GetALL")
	//beego.Router("/k8s/statefulset", &k8s.StatefulSetController{})




}

