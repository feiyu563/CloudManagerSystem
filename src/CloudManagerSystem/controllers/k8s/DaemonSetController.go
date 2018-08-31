package k8s

import (
	"CloudManagerSystem/controllers"
	"CloudManagerSystem/models/k8s"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"encoding/json"
	"fmt"
)

type DaemonSetController struct {
	controllers.BaseController
}

func (c DaemonSetController)GetALL(){
	daemonSetQueryParam := k8s.DaemonSetQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &daemonSetQueryParam)

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)

	if len(envUserCluster.ClusterId) == 0 {
		r := &models.JsonResult{enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), ""}
		c.Data["json"] = r
		return
	}

	clienthandle, err := models.GetApiServerHandle(envUserCluster.ClusterId, false)
	fmt.Println(err)
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		c.StopRun()
	}

	namespace := models.GetEnvNameSpaceByUserIdClusterId(u.Id, envUserCluster.ClusterId).Name
	if len(namespace) == 0 {
		//c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		r := &models.JsonResult{enums.JRCodeSucc, fmt.Sprintf("当前用户?命名空间环境变量为空", u.Id), ""}
		c.Data["json"] = r
		return
	}

	resultList,_ := k8s.GetDaemonSetList(clienthandle,namespace, &daemonSetQueryParam)
	c.Data["json"] = resultList
	c.ServeJSON()
}
