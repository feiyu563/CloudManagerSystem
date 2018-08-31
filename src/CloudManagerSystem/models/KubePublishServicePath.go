package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type KubePublishServicePathQueryParam struct {
	BaseQueryParam
	Id            string
	PserviceId    string
	ServiceId     string
	ServiceName   string
	ClusterId     string
	NamespaceId   string
	NamespaceName string
}

type KubePublishServicePath struct {
	Id         string    `orm:"pk;size(40)"` //       varchar(40) not null comment '主键',
	Path       string    `orm:"size(40)"`    //         varchar(40) comment '路径',
	ServiceId  string    `orm:"size(40)"`    //        varchar(40) comment '服务id',
	PortId     string    `orm:"size(40)"`    //        varchar(40) comment '端口id',
	PserviceId string    `orm:"size(40)"`    //        varchar(40) comment '部署服务Id',
	HostPort   string    `orm:"size(8)"`     //        varchar(40) comment '宿主机端口',
	CreateTime time.Time `orm:"auto_now;type(datetime)"`
}

type KubePublishServicePathVR struct {
	Id             string
	Name           string
	DomainName     string
	ContainerPort  string
	ServiceId      string
	ServiceName    string
	ServicePortId  string
	ClusterId      string
	ClusterName    string
	NamespaceName  string
	CreateUser     string
	CreateUserName string
	CreateTime     time.Time
	Remark         string
}

func (c *KubePublishServicePath) TableName() string {
	return KubePublishServicePathTBName()
}

func KubePublishServicePathTBName() string {
	return "kube_publish_service_path"
}

//扩展查询方法 ([]*KubeNameSpace, int64)
func KubePublishServicePathPageList(p *KubePublishServicePathQueryParam) ([]*KubePublishServicePath, int64) {

	sortorder := "id"
	data := make([]*KubePublishServicePath, 0)
	query := orm.NewOrm().QueryTable(KubePublishServicePathTBName())

	query = query.Filter("pservice_id", p.PserviceId)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(p.Limit, p.Offset).All(&data)
	return data, total
}

//批量删除
func KubePublishServicePathDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubePublishServicePathTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func KubePublishServicePathByIngressId(ingressId string) (int64, int64) {

	var rNum, errNum int64
	o := orm.NewOrm()
	sql := fmt.Sprintf("DELETE FROM %s WHERE pservice_id =  '%s'", KubePublishServicePathTBName(), ingressId)
	//fmt.Println(sql)
	o.Raw(sql).Exec()
	return rNum, errNum
}
