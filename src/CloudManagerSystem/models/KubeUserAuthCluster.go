package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strings"
	//"fmt"
)

func init() {
	orm.RegisterModel(new(KubeUserAuthCluster))
}

type KubeUserAuthClusterQueryParam struct {
	BaseQueryParam
	Id   string
	Name string
}

func KubeUserAuthClusterTBName() string {
	return "kube_auth_user_cluster"
}

type KubeUserAuthCluster struct {
	Id         string    `orm:"pk"`
	UserId     string    `orm:"size(40)"`
	UserType   int
	ClusterId  string    `orm:"size(40)"`
	CreateUser string    `orm:"size(40)"`
	CreateTime time.Time `orm:"auto_now;type(datetime)"`
}

func (a *KubeUserAuthCluster) TableName() string {
	return KubeUserAuthClusterTBName()
}

func GetAllKubeAuthUser(params *KubeClusterQueryParam, ClusterId string) ([]*KubeUserAuthCluster, int64) {
	query := orm.NewOrm().QueryTable(KubeUserAuthClusterTBName())
	data := make([]*KubeUserAuthCluster, 0)
	//dataall := make([]*KubeCluster, 0)

	//默认排序
	sortorder := "id"

	if (strings.Compare(params.Sort, "") != 0) {
		sortorder = params.Sort
	}

	if params.Order == "desc" || params.Order == "asc" {
		sortorder = "-" + sortorder
	}

	if params.Limit == 0 {
		params.Limit = 100
	}

	query = query.Filter("cluster_id", ClusterId)

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}
