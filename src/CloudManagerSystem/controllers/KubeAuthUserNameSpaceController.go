package controllers

import (
	"strings"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"encoding/json"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
)

type KubeAuthUserNameSpaceController struct {
	BaseController
}

//分页展示NameSpace
func (c *KubeAuthUserNameSpaceController) DataGrid() {
	queryParam := models.KubeAuthUserNameSpaceQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &queryParam)
	data, total := models.KubeAuthUserNameSpacePageList(&queryParam)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//获取一个
func (c *KubeAuthUserNameSpaceController) Get() {
	strs := c.GetString("id")
	if (len(strs) == 0) {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("id 不能为空"), 0)
	}
	c.Ctx.Input.RequestBody = []byte("{\"Id\":\"" + strs + "\"}")
	c.DataGrid()
}

//添加
func (c *KubeAuthUserNameSpaceController) Post() {
	c.Save()
}

//修改
func (c *KubeAuthUserNameSpaceController) Put() {
	c.Save()
}

//删除
func (c *KubeAuthUserNameSpaceController) Delete() {
	strs := c.GetString("ids")
	ids := make([]string, 0, len(strs))
	for _, id := range strings.Split(strs, ",") {
		ids = append(ids, id)
	}
	if num, err := models.KubeAuthUserNameSpaceDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

//保存方法 (实现 添加与修改功能)
func (c *KubeAuthUserNameSpaceController) Save() {
	var err error
	//var tt map[string][]models.JSONKubeAuthUserNameSpace
	getbodyData := models.JSONKubeAuthUserNameSpace{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &getbodyData)

	fmt.Printf("%+v", getbodyData)

	//得到NameSpace授权用户集合
	if (len(getbodyData.NameSpacesAuthUser) < 1) {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", getbodyData)
		return
	}

	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	if(len(envUserCluster.ClusterId) ==0){
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		return
	}

	var insertNum, updateNum, errNun int64
	updateIds := make(map[string]string)
	o := orm.NewOrm()
	for _, p := range getbodyData.NameSpacesAuthUser {

		m := models.KubeAuthUserNameSpace{Id: p.Id, UserId: p.UserId, UserType: p.UserType, NamespaceId: getbodyData.NamespaceId, ClusterId: getbodyData.ClusterId,CreateUser:u.Id}

		if (strings.Compare(m.Id, "") == 0) {
			//m.Stype=1
			u4, _ := uuid.NewV4()
			m.Id = u4.String()

			_, err = o.Insert(&m)
			if err == nil {
				insertNum += 1
			} else {
				errNun += 1
			}
		} else {
			_, err = o.Update(&m)
			if err == nil {
				updateNum += 1
			} else {
				errNun += 1
			}
		}

		getId := fmt.Sprintf("'%s'", m.Id)
		if _, ok := updateIds[m.NamespaceId]; !ok {
			updateIds[m.NamespaceId] = getId
		} else {
			updateIds[m.NamespaceId] = (updateIds[m.NamespaceId] + "," + getId)
		}
	}

	//再次编辑删除去除的项
	models.KubeAuthUserNameSpaceDeleteByNsIdNotInId(&updateIds)
	//errNun += deleErr

	// fmt.Sprintf("保存成功", 0)
	resultIds := strings.Replace( updateIds[getbodyData.NamespaceId],"'","",-1)
	c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("保存成功"), resultIds)
	//c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("新增%d项，修改%d项，删除%d项，失败%d项", insertNum, updateNum, deleteNum, errNun), 0)
}
