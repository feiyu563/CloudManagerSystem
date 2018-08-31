package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type KubeNameSpaceQueryParam struct {
	BaseQueryParam
	Id          	string
	Name          	string
	ClusterId       string
}

func KubeNameSpaceTBName()string{
	return "kube_namespace"
}
type KubeNameSpace struct {
	Id             string     `orm:"pk;size(40)"`          //varchar(40) not null comment '主键',
	Name           string      `orm:"size(40)"`            //varchar(40) not null comment '名称',
	Lable          string      `orm:"size(100)"`           //varchar(100) comment '标签',
	ClusterId      string      `orm:"size(40)"`            //varchar(40) not null comment '集群id',
	CreateUser     string      `orm:"size(40)"`            //varchar(40) comment '创建人',
	CreateTime     time.Time   `orm:"auto_now;type(datetime)"`      //datetime comment '创建时间',
	Remark         string      `orm:"size(200)"`           //varchar(200) comment '备注'
	Stype          int                                     //numeric(1) comment '类型',
}

type KubeNameSpaceVR struct {
	Id             string
	Name           string
	Lable          string
	ClusterId      string
	ClusterName    string
	CreateUser     string
	CreateTime     time.Time
	Remark         string
	Stype          int
}

type ProcSearchCount struct {
	SearchCount int64
}

func(c *KubeNameSpace)TableName()string{
	return KubeNameSpaceTBName()
}


//获取分页数据
func KubeNameSpacePageList(params *KubeNameSpaceQueryParam) ([]*KubeNameSpaceVR, int64) {
	//query := orm.NewOrm().QueryTable(KubeNameSpaceTBName())
	//
	//data := make([]*KubeNameSpace, 0)
	////默认排序
	//sortorder := "id"
	//
	//if(strings.Compare(params.Sort,"")!=0){
	//	sortorder = params.Sort
	//}
	//
	//if params.Order == "desc" || params.Order == "asc" {
	//	sortorder = "-" + sortorder
	//}
	//
	////
	//if(strings.Compare(params.Name,"")!=0){
	//	query = query.Filter("name", params.Name)
	//}
	//
	//if(strings.Compare(params.Id,"")!=0){
	//	query = query.Filter("id", params.Id)
	//}
	//
	//total, _ := query.Count()
	//query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	//
	//return data, total

	return ExtendPageList(params)
}

//批量删除
func KubeNameSpaceDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeNameSpaceTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

//扩展查询方法 ([]*KubeNameSpace, int64)
func ExtendPageList(p *KubeNameSpaceQueryParam)([]*KubeNameSpaceVR, int64){

	data := make([]*KubeNameSpaceVR, 0)
	var sCount ProcSearchCount
	o := orm.NewOrm()

	var sql ,sqlTotal string
	//sql =  fmt.Sprintf("select ns.*,cr.`name` cluster_name from kube_namespace ns left JOIN kube_cluster cr on (ns.cluster_id = cr.id)")
	sqlTotal =  fmt.Sprintf("CALL proc_KubeNameSpaceQTC (?,?,?,?,?,?,?)")
	sql =  fmt.Sprintf("CALL proc_KubeNameSpaceQPL (?,?,?,?,?,?,?)")

	fmt.Println(sql)

	o.Raw(sql,p.Id,p.Name,p.ClusterId,p.Sort,p.Order,p.Offset,p.Limit).QueryRows(&data)
	o.Raw(sqlTotal,p.Id,p.Name,p.ClusterId,p.Sort,p.Order,p.Offset,p.Limit).QueryRow(&sCount)


	return data, sCount.SearchCount
	/*fmt.Println("")
	fmt.Printf("%+v", data)
	fmt.Println("")
	fmt.Printf("%+v",sCount)
	fmt.Println("")*/
}