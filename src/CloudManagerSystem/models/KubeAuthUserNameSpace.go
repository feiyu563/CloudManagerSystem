package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type KubeAuthUserNameSpaceQueryParam struct {
	BaseQueryParam
	Id          	    string
	UserId          	string
	UserName          	string
	ClusterId          	string
	ClusterName         string
	NamespaceId         string
	NamespaceName       string
}

type KubeAuthUserNameSpace struct {
	Id                  string    `orm:"pk;size(40)"`//varchar(40) not null comment '主键',
	UserId              string    `orm:"size(40)"`   //varchar(40) not null,
	UserType            int                       //numeric(1) comment '用户类型：user:0/group:1/sa:2',
	NamespaceId         string    `orm:"size(40)"`   //varchar(40) not null,
	ClusterId           string    `orm:"size(40)"`   //varchar(40) not null comment '集群id',
	CreateUser          string    `orm:"size(40)"`  //varchar(40) comment '创建人',
	CreateTime          time.Time `orm:"auto_now;type(datetime)"`//datetime comment '创建时间',
}

type KubeAuthUserNameSpaceMini struct {
	Id                  string
	UserId              string
	UserType            int        //numeric(1) comment '用户类型：user:0/group:1/sa:2',
}

type JSONKubeAuthUserNameSpace struct {
	NamespaceId string
	ClusterId string
	NameSpacesAuthUser []*KubeAuthUserNameSpaceMini
}


type KubeAuthUserNameSpaceVR struct {
	Id                  string
	UserId              string
	UserName            string
	UserType            int
	NamespaceId         string
	NamespaceName       string
	ClusterId           string
	ClusterName         string
	CreateUser          string
	CreateTime          time.Time
}

func KubeAuthUserNameSpaceTBName()string{
	return "kube_auth_user_namespace"
}
 // beego orm 必须的函数名
func(c *KubeAuthUserNameSpace)TableName()string{
	return KubeAuthUserNameSpaceTBName()
}

//批量删除
func KubeAuthUserNameSpaceDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeAuthUserNameSpaceTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

func KubeAuthUserNameSpaceDeleteByNsIdNotInId(d *map[string]string) (int64, int64) {

	var rNum ,errNum int64
	o := orm.NewOrm()
	for nsId,idstr := range *d{

		sql  := fmt.Sprintf("DELETE FROM %s WHERE id not in ( %s) and namespace_id =  '%s'", KubeAuthUserNameSpaceTBName(),idstr,nsId)

		//fmt.Println(sql)
		o.Raw(sql).Exec()

	}
	return rNum, errNum
}


//获取分页数据
func KubeAuthUserNameSpacePageList(p *KubeAuthUserNameSpaceQueryParam) ([]*KubeAuthUserNameSpaceVR, int64) {
	data := make([]*KubeAuthUserNameSpaceVR, 0)
	var sCount ProcSearchCount
	o := orm.NewOrm()

	var sql ,sqlTotal string
	sql =  fmt.Sprintf("CALL proc_KubeAuthUserNameSpaceQPL (?,?,?,?,?,?,?,?,?,?,?)")
	sqlTotal =  fmt.Sprintf("CALL proc_KubeAuthUserNameSpaceQTC (?,?,?,?,?,?,?,?,?,?,?)")

	//fmt.Println(sql)
	o.Raw(sql,p.Id,p.UserId,p.UserName,p.ClusterId,p.ClusterName,p.NamespaceId,p.NamespaceName,p.Sort,p.Order,p.Offset,p.Limit).QueryRows(&data)
	o.Raw(sqlTotal,p.Id,p.UserId,p.UserName,p.ClusterId,p.ClusterName,p.NamespaceId,p.NamespaceName,p.Sort,p.Order,p.Offset,p.Limit).QueryRow(&sCount)

	return data, sCount.SearchCount
}