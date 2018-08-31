package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"fmt"
	"github.com/satori/go.uuid"
)

func init() {
	orm.RegisterModel(new(ResourceOper))
}
func ClusterResourceTBName() string {
	return "kube_resource"
}

//Resource 权限控制资源表
type ClusterResource struct {
	Id            string `orm:"pk"`
	ResType       int                      //`orm:""`// numeric(1) not null comment '资源类型 resources:0/url:1/node:2',
	OperRole      string `orm:"size(100)"` //varchar(100) not null comment '操作权限',
	ApiGroups     string `orm:"size(1)"`   //varchar(1) comment 'apiGroups',
	ResourceNames string `orm:"size(40)"`  //varchar(40) not null comment '资源名称',
	RoleType      int                      //numeric(1) not null comment 'role:0/cluterrole:1',
	RoleId        string `orm:"size(40)"`  //varchar(40) not null comment '角色id',
	UserOperType  int                      //varchar(40) not null comment '主键',
}
type ClusterResourceQueryParam struct {
	BaseQueryParam
	Id   string
	Name string
}

type ResourceOperQueryParam struct {
	//BaseQueryParam
	Id            string
	ResourceNames string
	OperName      string
}
type JSONResourceOperQueryParam struct {
	RoleID                 string
	RoleName               string
	ResourceOperQueryParam []*ResourceOperQueryParam
}

func (a *ClusterResource) TableName() string {
	return ClusterResourceTBName()
}

func ResourceOper_DeFTBName() string {
	return "kube_oper_definition"
}

type ResourceOper struct {
	Id       string `orm:"pk" orm:"size(40)"`
	OperType int
	OperName string `orm:"size(100)"`
}

func (a *ResourceOper) TableName() string {
	return ResourceOper_DeFTBName()
}

type ResourceReturn struct {
	Id            string
	ResourceNames string
	OperName      string
}

type ResourceReturns struct {
	Count int
	Res   []ResourceReturn
}

//获取分页数据
//SELECT T0.id, T0.resource_names, T1.oper_name FROM kube_resource T0
//inner join kube_oper_definition AS T1 ON T0.user_oper_type = T1.oper_type ORDER BY T0.`id` ASC LIMIT 1000 offset 10
func KubeResourceOperPageList(params *ClusterResourceQueryParam, rolename string) ([]ResourceReturn, int64) {
	query := orm.NewOrm().QueryTable(ClusterResourceTBName())
	var data []ResourceReturn
	//默认排序
	sortorder := "id"

	if (strings.Compare(params.Sort, "") != 0) {
		sortorder = params.Sort
	}

	if params.Order == "desc" || params.Order == "asc" {
		sortorder = "-" + sortorder
	}

	//
	if (strings.Compare(params.Name, "") != 0) {
		query = query.Filter("name", params.Name)
	}

	if (strings.Compare(params.Id, "") != 0) {
		query = query.Filter("id", params.Id)
	}
	//total, _ := query.Count()
	if params.Limit == 0 {
		params.Limit = 100
	}
	var sql string
	sql = fmt.Sprintf(`SELECT T0.id , T0.resource_names , T1.oper_name
  		FROM %s AS T0 inner join %s AS T1 ON T0.user_oper_type = T1.oper_type
		inner join %s T2 on T2.id = T0.role_id
		WHERE T2.name = ?
		ORDER BY T0.id ASC LIMIT %d OFFSET %d`, ClusterResourceTBName(), ResourceOper_DeFTBName(), ClusterRoleTBName(), params.Limit, params.Offset)

	o := orm.NewOrm()
	fmt.Println(sql)
	o.Raw(sql, rolename).QueryRows(&data)

	total := len(data)
	fmt.Println(total)

	return data, int64(total)
}

func ResourceGetONE(id string) ([]ResourceReturn, int64) {
	var data []ResourceReturn
	query := orm.NewOrm().QueryTable(ClusterResourceTBName())
	query = query.Filter("id", id)
	total, _ := query.Count()

	var sql string
	sql = fmt.Sprintf(`SELECT T0.id , T0.resource_names , T1.oper_name
  		FROM %s AS T0 inner join %s AS T1 ON T0.user_oper_type = T1.oper_type
		WHERE T0.id = ?`, ClusterResourceTBName(), ResourceOper_DeFTBName())

	o := orm.NewOrm()
	fmt.Println(sql)
	o.Raw(sql, id).QueryRows(&data)
	return data, total
}

func UpdateResourceOper(parms ResourceOperQueryParam) (bool, error) {
	op := orm.NewOrm()
	var oper []ResourceOper
	var sql string
	//fmt.Println("--1--")

	sql = fmt.Sprintf(`SELECT id , oper_type , oper_name
  		FROM %s 
		WHERE oper_name = ?`, ResourceOper_DeFTBName())
	//fmt.Println("---2-")
	//fmt.Println(sql)

	op.Raw(sql, parms.OperName).QueryRows(&oper)
	//fmt.Println("----")
	fmt.Println(oper)

	var err error

	rs := orm.NewOrm()
	resource := ClusterResource{Id: parms.Id}
	if rs.Read(&resource) == nil {
		//resource.Name = "MyName"
		resource.ResourceNames = parms.ResourceNames
		resource.UserOperType = oper[0].OperType
		if num, err := rs.Update(&resource); err == nil {
			fmt.Println(num)
			return true, nil
		}
	}
	fmt.Println("error update exit")
	return false, err
}

func InsertResourceOper(parms ResourceOperQueryParam, RoleId string) (bool, error) {
	op := orm.NewOrm()
	//var crole []ClusterRole
	var sql string

	//sql = fmt.Sprintf(`SELECT id FROM %s
	//	WHERE name = ?`, ClusterRoleTBName())
	//op.Raw(sql, RoleName).QueryRows(&crole)
	////fmt.Println("----")
	//fmt.Println(crole)

	var oper []ResourceOper

	sql = fmt.Sprintf(`SELECT oper_type FROM %s 
		WHERE oper_name = ?`, ResourceOper_DeFTBName())

	op.Raw(sql, parms.OperName).QueryRows(&oper)
	//fmt.Println("----")
	fmt.Println(oper)

	//	title = "添加"
	tem_uuid_t, _ := uuid.NewV4()

	rs := orm.NewOrm()
	resource := ClusterResource{}
	resource.Id = tem_uuid_t.String()
	resource.RoleId = RoleId
	resource.ResourceNames = parms.ResourceNames
	resource.UserOperType = oper[0].OperType
	_, err := rs.Insert(&resource)
	if err == nil {
		return true, nil
	}

	fmt.Println("-----error insert exit")
	return false, err
}

func GetResourceALL() ([]ClusterResource, int64) {
	var data []ClusterResource

	query := orm.NewOrm().QueryTable(ClusterResourceTBName())
	query.All(&data)
	total, _ := query.Count()

	return data, total
}

func GetOperALL() ([]ResourceOper, int64) {
	var data []ResourceOper
	query := orm.NewOrm().QueryTable(ResourceOper_DeFTBName())
	query.All(&data)
	total, _ := query.Count()
	return data, total
}

//select*from content_node_relation where content_id like '%ddd%' and node_id like '%%';

//批量删除
func ResourceDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(ClusterResourceTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
