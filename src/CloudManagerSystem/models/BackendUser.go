package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"strconv"
)
//获取 BackendUser 对应的表名称
func BackendUserTBName() string {
	return "rms_backend_user"
}

type BackendUserQueryParam struct {
	BaseQueryParam
	UserNameLike string //模糊查询
	RealNameLike string //模糊查询
	Mobile       string //精确查询
	SearchStatus string //为空不查询，有值精确查询
	UserTypes	 string
}
type BackendUser struct {
	Id                 string `orm:"pk"`
	RealName           string `orm:"size(32)"`
	UserName           string `orm:"size(24)"`
	UserPwd            string `json:"-"`
	IsSuper            bool
	Status             int
	UserType		   int
	Mobile             string                `orm:"size(16)"`
	Email              string                `orm:"size(256)"`
	Avatar             string                `orm:"size(256)"`
	RoleIds            []int                 `orm:"-" form:"RoleIds"`
	RoleBackendUserRel []*RoleBackendUserRel `orm:"reverse(many)"` // 设置一对多的反向关系
	ResourceUrlForList []string              `orm:"-"`
	GroupIds 		   []string              `orm:"-" form:"GroupIds"`
}

func (a *BackendUser) TableName() string {
	return BackendUserTBName()
}

//获取分页数据
func BackendUserPageList(params *BackendUserQueryParam) ([]*BackendUser, int64) {
	query := orm.NewOrm().QueryTable(BackendUserTBName())
	data := make([]*BackendUser, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("username__istartswith", params.UserNameLike)
	query = query.Filter("realname__istartswith", params.RealNameLike)
	if len(params.Mobile) > 0 {
		query = query.Filter("mobile", params.Mobile)
	}
	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

//获取分页数据
func BackendUserAllList(params *BackendUserQueryParam) ([]*BackendUser, error) {
	query := orm.NewOrm().QueryTable(BackendUserTBName())
	data := make([]*BackendUser, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}

	if len(params.Mobile) > 0 {
		query = query.Filter("mobile", params.Mobile)
	}

	if len(params.SearchStatus) > 0 {
		query = query.Filter("status", params.SearchStatus)
	}
	if(len(params.UserTypes)>0){
		list:=strings.Split(params.UserTypes,",")
		types := make([]int, 0, len(list))
		for _, str := range list {
			if id, err := strconv.Atoi(str); err == nil {
				types = append(types, id)
			}
		}
		query = query.Filter("user_type__in", types)
	}
	_,err:=query.OrderBy(sortorder).All(&data)
	if(err==nil){
		return data,nil
	}
	return nil,err
}

// 根据id获取单条
func BackendUserOne(id string) (*BackendUser, error) {
	o := orm.NewOrm()
	m := BackendUser{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// 根据用户名密码获取单条
func BackendUserOneByUserName(username, userpwd string) (*BackendUser, error) {
	m := BackendUser{}
	err := orm.NewOrm().QueryTable(BackendUserTBName()).Filter("username", username).Filter("userpwd", userpwd).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
