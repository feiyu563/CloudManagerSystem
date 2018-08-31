package models

import (
	"github.com/astaxie/beego/orm"
)

func KubeServicePortTBName() string {
	return "kube_service_port"
}

type KubeServicePort struct {
	Id            string `orm:"pk" orm:"size(40)"`
	Name          string `orm:"size(40)"`
	Protocol      string `orm:"size(8)"`  // '协议'
	ContainerPort string                     // '容器端口',
	ServicePort   string                     //'服务端口',
	IsMain        string                  //'是否为主端口',
	ServiceId     string `orm:"size(40)"` //'服务id'
}

func (a *KubeServicePort) TableName() string {
	return KubeServicePortTBName()
}

func KubeServicePortGet(ServiceId string) ([]*KubeServicePort, int64) {
	data := make([]*KubeServicePort, 0)

	query := orm.NewOrm().QueryTable(KubeServicePortTBName())
	query.Filter("service_id", ServiceId).All(&data)
	total, _ := query.Count()
	return data, total
}




//批量删除
func ServicePortDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeServicePortTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
