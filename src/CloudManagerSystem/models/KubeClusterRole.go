package models

import (
	"strings"
	"time"
	"fmt"
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
)

func ClusterRoleTBName() string {
	return "kube_role"
}

type ClusterRoleQueryParam struct {
	BaseQueryParam
	Id   string
	Name string
}

//角色组名称 apiGroups 资源 操作
type ClusterSearchResource struct {
	Clustername string //`json:"clustername"`//orm:"size(40)"
	Rolename    string //`json:"rolename" form:"rolename"`//orm:"size(40)"
}

//Role  角色资源
type ClusterRole struct {
	Id          string `orm:"pk"`
	Name        string `orm:"size(40)"` //varchar(100)
	RoleType    int
	NamespaceId string `orm:"size(40)"`
	ClusterId   string `orm:"size(40)"`
	//Cluster            *Cluster           `orm:"rel(fk)"`//`orm:"size(40)"` //varchar(40)
	CreateUser string    `orm:"size(40)"` //varchar(40)
	CreateTime time.Time `orm:"auto_now;type(datetime)"`
	Remark     string    `orm:"size(200)"`
}

func (a *ClusterRole) TableName() string {
	return ClusterRoleTBName()
}

//获取分页数据
func KubeClusterPageList(params *ClusterRoleQueryParam, ClusterId string) ([]*ClusterRole, int64) {
	query := orm.NewOrm().QueryTable(ClusterRoleTBName())

	data := make([]*ClusterRole, 0)
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
	query = query.Filter("cluster_id", ClusterId)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func KubeClusterGetOne(RoleId string) ClusterRole {
	data := ClusterRole{}
	query := orm.NewOrm().QueryTable(ClusterRoleTBName())
	query.Filter("id", RoleId).One(&data)
	return data
}

func KubeClusterRoleGet(clusterid string) ([]*ClusterRole, int64) {
	data := make([]*ClusterRole, 0)

	query := orm.NewOrm().QueryTable(ClusterRoleTBName())
	query.Filter("cluster_id", clusterid).All(&data)
	total, _ := query.Count()
	return data, total
}

func ClusterRoleInsert(rolename string, ClusterId string, Username string) (string, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(ClusterRoleTBName())
	bExist := qs.Filter("name", rolename).Exist()
	var err error
	if bExist {
		err = errors.New("Already Exist ")
		return "", err
	}
	tem_uuid_t, _ := uuid.NewV4()

	rl := orm.NewOrm()
	role := ClusterRole{}
	role.Id = tem_uuid_t.String()
	role.Name = rolename
	role.ClusterId = ClusterId
	role.CreateUser = Username
	//fmt.Println("######################",role.Id)
	rl.Insert(&role)
	return role.Id, nil
}

//select*from content_node_relation where content_id like '%ddd%' and node_id like '%%';
func ResourceSearch(sr ClusterSearchResource) []ClusterSearchResource {
	var rev []ClusterSearchResource
	var rolename_str string
	o := orm.NewOrm()
	rolename_str = sr.Rolename + "%"
	var sql string
	sql = fmt.Sprintf(`SELECT T0.name AS clustername, T1.name AS rolename
  		FROM %s AS T0 inner join %s AS T1 ON T0.cluster_id = T1.id
		WHERE T0.name like '%s'`, ClusterRoleTBName(), KubeClusterTBName(), rolename_str)

	fmt.Println(sql)
	//fmt.Println(sr.Rolename)
	o.Raw(sql).QueryRows(&rev)
	return rev
}

func ClusterRoleGetALL() []string {
	var ret []string
	o := orm.NewOrm()

	var sql string
	sql = fmt.Sprintf(`SELECT name FROM %s `, ClusterRoleTBName())

	fmt.Println(sql)
	o.Raw(sql).QueryRows(&ret)
	return ret
}

//批量删除
func KubeClusterRoleDelete(ids []string) (int64, error) {
	o := orm.NewOrm()
	for i, _ := range ids {
		qs := o.QueryTable(ClusterResourceTBName())
		if _, err := qs.Filter("role_id", ids[i]).Delete(); err != nil {
			return 0, err
		}
	}
	query := orm.NewOrm().QueryTable(ClusterRoleTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
