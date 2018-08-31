package controllers



import (
	"encoding/json"
	"strings"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
)

type KubeUserAuthClusterController struct {
	BaseController
}

func (c *KubeUserAuthClusterController) DataGrid() {
	var ClusterId string
	c.Ctx.Input.Bind(&ClusterId, "ClusterId")

	if strings.Compare(ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeFailed, "ClusterId NULL", "")
	}


	kubeCluster := models.KubeClusterQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeCluster)
	data, total := models.GetAllKubeAuthUser(&kubeCluster,ClusterId)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
