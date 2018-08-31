package k8s

import (
	"CloudManagerSystem/controllers"
	"CloudManagerSystem/models/k8s"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"encoding/json"
	"fmt"
)

type PersistentVolumeController struct {
	controllers.BaseController
}

func (c *PersistentVolumeController) GetALL() {

	persistentQueryParam := k8s.PersistentVolumeQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &persistentQueryParam)

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)
	//envUserCluster :=  models.KubeEnvUserCluster{ ClusterId:"1" }
	//envUserCluster.ClusterId ="1"
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

	resultList := k8s.GetPersistentVolumeList(clienthandle, &persistentQueryParam)
	c.Data["json"] = resultList
	c.ServeJSON()
}

func (c *PersistentVolumeController) Get() {

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

	clienthandle, err := models.GetApiServerHandle(envUserCluster.ClusterId, false)

	if err != nil {
		fmt.Println(err)
		c.Data["json"] = err
		c.ServeJSON()
		c.StopRun()
	}

	data, err := k8s.GetPersistentVolumeDetail(clienthandle, name)
	c.Data["json"] = data
	c.ServeJSON()
}
