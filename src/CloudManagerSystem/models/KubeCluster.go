package models

import (
	"time"
	"fmt"
	"strings"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
)

//func init() {
//	orm.RegisterModel(new(KubeCluster))
//}

func KubeClusterTBName() string {
	return "kube_cluster"
}

type KubeClusterQueryParam struct {
	BaseQueryParam
	Id     string
	Name   string
	UserId string
	Ip     string
}

//type JSONKubeClusterRelationT struct{
//	BaseQueryParam
//	Id         string
//
//}
const (
	ERRORTYPE = -1
	USERTYPE  = 0x11
	HOSTTYPE  = 0x12
)

var TypeText = map[int]string{
	ERRORTYPE: "Unknow Error",
	USERTYPE:  "Authentication Action!",
	HOSTTYPE:  "Host Action",
}

type JSONKubeClusterRelation struct {
	Type      int //11 host , 12 user
	ClusterId string
	//HostName           		[]*string
	//Users              		[]*string
	KubeClusterQueryParam []*KubeClusterQueryParam
}

type KubeCluster struct {
	Id         string    `orm:"pk"`
	Name       string
	CreateUser string
	CreateTime time.Time `orm:"auto_now;type(datetime)"`
	Remark     string
}

type KubeClusterRT struct {
	KubeCluster
	//Id       string
	//Name     string
	Ips        string
	HostName   string
	HostStatus bool
	Remark     string
}

func (a *KubeCluster) TableName() string {
	return KubeClusterTBName()
}

//集群名称 节点 是否部署 备注
func GetAllKubeCluster(params *KubeClusterQueryParam) ([]*KubeClusterRT, int64) {

	query := orm.NewOrm().QueryTable(KubeClusterTBName())
	data := make([]*KubeClusterRT, 0)
	//dataall := make([]*KubeCluster, 0)

	//默认排序
	sortorder := "id"

	if (strings.Compare(params.Sort, "") != 0) {
		sortorder = params.Sort
	}

	if params.Order == "desc" || params.Order == "asc" {
		sortorder = "-" + sortorder
	}

	//
	if (strings.Compare(params.Name, "") != 0) {
		query = query.Filter("name", params.Name)
	}

	if (strings.Compare(params.Id, "") != 0) {
		query = query.Filter("id", params.Id)
	}

	total, _ := query.Count()

	if params.Limit == 0 {
		params.Limit = 100
	}
	var sql string
	//sql = fmt.Sprintf(`SELECT T0.id , T0.name , T1.ip ,T1.host_name, T0.remark
  	//	FROM %s  T0 left join %s  T1 ON T0.id = T1.cluster_id
	//	ORDER BY T0.id ASC LIMIT %d OFFSET %d`, KubeClusterTBName(), KubeHostTBName(), params.Limit, params.Offset)
	o := orm.NewOrm()

	if strings.Compare(params.Name, "") != 0 {
		sql = fmt.Sprintf(`SELECT kc.*,t_host.ips,t_host.host_name,t_host.host_status
			from %s kc
			LEFT JOIN (
			select kh.cluster_id cluster_id,kh.host_status host_status ,kh.host_name host_name ,group_concat(kh.ip order by kh.ip desc separator ',') ips
			from %s kh
			group by kh.cluster_id
			)t_host on t_host.cluster_id=kc.id
			WHERE kc.name = '%s'
			LIMIT %d OFFSET %d`, KubeClusterTBName(), KubeHostTBName() , params.Name , params.Limit, params.Offset)
		fmt.Println(sql)
		o.Raw(sql).QueryRows(&data)
	} else {
		sql = fmt.Sprintf(`SELECT kc.*,t_host.ips,t_host.host_name,t_host.host_status
			from %s kc
			LEFT JOIN (
			select kh.cluster_id cluster_id,kh.host_status host_status ,kh.host_name host_name ,group_concat(kh.ip order by kh.ip desc separator ',') ips
			from %s kh
			group by kh.cluster_id
			)t_host on t_host.cluster_id=kc.id
			LIMIT %d OFFSET %d`, KubeClusterTBName(), KubeHostTBName(), params.Limit, params.Offset)
		fmt.Println(sql)
		o.Raw(sql).QueryRows(&data)
	}

	//total := len(data)
	//fmt.Println(total)

	return data, int64(total)
}

