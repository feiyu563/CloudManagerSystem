package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strings"
)

func KubeServiceAccountsTBName() string {
	return "kube_service_accounts"
}

type KubeServiceAccountsQueryParam struct {
	BaseQueryParam
	ClusterId		string
	ClusterName		string
	GroupName      	string
}

type KubeServiceAccountsVO struct {
	KubeServiceAccounts	*KubeServiceAccounts
	KubeBinds		[]*KubeBind
}

type KubeServiceAccounts struct {
	Id          	string `orm:"pk"`
	Name           	string
	ClusterId		string
	ClusterName		string `orm:"-"`
	NamespaceId     string
	CreateUser		string
	CreateTime		time.Time	`orm:"auto_now;type(datetime)"`
	Remark			string
}


//获取分页数据
func KubeServiceAccountsPageList(params *KubeServiceAccountsQueryParam) ([]*KubeServiceAccounts, int64) {
	query := orm.NewOrm().QueryTable(KubeServiceAccountsTBName())
	data := make([]*KubeServiceAccounts, 0)
	//默认排序
	sortorder := "create_time"
	
	if(strings.Compare(params.Sort,"")!=0){
		sortorder = params.Sort
	}

	if params.Order == "desc" || params.Order == "asc" {
		sortorder = "-" + sortorder
	}

	if(strings.Compare(params.GroupName,"")!=0){
		query = query.Filter("name", params.GroupName)
	}
	
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

//批量删除
func KubeServiceAccountsDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeServiceAccountsTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}