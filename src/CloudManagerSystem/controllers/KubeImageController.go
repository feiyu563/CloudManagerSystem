package controllers
import (
	"encoding/json"
	//"fmt"
	//"strings"
	"CloudManagerSystem/models"
	//"CloudManagerSystem/enums"
	//"github.com/astaxie/beego/orm"
	//"github.com/satori/go.uuid"
	//"fmt"
	"CloudManagerSystem/enums"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/satori/go.uuid"
)

type KubeImageController struct {
	BaseController
}

func (c *KubeImageController) DataGrid() {
	kubeservice := models.KubeImageQueryParam{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &kubeservice)
	data, total := models.KubeImagePageList(&kubeservice)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *KubeImageController) GetALL(){

	data, total := models.KubeImageList()

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}


//获取一个
func(c *KubeImageController) Get(){
	strs := c.GetString("id")
	if(len(strs) == 0){
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("id 不能为空"), 0)
	}
	c.Ctx.Input.RequestBody =[]byte("{\"Id\":\""+strs+"\"}")
	c.DataGrid()
}

//添加
func(c *KubeImageController)Post(){
	c.Save()
}

//修改
func(c *KubeImageController)Put(){
	c.Save()
}

//删除
func(c *KubeImageController)Delete(){
	strs := c.GetString("ids")
	ids := make([]string, 0, len(strs))
	for _, id := range strings.Split(strs, ",") {
		ids = append(ids, id)
	}
	if num, err := models.KubeImageDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

//保存方法 (实现 添加与修改功能)
func (c *KubeImageController) Save(){
	var err error
	m := models.KubeImage{}
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
		return
	}

	u:=c.GetSessionUser()
	envUserCluster,_ :=models.EnvUserCluster(u.Id)
	if(len(envUserCluster.ClusterId) ==0){
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前用户?集群环境变量为空",u.Id), 0)
		return
	}

	o := orm.NewOrm()
	var title string

	if(strings.Compare(m.Id,"")==0){
		u4,_:= uuid.NewV4()
		m.Id=u4.String()
		title="添加"
		_, err = o.Insert(&m)
	}else{
		title="编辑"
		_, err = o.Update(&m)
	}
	if err == nil {
		c.jsonResult(enums.JRCodeSucc, title+"成功", m.Id)
	} else {
		c.jsonResult(enums.JRCodeFailed, title+"失败", m.Id)
	}
}