package models

import (
	"time"
)

func KubeServiceIssueTBName() string {
	return "kube_service_issue"
}

//发布
type KubeServiceIssue struct {
	Id          string    `orm:"pk" orm:"size(40)"`
	Name        string    `orm:"size(40)"`
	ServiceId   string    `orm:"size(40)"` //服务id,
	CreateTime  time.Time `orm:"auto_now;type(datetime)"` //'发布时间',
	Remark      string    `orm:"size(255)"` //'发布原因'
	ServiceMark string    `orm:"size(40)"`//服务标识',
	Type        string    `orm:"size(6)"`
}

func (a *KubeServiceIssue) TableName() string {
	return KubeServiceIssueTBName()
}
