package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strings"
	"fmt"
)

func KubeUserGroupTBName() string {
	return "user_group"
}

type KubeUserGroupQueryParam struct {
	BaseQueryParam
	ClusterId		string
	ClusterName		string
	GroupName      	string
}

type KubeUserGroupVO struct {
	ClusterId		string
	KubeUserGroup   *KubeUserGroup
	KubeBinds		[]*KubeBind
}

type KubeUserGroup struct {
	Id          	string `orm:"pk"`
	GroupName      	string
	ClusterId      	string
	ClusterName		string	`orm:"-"`
	CreateUser		string
	CreateTime		time.Time	`orm:"auto_now;type(datetime)"`
	Remark			string
}

func (a *KubeUserGroup) TableName() string {
	return KubeUserGroupTBName()
}

//获取分页数据
func KubeUserGroupPageList(params *KubeUserGroupQueryParam) ([]*KubeUserGroup, int64) {
	query := orm.NewOrm().QueryTable(KubeUserGroupTBName())
	
	data := make([]*KubeUserGroup, 0)
	//默认排序
	sortorder := "create_time"
	
	if(strings.Compare(params.Sort,"")!=0){
		sortorder = params.Sort
	}

	if params.Order == "desc" || params.Order == "asc" {
		sortorder = "-" + sortorder
	}

	if(strings.Compare(params.ClusterId,"")!=0){
		query = query.Filter("cluster_id", params.ClusterId)
	}

	if(strings.Compare(params.GroupName,"")!=0){
		query = query.Filter("group_name", params.GroupName)
	}
	
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

//获取分页数据
func KubeUserGroupAllList(params *KubeUserGroupQueryParam) ([]*KubeUserGroup, error) {
	query := orm.NewOrm().QueryTable(KubeUserGroupTBName())

	data := make([]*KubeUserGroup, 0)
	//默认排序
	sortorder := "create_time"

	if(strings.Compare(params.Sort,"")!=0){
		sortorder = params.Sort
	}

	if params.Order == "desc" || params.Order == "asc" {
		sortorder = "-" + sortorder
	}

	if(strings.Compare(params.ClusterId,"")!=0){
		query = query.Filter("cluster_id", params.ClusterId)
	}

	if(strings.Compare(params.GroupName,"")!=0){
		query = query.Filter("group_name", params.GroupName)
	}
	_,err:=query.OrderBy(sortorder).All(&data)
	if(err==nil){
		return data, nil
	}
	return nil, err
}

//获取分页数据
func KubeUserGroupSqlPageList(params *KubeUserGroupQueryParam) ([]*KubeUserGroup, int64) {
	var rev []*KubeUserGroup
	o := orm.NewOrm()
	//默认排序
	sortorder := "create_time"
	var sql string
	sql=`SELECT ug.*,kc.name cluster_name
		FROM %s ug
		INNER JOIN %s kc on kc.id=ug.cluster_id`


	slq_date := fmt.Sprintf(sql+`ORDER BY %s ASC LIMIT %d OFFSET %d`, KubeUserGroupTBName(), KubeClusterTBName(),sortorder, params.Limit, params.Offset)


	var err error
	var total int64
	sql_total :=`SELECT count(ug.id)
		FROM %s ug`
	_,err=o.Raw(sql_total).QueryRows(&rev)
	if(err!=nil){
		fmt.Println("======err 1======")
		return nil,0
	}
	re, err:=o.Raw(slq_date).Exec()
	if(err!=nil){
		fmt.Println("======err 2======")
		return nil,0
	}
	total,err=re.LastInsertId()
	if(err!=nil){
		fmt.Println("======err 3======")
		return nil,0
	}
	return rev,total
}

//批量删除
func KubeUserGroupDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeUserGroupTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}




