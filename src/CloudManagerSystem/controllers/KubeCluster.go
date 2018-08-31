package controllers

import (
	"encoding/json"
	"fmt"
	"strings"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
)

type KubeClusterController struct {
	BaseController
}

func (c *KubeClusterController) DataGrid() {
	kubeCluster := models.KubeClusterQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeCluster)
	data, total := models.GetAllKubeCluster(&kubeCluster)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *KubeClusterController) Save() {
	var err error
	u := c.GetSessionUser()

	m := models.KubeCluster{}
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	o := orm.NewOrm()
	var title string
	if strings.Compare(m.Id, "") == 0 {
		u4, _ := uuid.NewV4()
		m.Id = u4.String()
		title = "添加"
		m.CreateUser = u.UserName
		_, err = o.Insert(&m)
	} else {
		title = "编辑"
		aa := &models.KubeCluster{Id: m.Id}
		o.Read(aa)
		//not update
		m.CreateUser = aa.CreateUser
		m.CreateTime = aa.CreateTime

		_, err = o.Update(&m)
	}
	if err == nil {
		c.jsonResult(enums.JRCodeSucc, title+"成功", m.Id)
	} else {
		c.jsonResult(enums.JRCodeFailed, title+"失败", m.Id)
	}

}

func (c *KubeClusterController)DeleteNodeRelation() {
	var err error
	kubeCluster := models.JSONKubeClusterRelation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeCluster)
	fmt.Println(kubeCluster)
	//if strings.Compare(kubeCluster.ClusterId, "") == 0 {
	//	c.jsonResult(enums.JRCodeSucc, "Del成功", kubeCluster)
	//}
	ClusterParm := kubeCluster.KubeClusterQueryParam
	for k, _ := range ClusterParm {
		if strings.Compare((*ClusterParm[k]).Id, "") == 0 {
			c.jsonResult(enums.JRCodeSucc, "Del成功", kubeCluster)
		}
		err = models.KubeClusterHostRelationDel((*ClusterParm[k]).Id)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "DELETE Auth Relation ERROR", kubeCluster)
		}
	}
	c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除"), kubeCluster)
}

func (c *KubeClusterController)DeleteAuthRelation(){
	var err error
	kubeCluster := models.JSONKubeClusterRelation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeCluster)
	fmt.Println(kubeCluster)
	//if strings.Compare(kubeCluster.ClusterId, "") == 0 {
	//	c.jsonResult(enums.JRCodeSucc, "Del成功", kubeCluster)
	//}
	ClusterParm := kubeCluster.KubeClusterQueryParam
	for k, _ := range ClusterParm {
		if strings.Compare((*ClusterParm[k]).Id, "") == 0 {
			c.jsonResult(enums.JRCodeSucc, "Del成功", kubeCluster)
		}
		err = models.KubeClusterAuthRelationDel((*ClusterParm[k]).Id)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "DELETE Auth Relation ERROR", kubeCluster)
		}
	}

	c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除"), kubeCluster)
}

func (c *KubeClusterController) Relation() {
	var err error

	kubeCluster := models.JSONKubeClusterRelation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeCluster)
	fmt.Println(kubeCluster)
	if strings.Compare(kubeCluster.ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeFailed, "分配失败", kubeCluster)
	}

	ClusterParm := kubeCluster.KubeClusterQueryParam
	if kubeCluster.Type == models.USERTYPE {
		//err = models.KubeClusterNodeRelationDel(kubeCluster.ClusterId)
		//if err != nil {
		//	c.jsonResult(enums.JRCodeFailed, "DELETE Relation ERROR", kubeCluster)
		//}
		for k, _ := range ClusterParm {
			err = models.KubeClusterAuthRelation(*ClusterParm[k], kubeCluster.ClusterId)
			if err != nil {
				c.jsonResult(enums.JRCodeFailed, fmt.Sprintf("Auth %s Relation ERROR",(*ClusterParm[k]).Name), kubeCluster)
			}
		}
	} else if kubeCluster.Type == models.HOSTTYPE {
		err = models.KubeClusterNodeRelationDel(kubeCluster.ClusterId)
		if err != nil {
			c.jsonResult(enums.JRCodeFailed, "DELETE Relation ERROR", kubeCluster)
		}

		for k, _ := range ClusterParm {
			err = models.KubeClusterNodeRelation(*ClusterParm[k], kubeCluster.ClusterId)
			if err != nil {
				c.jsonResult(enums.JRCodeFailed, fmt.Sprintf("Host %s Relation ERROR",(*ClusterParm[k]).Name), kubeCluster)
			}
		}
	} else {
		c.jsonResult(enums.JRCodeFailed, "Function ERROR", kubeCluster)
	}
	c.jsonResult(enums.JRCodeSucc, "分配成功", kubeCluster)

}


func (c *KubeClusterController) Delete() {

	strs := c.GetString("ids")
	ids := strings.Split(strs, ",")

	if num, err := models.KubeClusterDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}
