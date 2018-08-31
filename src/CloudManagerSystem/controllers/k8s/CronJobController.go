package k8s

import (
	"CloudManagerSystem/controllers"
	"CloudManagerSystem/models/k8s"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"fmt"
	"encoding/json"
)

type CronJobController struct {
	controllers.BaseController
}

func (c *CronJobController) GetALL(){
	cronJobQueryParam := k8s.CronJobQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &cronJobQueryParam)

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

	resultList,_ := k8s.GetCronJobList(clienthandle,namespace, &cronJobQueryParam)
	c.Data["json"] = resultList
	c.ServeJSON()
}
