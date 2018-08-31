package routers

import (
	"github.com/astaxie/beego"
	"CloudManagerSystem/controllers"
)

func init()  {
	//kubernetes集群管理
	beego.Router("/kube/nodes", &controllers.KubeDashBoardController{}, "*:Nodes")
	beego.Router("/kube/deployment", &controllers.KubeDashBoardController{}, "*:Deployment")
	beego.Router("/kube/edit_deployment", &controllers.KubeDashBoardController{}, "*:Edit_Deployment")
	beego.Router("/kube/deployment_detail", &controllers.KubeDashBoardController{}, "*:Deployment_detail")
	beego.Router("/kube/pod", &controllers.KubeDashBoardController{}, "*:Pod")
	beego.Router("/kube/services", &controllers.KubeDashBoardController{}, "*:Services")
	beego.Router("/kube/services_detail", &controllers.KubeDashBoardController{}, "*:Services_detail")
	beego.Router("/kube/namespaces", &controllers.KubeDashBoardController{}, "*:Namespaces")
	beego.Router("/kube/pvs", &controllers.KubeDashBoardController{}, "*:Pvs")
	beego.Router("/kube/pvs_detail", &controllers.KubeDashBoardController{}, "*:Pvs_detail")
	beego.Router("/kube/storageclass", &controllers.KubeDashBoardController{}, "*:StorageClass")
	beego.Router("/kube/storageclass_detail", &controllers.KubeDashBoardController{}, "*:StorageClass_detail")
	beego.Router("/kube/createcephstorageclass", &controllers.KubeDashBoardController{}, "*:CreateCephStorageClass")
	beego.Router("/kube/statefulset", &controllers.KubeDashBoardController{}, "*:Statefulset")
	beego.Router("/kube/daemonset", &controllers.KubeDashBoardController{}, "*:DaemonSet")
	beego.Router("/kube/cronjob", &controllers.KubeDashBoardController{}, "*:CronJob")




	//部署服务
	beego.Router("/kube/deploy", &controllers.KubeController{}, "*:Deploy")
	beego.Router("/kube/editdeploy", &controllers.KubeController{}, "*:EditDeploy")
	beego.Router("/kube/publish", &controllers.KubeController{}, "*:Publish")
	beego.Router("/kube/history", &controllers.KubeController{}, "*:History")
	beego.Router("/kube/manage", &controllers.KubeController{}, "*:Manage")
	beego.Router("/kube/scale", &controllers.KubeController{}, "*:Scale")

	//服务发布
	beego.Router("/kube/deployproxy", &controllers.KubeController{}, "*:DeployProxy")
	beego.Router("/kube/editdeployproxy", &controllers.KubeController{}, "*:EditDeployProxy")

	beego.Router("/kube/deployIngress", &controllers.KubeController{}, "*:DeployIngress")
	beego.Router("/kube/editdeployingress", &controllers.KubeController{}, "*:EditDeployIngress")
	//Images镜像管理
	beego.Router("/kube/images", &controllers.KubeController{}, "*:Images")
	beego.Router("/kube/repositories", &controllers.KubeController{}, "*:Repositories")
	beego.Router("/kube/tags", &controllers.KubeController{}, "*:Tags")
	beego.Router("/kube/editimages", &controllers.KubeController{}, "*:EditImages")

	//Logs管理
	beego.Router("/kube/logs", &controllers.MessagesController{}, "*:Logs")
	beego.Router("/kube/getlogs", &controllers.MessagesController{}, "*:GetLogs")

	//Appshop
	beego.Router("/kube/apps", &controllers.AppshopController{}, "*:Apps")

	//首页图表加载
	beego.Router("/kube/procpumemdisk", &controllers.KubePrometheusController{}, "*:ProCpuMemDisk")
	beego.Router("/kube/procount", &controllers.KubePrometheusController{}, "*:ProCount")
}

