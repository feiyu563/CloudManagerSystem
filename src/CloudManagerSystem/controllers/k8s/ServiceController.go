package k8s

import (
	"CloudManagerSystem/controllers"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"CloudManagerSystem/models/k8s"
	"encoding/json"
	"fmt"
)

type ServiceController struct {
	controllers.BaseController
}

func (c *ServiceController)GetALL(){
	serviceQueryParam :=k8s.ServiceQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &serviceQueryParam)
	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	//envUserCluster := models.KubeEnvUserCluster{ClusterId: "1"}
	//envUserCluster.ClusterId = "1"
	if len(envUserCluster.ClusterId) == 0 {
		//c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		r := &models.JsonResult{enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), ""}
		c.Data["json"] = r
		return
	}
	namespace := models.GetEnvNameSpaceByUserIdClusterId(u.Id,envUserCluster.ClusterId).Name
	if len(namespace) == 0 {
		//c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		r := &models.JsonResult{enums.JRCodeSucc, fmt.Sprintf("当前用户?命名空间环境变量为空",u.Id), ""}
		c.Data["json"] = r
		return
	}

	nodeList,error:= k8s.GetServiceList(envUserCluster.ClusterId,namespace,&serviceQueryParam)
	if error !=nil {
		r := &models.JsonResult{enums.JRCodeSucc, error.Error(), ""}
		c.Data["json"] = r
	}else{
		c.Data["json"] =nodeList
	}


	c.ServeJSON()
}


func (c *ServiceController) Get() {

	name := c.Ctx.Input.Query("name")
	if len(name)  ==0 {
		r := &models.JsonResult{enums.JRCodeSucc, "请传name参数", ""}
		c.Data["json"] = r
		return
	}

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)

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

	clienthandle, err := models.GetApiServerHandle(envUserCluster.ClusterId, false)

	if err != nil {
		fmt.Println(err)
		c.Data["json"] = err
		c.ServeJSON()
		c.StopRun()
	}

	data, err := k8s.GetServiceDetail(clienthandle, namespace,name)
	c.Data["json"] = data
	c.ServeJSON()
}
