package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strings"
)

func KubeHostTBName() string {
	return "kube_host"
}

type KubeHostQueryParam struct {
	BaseQueryParam
	Id          	string `form:"hostId"`
	Ip           	string
	Lable           string `form:"lable"`
}

type KubeHost struct {
	Id          	string `orm:"pk"`
	Ip           	string
	Lable           string
	Role			string
	HostStatus      bool    //0,1
	HostName		string
	PassWord		string
	ClusterId		string
	CreateUser		string
	CreateTime		time.Time	`orm:"auto_now;type(datetime)"`
	Remark			string
	IsDeploy   		bool
	User       		string
	IsInstallNode 	bool
}

func (a *KubeHost) TableName() string {
	return KubeHostTBName()
}

//获取分页数据
func KubeHostPageList(params *KubeHostQueryParam) ([]*KubeHost, int64) {
	query := orm.NewOrm().QueryTable(KubeHostTBName())

	data := make([]*KubeHost, 0)
	//默认排序
	sortorder := "create_time"

	if(strings.Compare(params.Sort,"")!=0){
		sortorder = params.Sort
	}

	if params.Order == "desc" || params.Order == "asc" {
		sortorder = "-" + sortorder
	}
	if(strings.Compare(params.Ip,"")!=0){
		query = query.Filter("ip", params.Ip)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func GetAllKubeHostWithNOCluster(ClusterID string)([]*KubeHost, int64){
	query := orm.NewOrm().QueryTable(KubeHostTBName())

	data := make([]*KubeHost, 0)
	query = query.Filter("cluster_id", ClusterID)

	total, _ := query.Count()
	query.All(&data)
	//query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}
//批量删除
func KubeHostDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeHostTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

