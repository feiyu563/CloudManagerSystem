package controllers

import (
	"encoding/json"
	"fmt"
	"strings"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego/orm"
)

type KubeServiceVersionController struct {
	BaseController
}

func (c *KubeServiceVersionController) Publish() {

	jsonkubeserviceversion := models.JSONKubeServiceVersion{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &jsonkubeserviceversion)

	var ClusterId string

	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)

	if strings.Compare(envUserCluster.ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
	} else {
		ClusterId = envUserCluster.ClusterId
	}

	if strings.Compare(jsonkubeserviceversion.Id, "") == 0 {
		c.jsonResult(enums.JRCodeFailed, "版本发布失败", jsonkubeserviceversion)
	}
	fmt.Println("---------------------", jsonkubeserviceversion)

	var err error
	kubeservice := models.KubeService{}
	kubeserviceversion := models.KubeServiceVersion{}
	o := orm.NewOrm()

	//
	u3, _ := uuid.NewV4()
	kubeserviceversion.Id = u3.String()
	kubeserviceversion.CreateUser = u.UserName
	kubeserviceversion.VersionName = jsonkubeserviceversion.VersionName
	kubeserviceversion.Remark = jsonkubeserviceversion.VersionRemark

	_, err = o.Insert(&kubeserviceversion)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "版本发布失败", err)
	}

	//
	aa := &models.KubeService{Id: jsonkubeserviceversion.Id}
	o.Read(aa)
	u4, _ := uuid.NewV4()
	kubeservice = *aa
	kubeservice.Id = u4.String()
	kubeservice.IsVersion = 1
	kubeservice.FatherId = jsonkubeserviceversion.Id
	kubeservice.VersionId = kubeserviceversion.Id
	kubeservice.ClusterId = ClusterId

	_, err = o.Insert(&kubeservice)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "版本发布失败", err)
	} else {
		c.jsonResult(enums.JRCodeSucc, "版本发布成功", kubeservice)
	}
}

func (c *KubeServiceVersionController) DataGrid() {
	kubeservice := models.KubeServiceVersionQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeservice)
	//var ClusterId string
	u := c.GetSessionUser()
	envUserCluster, _ := models.EnvUserCluster(u.Id)
	//fmt.Println(u.UserName)
	if strings.Compare(envUserCluster.ClusterId, "") == 0 {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空", u.Id), 0)
	} else {
		//ClusterId = envUserCluster.ClusterId
	}
	if strings.Compare(kubeservice.Id,"") == 0 {
		c.jsonResult(enums.JRCodeFailed, "error", kubeservice)
	}
	data, total := models.KubeServiceVersionPageList(&kubeservice)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
