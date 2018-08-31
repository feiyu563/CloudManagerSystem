package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type KubePublishProxyQueryParam struct {
	BaseQueryParam
	Id            string
	Name          string
	ServiceId     string
	ServiceName   string
	ClusterId     string
	NamespaceId   string
	NamespaceName string
	Remark        string
}

type KubePublishProxy struct {
	Id          string    `orm:"pk;size(40)"`             //       varchar(40) not null comment '主键',
	ServiceId   string    `orm:"size(40)"`                //varchar(40) not null comment '服务id',
	PortId      string    `orm:"size(40)"`                // varchar(40) comment '服务端口id',
	CreateUser  string    `orm:"size(40)"`                //   varchar(40) comment '创建人',
	CreateTime  time.Time `orm:"auto_now;type(datetime)"` //  datetime comment '创建时间',
	Remark      string    `orm:"size(200)"`               //  varchar(200) comment '备注',
	ClusterId   string    `orm:"size(40)"`                //  varchar(40) not null comment '集群id',
	NamespaceId string    `orm:"size(40)"`                //      varchar(40) comment '命名空间id',
	Port        string    `orm:"size(40)"`                //       numeric(4) comment '主机端口',
	Name        string    `orm:"size(40)"`                //        varchar(40) comment '部署名称',
}

type KubePublishProxyVR struct {
	Id             string
	Name           string
	Port           string
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

func (c *KubePublishProxy) TableName() string {
	return KubePublishProxyTBName()
}

func KubePublishProxyTBName() string {
	return "kube_publish_proxy"
}

//扩展查询方法 ([]*KubeNameSpace, int64)
func KubePublishProxyPageList(p *KubePublishProxyQueryParam) ([]*KubePublishProxyVR, int64) {

	data := make([]*KubePublishProxyVR, 0)
	var sCount ProcSearchCount
	o := orm.NewOrm()

	var sql, sqlTotal string
	//sql =  fmt.Sprintf("select ns.*,cr.`name` cluster_name from kube_namespace ns left JOIN kube_cluster cr on (ns.cluster_id = cr.id)")
	sqlTotal = fmt.Sprintf("CALL proc_KubePublishProxyQTC (?,?,?,?,?,?,?,?,?,?,?,?)")
	sql = fmt.Sprintf("CALL proc_KubePublishProxyQPL (?,?,?,?,?,?,?,?,?,?,?,?)")

	fmt.Println(sql)

	o.Raw(sql, p.Id, p.Name, p.ServiceId, p.ServiceName, p.ClusterId, p.NamespaceId, p.NamespaceName, p.Remark, p.Sort, p.Order, p.Offset, p.Limit).QueryRows(&data)
	o.Raw(sqlTotal, p.Id, p.Name, p.ServiceId, p.ServiceName, p.ClusterId, p.NamespaceId, p.NamespaceName, p.Remark, p.Sort, p.Order, p.Offset, p.Limit).QueryRow(&sCount)

	return data, sCount.SearchCount
}

//批量删除
func KubePublishProxyDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubePublishProxyTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