//INSERT INTO db1_name(field1,field2) SELECT field1,field2 FROM db2_name
func KubeClusterAuthRelation(parms KubeClusterQueryParam, ClusterId string) error {
	var err error
	if strings.Compare(parms.Id, "") != 0 {
		query := orm.NewOrm().QueryTable(KubeUserAuthClusterTBName())
		_, err := query.Filter("id", parms.Id).Delete()
		if err != nil {
			return err
		}
	}
	ur := orm.NewOrm()
	user := BackendUser{Id: parms.UserId}
	err = ur.Read(&user)
	if ur.Read(&user) != nil {
		return err
	}

	tem_uuid_t, _ := uuid.NewV4()

	kuac := orm.NewOrm()
	authcluster := KubeUserAuthCluster{}
	authcluster.Id = tem_uuid_t.String()
	authcluster.UserId = user.Id
	authcluster.ClusterId = ClusterId
	//authcluster.UserType
	_, err1 := kuac.Insert(&authcluster)
	if err1 == nil {
		return err1
	}

	return nil
}

func KubeClusterHostRelationDel(Id string)(error) {

	var sql string
	sql = fmt.Sprintf(`UPDATE %s T0 
  		SET T0.cluster_id = ""
		WHERE T0.id = ?`, KubeHostTBName())

	o := orm.NewOrm()
	fmt.Println(sql)
	_, err := o.Raw(sql, Id).Exec()
	if err != nil {
		return err
	}
	return nil
}

func KubeClusterAuthRelationDel(Id string)(error) {
	var sql string
	sql = fmt.Sprintf(`DELETE from %s
		WHERE id = '%s'`, KubeUserAuthClusterTBName(),Id)

	o := orm.NewOrm()
	fmt.Println(sql)
	_, err := o.Raw(sql).Exec()
	if err != nil {
		return err
	}
	return nil
}

func KubeClusterNodeRelation(parms KubeClusterQueryParam, ClusterId string) error {
	var sql string
	sql = fmt.Sprintf(`UPDATE %s T0 
  		SET T0.cluster_id = ?
		WHERE T0.ip = '%s'`, KubeHostTBName(), parms.Ip)

	o := orm.NewOrm()
	fmt.Println(sql)
	_, err := o.Raw(sql, ClusterId).Exec()
	if err != nil {
		return err
	}
	return nil
}

func KubeClusterNodeRelationDel(ClusterId string) error {
	var sql string
	sql = fmt.Sprintf(`UPDATE %s T0 
  		inner join %s T1 ON T0.id = T1.cluster_id SET T1.cluster_id = ""
		WHERE T1.cluster_id = ?`, KubeClusterTBName(), KubeHostTBName())

	o := orm.NewOrm()
	fmt.Println(sql)
	_, err := o.Raw(sql, ClusterId).Exec()
	if err != nil {
		return err
	}
	return nil
}

//批量删除
func KubeClusterDelete(ids []string) (int64, error) {
	o := orm.NewOrm()
	for i, _ := range ids {
		qs := o.QueryTable(KubeUserAuthClusterTBName())
		if _, err := qs.Filter("cluster_id", ids[i]).Delete(); err != nil {
			return 0, err
		}

		var sql string
		sql = fmt.Sprintf(`UPDATE %s T0 
  		SET T0.cluster_id = ""
		WHERE T0.cluster_id = '%s'`, KubeHostTBName(),ids[i])

		fmt.Println(sql)
		_, err := o.Raw(sql).Exec()
		if err != nil {
			return 0, err
		}

	}
	query := orm.NewOrm().QueryTable(KubeClusterTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

