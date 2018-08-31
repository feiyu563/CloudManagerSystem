package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type KubePublishServiceQueryParam struct {
	BaseQueryParam
	Id            string
	Stype         string
	Name          string
	DomainName    string
	ServiceId     string
	ServiceName   string
	ClusterId     string
	NamespaceId   string
	NamespaceName string
	Remark        string
}

type KubePublishService struct {
	Id          string                    `orm:"pk;size(40)"`             //       varchar(40) not null comment '主键',
	DomainName  string                    `orm:"size(40)"`                //        varchar(40) comment '域名',
	Name        string                    `orm:"size(40)"`                //         varchar(40) comment 'name',
	Ramark      string                    `orm:"size(100)"`               //         varchar(100) comment '备注',
	CreateUser  string                    `orm:"size(40)"`                //        varchar(40) comment '创建人',
	CreateTime  time.Time                 `orm:"auto_now;type(datetime)"` //        datetime comment '创建时间',
	ClusterId   string                    `orm:"size(40)"`                //        varchar(40) comment '集群id',
	NamespaceId string                    `orm:"size(40)"`                //        varchar(40) comment '命名空间id',
	Stype       string                    `orm:"size(8)"`                 //        协议类型 tcp/http
	ServiceId   string                    `orm:"size(40)"`                //      服务id
	DeployName  string                    `orm:"size(40)"`                //      服务id
	Paths       []*KubePublishServicePath `orm:"-"`
}

type KubePublishServiceVR struct {
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
	Stype          string
}

func (c *KubePublishService) TableName() string {
	return KubePublishServiceTBName()
}

func KubePublishServiceTBName() string {
	return "kube_publish_service"
}

//扩展查询方法 ([]*KubeNameSpace, int64)
func KubePublishServicePageList(p *KubePublishServiceQueryParam) ([]*KubePublishServiceVR, int64) {

	data := make([]*KubePublishServiceVR, 0)
	var sCount ProcSearchCount
	o := orm.NewOrm()

	var sql, sqlTotal string
	//sql =  fmt.Sprintf("select ns.*,cr.`name` cluster_name from kube_namespace ns left JOIN kube_cluster cr on (ns.cluster_id = cr.id)")
	sqlTotal = fmt.Sprintf("CALL proc_KubePublishServiceQTC (?,?,?,?,?,?,?,?,?,?,?,?,?)")
	sql = fmt.Sprintf("CALL proc_KubePublishServiceQPL (?,?,?,?,?,?,?,?,?,?,?,?,?)")

	fmt.Println(sql)

	o.Raw(sql, p.Id, p.Stype, p.Name, p.ServiceId, p.ServiceName, p.ClusterId, p.NamespaceId, p.NamespaceName, p.Remark, p.Sort, p.Order, p.Offset, p.Limit).QueryRows(&data)
	o.Raw(sqlTotal, p.Id, p.Stype, p.Name, p.ServiceId, p.ServiceName, p.ClusterId, p.NamespaceId, p.NamespaceName, p.Remark, p.Sort, p.Order, p.Offset, p.Limit).QueryRow(&sCount)

	return data, sCount.SearchCount
}

//批量删除
func KubePublishServiceDelete(ids []string) (int64, error) {
	//删除子表
	queryPath := orm.NewOrm().QueryTable(KubePublishServicePathTBName())
	num, err := queryPath.Filter("pservice_id__in", ids).Delete()

	query := orm.NewOrm().QueryTable(KubePublishServiceTBName())
	num, err = query.Filter("id__in", ids).Delete()

	return num, err
}
