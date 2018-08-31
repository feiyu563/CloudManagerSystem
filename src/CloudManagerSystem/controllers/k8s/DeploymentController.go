package k8s

import (
	"encoding/json"
	"fmt"
	//"strings"

	"CloudManagerSystem/models"
	"CloudManagerSystem/controllers"
	"CloudManagerSystem/models/k8s"
	"CloudManagerSystem/enums"
)

type DeploymentController struct {
	controllers.BaseController
}

func (c *DeploymentController) GetALL() {
	fmt.Println("GetALL")
	deploymentquery := k8s.DeploymentQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &deploymentquery)

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)
	//envUserCluster.ClusterId ="1"
	if len(envUserCluster.ClusterId) == 0 {
		r := &models.JsonResult{enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), ""}
		c.Data["json"] = r
		c.ServeJSON()
		c.StopRun()
	}
	clienthandle, err := models.GetApiServerHandle(envUserCluster.ClusterId, false)
	fmt.Println(err)
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		c.StopRun()
	}
	namespace := models.GetEnvNameSpaceByUserIdClusterId(u.Id,envUserCluster.ClusterId).Name

	data := k8s.GetDeploymentList(&deploymentquery, clienthandle, namespace)

	c.Data["json"] = data
	c.ServeJSON()
}

func (c *DeploymentController) GetDetail() {
	fmt.Println("GetALL")
	deploymentquery := k8s.DeploymentDetailQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &deploymentquery)

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)
	if len(envUserCluster.ClusterId) == 0 {
		r := &models.JsonResult{enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), ""}
		c.Data["json"] = r
		c.ServeJSON()
		c.StopRun()
	}
	clienthandle, err := models.GetApiServerHandle(envUserCluster.ClusterId, false)
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		c.StopRun()
	}
	namespace := models.GetEnvNameSpaceByUserIdClusterId(u.Id,envUserCluster.ClusterId).Name

	//data, err := k8s.GetDeploymentDetail(&deploymentquery, clienthandle, "default", "bikemaintainserviceweb") // deploymentquery.Name)
	data, err := k8s.GetDeploymentDetail(&deploymentquery, clienthandle, namespace, deploymentquery.Name) // deploymentquery.Name)

	c.Data["json"] = data
	c.ServeJSON()
}
