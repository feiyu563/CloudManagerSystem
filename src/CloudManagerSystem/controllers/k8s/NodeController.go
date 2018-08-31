package k8s


import (

	//"strings"
	"CloudManagerSystem/controllers"
	"CloudManagerSystem/models/k8s"
	"CloudManagerSystem/enums"
	"CloudManagerSystem/models"
	"encoding/json"
)

type NodeController struct {
	controllers.BaseController
}

func (c *NodeController)GetALL(){

	nodeQueryParam :=k8s.NodeQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &nodeQueryParam)

	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	//envUserCluster :=  models.KubeEnvUserCluster{ ClusterId:"1" }
	//envUserCluster.ClusterId ="1"
	if len(envUserCluster.ClusterId) ==0 {
		//c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		r := &models.JsonResult{enums.JRCodeSucc, "当前用户?集群环境变量为空", ""}
		c.Data["json"] = r
		return
	}

	nodeList,error:= k8s.GetNodeList(envUserCluster.ClusterId,&nodeQueryParam)
	if error !=nil {
		r := &models.JsonResult{enums.JRCodeSucc, error.Error(), ""}
		c.Data["json"] = r
	}else{
		c.Data["json"] =nodeList
	}


	c.ServeJSON()
}
