package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"CloudManagerSystem/enums"
	"CloudManagerSystem/models"
	"CloudManagerSystem/utils"

	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"github.com/astaxie/beego"
)

type BackendUserController struct {
	BaseController
}

func (c *BackendUserController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()

}
func (c *BackendUserController) Index() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backenduser/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "backenduser/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("BackendUserController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("BackendUserController", "Delete")
}

func (c *BackendUserController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.BackendUserQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.BackendUserPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *BackendUserController) AllList() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.BackendUserQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, _ := models.BackendUserAllList(&params)
	//定义返回的数据结构
	c.Data["json"] = data
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *BackendUserController) Edit() {
	Id := c.GetString("id")
	m := &models.BackendUser{}
	var err error
	if strings.Compare(Id, "") != 0 {
		m, err = models.BackendUserOne(Id)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
		o := orm.NewOrm()
		o.LoadRelated(m, "RoleBackendUserRel")
		//获取关联的roleId列表
		var roleIds []string
		for _, item := range m.RoleBackendUserRel {
			roleIds = append(roleIds, strconv.Itoa(item.Role.Id))
		}
		c.Data["roles"] = strings.Join(roleIds, ",")

		list, _ := models.KubeUser2GroupAllList(&models.KubeUser2GroupQueryParam{UserId: Id})
		var groupIds []string
		for _, item := range list {
			groupIds = append(groupIds, item.GroupId)
		}
		c.Data["groupIds"] = strings.Join(groupIds, ",")
	} else {
		//添加用户时默认状态为启用
		m.Status = enums.Enabled
		m.UserType = -1
	}
	c.Data["m"] = m
	c.setTpl("backenduser/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "backenduser/edit_footerjs.html"
}
func (c *BackendUserController) Save() {
	m := models.BackendUser{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}

	var apiuser models.UserVo
	if strings.Compare(m.Id, "") == 0 {
		u4, _ := uuid.NewV4()
		m.Id = u4.String()
		//对密码进行加密
		apiuser.PassWord = m.UserPwd

		m.UserPwd = utils.String2md5(m.UserPwd)
		if _, err := o.Insert(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
		apiuser.Oper = "ADD"
		apiuser.UID = m.Id
		apiuser.Name = m.UserName
		for _, groupId := range m.GroupIds {
			data := models.KubeUserGroup{}

			query := orm.NewOrm().QueryTable(models.KubeUserGroupTBName())
			query.Filter("id", groupId).One(&data)
			apiuser.Groups = data.GroupName + "," + apiuser.Groups
		}
		fmt.Println("######################add##########")
		//models.SendUserModifyMsgToApiserver(&apiuser)
		fmt.Printf("%+v", apiuser)
		//apiuser.Groups = m.GroupIds
	} else {
		fmt.Println("#####################update###########")
		apiuser.Oper = "UPDATE"
		apiuser.UID = m.Id
		apiuser.Name = m.UserName
		for _, groupId := range m.GroupIds {
			data := models.KubeUserGroup{}

			query := orm.NewOrm().QueryTable(models.KubeUserGroupTBName())
			query.Filter("id", groupId).One(&data)
			apiuser.Groups = data.GroupName + "," + apiuser.Groups
		}
		//models.SendUserModifyMsgToApiserver(&apiuser)

		//删除已关联的历史数据
		if _, err := o.QueryTable(models.RoleBackendUserRelTBName()).Filter("backenduser__id", m.Id).Delete(); err != nil {
			c.jsonResult(enums.JRCodeFailed, "删除历史关系失败", "")
		}

		if oM, err := models.BackendUserOne(m.Id); err != nil {
			c.jsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", m.Id)
		} else {
			m.UserPwd = strings.TrimSpace(m.UserPwd)
			if len(m.UserPwd) == 0 {
				//如果密码为空则不修改
				m.UserPwd = oM.UserPwd
			} else {
				m.UserPwd = utils.String2md5(m.UserPwd)
			}
			//本页面不修改头像和密码，直接将值附给新m
			m.Avatar = oM.Avatar
		}
		if _, err := o.Update(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
	err = models.SendUserModifyMsgToApiserver(&apiuser)
	if err != nil {
		beego.Error("---------k8s add or update error-------------")
	}
	//删除所有组关系
	models.KubeUser2GroupDeleteByUserId(m.Id)
	for _, groupId := range m.GroupIds {
		u2g := models.KubeUser2Group{UserId: m.Id, GroupId: groupId}
		u4, _ := uuid.NewV4()
		u2g.Id = u4.String()
		o.Insert(&u2g)
	}
	//添加关系
	var relations []models.RoleBackendUserRel
	for _, roleId := range m.RoleIds {
		r := models.Role{Id: roleId}
		relation := models.RoleBackendUserRel{BackendUser: &m, Role: &r}
		relations = append(relations, relation)
	}
	if len(relations) > 0 {
		//批量添加
		if _, err := o.InsertMulti(len(relations), relations); err == nil {
			c.jsonResult(enums.JRCodeSucc, "保存成功", m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed, "保存失败", m.Id)
		}
	} else {
		c.jsonResult(enums.JRCodeSucc, "保存成功", m.Id)
	}

}
func (c *BackendUserController) Delete() {
	strs := c.GetString("ids")
	if (strings.Compare(strs, "") != 0) {
		query := orm.NewOrm().QueryTable(models.BackendUserTBName())
		if num, err := query.Filter("id__in", strings.Split(strs, ",")).Delete(); err == nil {
			c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
		} else {
			c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
		}
	} else {
		c.jsonResult(enums.JRCodeFailed, "请选择用户", 0)
	}
}
