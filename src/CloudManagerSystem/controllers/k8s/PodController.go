package k8s

import (
	"CloudManagerSystem/controllers"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"CloudManagerSystem/models/k8s"
	"encoding/json"
	"fmt"
)

type PodController struct {
	controllers.BaseController
}

func (c *PodController) GetALL() {

	podQueryParam := k8s.PodQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &podQueryParam)

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)
	//envUserCluster := models.KubeEnvUserCluster{ClusterId: "1"}
	//envUserCluster.ClusterId = "1"
	if len(envUserCluster.ClusterId) == 0 {
		//c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		r := &models.JsonResult{enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), ""}
		c.Data["json"] = r
		return
	}
	namespace := models.GetEnvNameSpaceByUserIdClusterId(u.Id, envUserCluster.ClusterId).Name
	if len(namespace) == 0 {
		//c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		r := &models.JsonResult{enums.JRCodeSucc, fmt.Sprintf("当前用户?命名空间环境变量为空", u.Id), ""}
		c.Data["json"] = r
		return
	}

	podList, error := k8s.GetPodList(envUserCluster.ClusterId, namespace, &podQueryParam)
	if error != nil {
		r := &models.JsonResult{enums.JRCodeSucc, error.Error(), ""}
		c.Data["json"] = r
	} else {
		c.Data["json"] = podList
	}

	c.ServeJSON()
}
