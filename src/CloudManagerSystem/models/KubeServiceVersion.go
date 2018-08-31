package models

import (
	"time"
	"fmt"
	//"strings"
	"github.com/astaxie/beego/orm"
)

func KubeServiceVersionTBName() string {
	return "kube_service_version"
}

//版本
type KubeServiceVersion struct {
	Id          string    `orm:"pk" orm:"size(40)"`
	VersionName string    `orm:"size(40)"`
	Remark      string    `orm:"size(255)"` //'发布原因'
	CreateUser  string    `orm:"size(40)"`
	CreateTime  time.Time `orm:"auto_now;type(datetime)"` //'发布时间',
}

func (a *KubeServiceVersion) TableName() string {
	return KubeServiceVersionTBName()
}

type JSONKubeServiceVersion struct {
	Id            string
	VersionName   string
	VersionRemark string
}

type KubeServiceVersionQueryParam struct {
	BaseQueryParam
	Id   string
	Name string
}

func KubeServiceVersionPageList(params *KubeServiceVersionQueryParam) ([]*KubeServiceVersion, int64) {
	query := orm.NewOrm().QueryTable(KubeServiceTBName())

	data := make([]*KubeServiceVersion, 0)
	//默认排序
    if params.Limit == 0 {
    	params.Limit = 1000
	}
	o := orm.NewOrm()
	var sql string
	sql = fmt.Sprintf(`SELECT ksv.id id ,ksv.version_name version_name, ksv.remark remark ,ksv.create_user create_user,ksv.create_time create_time,ks.father_id
  		FROM %s ks inner join %s ksv ON ksv.id = ks.version_id
			where ks.father_id = '%s' LIMIT %d OFFSET %d `, KubeServiceTBName(), KubeServiceVersionTBName(),params.Id,params.Limit, params.Offset)

	fmt.Println(sql)
	o.Raw(sql).QueryRows(&data)

	total, _ := query.Filter("father_id",params.Id).Count()
	return data, total
}
