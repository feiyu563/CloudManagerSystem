package k8s

import (
	"encoding/json"
	"fmt"
	//"strings"

	//"k8s.io/apimachinery/pkg/runtime"

	"CloudManagerSystem/enums"
	"CloudManagerSystem/models"
	"CloudManagerSystem/controllers"
	"CloudManagerSystem/models/k8s"
)

type RawResourceController struct {
	controllers.BaseController
}

func (c *RawResourceController) GetRaw() {
	fmt.Println("GetALL")
	deploymentquery := k8s.ResourceVerberQueryParam{}
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
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		c.StopRun()
	}
	resourceVerberHandle := k8s.VerberClientHandle(clienthandle)

	data, _ := resourceVerberHandle.Get(deploymentquery.Kind, true, deploymentquery.Namespace, deploymentquery.Name)
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *RawResourceController) PutRaw() {
	fmt.Println("GetALL")
	deploymentquery := k8s.ResourceVerberQueryParam{}
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
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		c.StopRun()
	}
	//fmt.Printf("%+v\n",deploymentquery)
	//fmt.Println("###################################",string(deploymentquery.PutSpec.Raw))
	resourceVerberHandle := k8s.VerberClientHandle(clienthandle)

	//putSpec := &runtime.Unknown{}
	//json.Unmarshal(deploymentquery.PutSpec, putSpec)

	err = resourceVerberHandle.Put(deploymentquery.Kind, true, deploymentquery.Namespace, deploymentquery.Name, deploymentquery.PutSpec)
	//err = resourceVerberHandle.Put("deployment", true, "default", "bikemaintainserviceweb", putSpec)
	if err != nil {

		r := &models.JsonResult{enums.JRCodeSucc, err.Error(), ""}
		c.Data["json"] = r
		c.ServeJSON()
	} else {
		r := &models.JsonResult{enums.JRCodeSucc, "OK", ""}
		c.Data["json"] = r
		c.ServeJSON()
	}

}

func (c *RawResourceController) DeleteRaw() {
	fmt.Println("GetALL")
	deploymentquery := k8s.ResourceVerberQueryParam{}
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
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		c.StopRun()
	}
	resourceVerberHandle := k8s.VerberClientHandle(clienthandle)

	err = resourceVerberHandle.Delete(deploymentquery.Kind, true, deploymentquery.Namespace, deploymentquery.Name)
	if err != nil {
		r := &models.JsonResult{enums.JRCodeFailed, err.Error(), ""}
		c.Data["json"] = r
		c.ServeJSON()
	} else {
		r := &models.JsonResult{enums.JRCodeSucc, "OK", ""}
		c.Data["json"] = r
		c.ServeJSON()
	}

}
