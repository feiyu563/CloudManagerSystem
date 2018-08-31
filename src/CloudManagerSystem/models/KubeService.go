package models

import (
	//"time"
	"fmt"
	"strings"
	"github.com/astaxie/beego/orm"
	"errors"
	//"github.com/satori/go.uuid"
	"time"
)

func init() {
	orm.RegisterModel(new(KubeService), new(KubeServicePort), new(KubeImage), new(KubeServiceIssue), new(KubeServiceVersion))
}

func KubeServiceTBName() string {
	return "kube_service"
}

type KubeServiceQueryParam struct {
	BaseQueryParam
	Id            string
	ServiceName   string
	NamespaceName string
	ImageName     string
}

type KubeServicePubORRollback struct {
	ServiceId string
	VersionId string
	Type      string
	Namespace string
}

// kind, namespace, name, count
type KubeServiceScale struct {
	Kind      string
	Namespace string
	Name      string
	Count     string
}

type JSONKubeService struct {
	Service     *KubeService
	ServicePort []*KubeServicePort
}

type KubeService struct {
	Id          string    `orm:"pk" orm:"size(40)"`
	Name        string    `orm:"size(40)"`
	ImageName     string    `orm:"size(40)"`  //镜像
	Env         string    `orm:"size(100)"` //环境变量
	Run         string    `orm:"size(100)"` //启动命令
	HostIp      bool
	CpuNeed     string
	CpuMax      string
	MemoryNeed  string
	MemoryMax   string
	ServiceNum  int                         //服务数量
	Heartbeat   string    `orm:"size(100)"` //监控检测
	RunTime     int                         //启动时间s
	SoketTime   int                         //超时时间s
	CreateUser  string    `orm:"size(40)"`
	CreateTime  time.Time `orm:"auto_now;type(datetime)"`
	Remark      string    `orm:"size(200)"`
	GroupId     string    `orm:"size(40)"` //运行区域
	NamespaceId string    `orm:"size(40)"`
	ClusterId   string    `orm:"size(40)"`
	IsVersion   int                         //是否为副本',
	FatherId    string    `orm:"size(40)"`  //主版本id',
	ServiceMark string    `orm:"size(40)"`  //服务标识',
	VersionId   string    `orm:"size(255)"` //'版本id'
}

func (a *KubeService) TableName() string {
	return KubeServiceTBName()
}

type KubeServiceRt struct {
	Id            string
	Name          string
	ImageName     string
	CpuNeed       string
	MemoryNeed    string
	ServicePort   string
	NamespaceName string
	NamespaceId   string
	//Tag           string
}
/*
	sql = fmt.Sprintf(`SELECT ks.id, ks.name,ks.namespace_id, ks.cpu_need , ks.memory_need , kn.name as namespace_name ,
	 ki.name as image_name , ki.tag as tag , ksp0.ports as service_port
	 FROM %s AS ks left join %s AS ki ON ki.id = ks.image_id
	 left join (
		select ksp.service_id service_id,group_concat(ksp.service_port order by ksp.service_port desc separator ',') ports
		from %s ksp
		group by ksp.service_id
		) ksp0 ON ksp0.service_id = ks.id
	 left join %s kn ON ks.namespace_id = kn.id
	 where ks.cluster_id = '%s'  and kn.name like '%s' and ki.name like '%s' and ks.name like '%s' and  is_version = 0
	 LIMIT %d OFFSET %d`,
		KubeServiceTBName(), KubeImageTBName(), KubeServicePortTBName(), KubeNameSpaceTBName(),
		ClusterId, namespacename_str, imagename_str, servicename_str, params.Limit, params.Offset)
*/
func KubeServicePageList(params *KubeServiceQueryParam, ClusterId string) ([]*KubeServiceRt, int64) {
	query := orm.NewOrm().QueryTable(KubeServiceTBName())

	data := make([]*KubeServiceRt, 0)
	//默认排序
	var total int64
	if params.Limit == 0 {
		params.Limit = 1000
	}
	var sql string

	o := orm.NewOrm()
	if strings.Compare(params.ServiceName, "") != 0 || strings.Compare(params.NamespaceName, "") != 0 || strings.Compare(params.ImageName, "") != 0 {

		var servicename_str, namespacename_str, imagename_str string

		servicename_str = "%" + params.ServiceName + "%"
		namespacename_str = "%" + params.NamespaceName + "%"
		imagename_str = "%" + params.ImageName + "%"
		//fmt.Println("-----------------------------", servicename_str, params.ServiceName)

		sql = fmt.Sprintf(`SELECT ks.id, ks.name, ks.namespace_id, ks.cpu_need , ks.memory_need , kn.name as namespace_name , 
         ks.image_name as image_name , ksp0.ports as service_port 
         FROM %s AS ks 
         left join (
			select ksp.service_id service_id,group_concat(ksp.service_port order by ksp.service_port desc separator ',') ports
			from %s ksp
			group by ksp.service_id
			) ksp0 ON ksp0.service_id = ks.id 
         left join %s kn ON ks.namespace_id = kn.id 
         where ks.cluster_id = '%s'  and kn.name like '%s' and ks.image_name like '%s' and ks.name like '%s' and  is_version = 0
         LIMIT %d OFFSET %d`,
			KubeServiceTBName(), KubeServicePortTBName(), KubeNameSpaceTBName(),
			ClusterId, namespacename_str, imagename_str, servicename_str, params.Limit, params.Offset)
		fmt.Println(sql)
		o.Raw(sql).QueryRows(&data)
		fmt.Printf("%+v", data)
		total, _ = query.Filter("cluster_id", ClusterId).Filter("is_version", 0).Count()

	} else if strings.Compare(params.Id, "") != 0 {
		sql = fmt.Sprintf(`SELECT ks.id, ks.name,ks.namespace_id, ks.cpu_need , ks.memory_need , kn.name as namespace_name , 
         ks.image_name as image_name , ksp0.ports as service_port 
         FROM %s AS ks 
         left join (
			select ksp.service_id service_id,group_concat(ksp.service_port order by ksp.service_port desc separator ',') ports
			from %s ksp
			group by ksp.service_id
			) ksp0 ON ksp0.service_id = ks.id 
         left join %s kn ON ks.namespace_id = kn.id 
         where ks.cluster_id = '%s' and  ks.id = '%s' and  is_version = 0
         LIMIT %d OFFSET %d`,
			KubeServiceTBName(), KubeServicePortTBName(), KubeNameSpaceTBName(), ClusterId, params.Id, params.Limit, params.Offset)

		fmt.Println(sql)
		o.Raw(sql).QueryRows(&data)
		total, _ = query.Filter("cluster_id", ClusterId).Filter("is_version", 0).Filter("id", params.Id).Count()

	} else {

		sql = fmt.Sprintf(`SELECT ks.id, ks.name,ks.namespace_id, ks.cpu_need , ks.memory_need , kn.name as namespace_name , 
         ks.image_name as image_name , ksp0.ports as service_port 
         FROM %s AS ks 
         left join (
			select ksp.service_id service_id,group_concat(ksp.service_port order by ksp.service_port desc separator ',') ports
			from %s ksp
			group by ksp.service_id
			) ksp0 ON ksp0.service_id = ks.id 
         left join %s kn ON ks.namespace_id = kn.id 
         where ks.cluster_id = '%s' and  is_version = 0 LIMIT %d OFFSET %d`,
			KubeServiceTBName(), KubeServicePortTBName(), KubeNameSpaceTBName(), ClusterId, params.Limit, params.Offset)

		fmt.Println(sql)
		o.Raw(sql).QueryRows(&data)
		total, _ = query.Filter("cluster_id", ClusterId).Filter("is_version", 0).Count()

	}

	return data, total
}

func KubeServiceIsExist(ServiceName string) (bool, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(KubeServiceTBName())
	bExist := qs.Filter("name", ServiceName).Exist()
	var err error
	if bExist {
		err = errors.New("Already Exist ")
		return false, err
	}

	return true, nil
}

func KubeServiceDelete(ids []string) (int64, error) {
	o := orm.NewOrm()
	for i, _ := range ids {
		qs := o.QueryTable(KubeServicePortTBName())
		if _, err := qs.Filter("service_id", ids[i]).Delete(); err != nil {
			err = errors.New("删除 " + ids[i] + "失败")
			return 0, err
		}
	}
	query := orm.NewOrm().QueryTable(KubeServiceTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

type KubeServiceToDeployment struct {
	Name      string
	ImageName string //镜像
	//ImageTag  string //镜像
	Env       string //环境变量
	Run       string //启动命令
	//HostIp      bool
	CpuNeed    string
	CpuMax     string
	MemoryNeed string
	MemoryMax  string
	ServiceNum int    //服务数量
	Heartbeat  string //监控检测
	RunTime    int    //启动时间s
	SoketTime  int    //超时时间s
	//GroupId     string //运行区域
	NamespaceName string
	//ServiceMark string //服务标识',
	//Version     string //'版本id'
}
//left join %s AS ki ON ki.id = ks.image_id
func KubeServiceVersionMessageGet(ServiceId string, ClusterId string, VersionId string) *KubeServiceToDeployment {
	var sql string
	data := make([]*KubeServiceToDeployment, 0)
	o := orm.NewOrm()
	sql = fmt.Sprintf(`SELECT  ks.name,ks.service_num,ks.cpu_max,ks.memory_max,ks.cpu_need,ks.memory_need,kn.name as namespace_name, 
         ks.env,ks.run,ks.heartbeat,ks.run_time,ks.soket_time,ks.image_name as image_name  
         FROM %s AS ks 
         left join %s kn ON ks.namespace_id = kn.id 
         where ks.cluster_id = '%s' and ks.father_id = '%s' and ks.version_id = '%s'`,
		KubeServiceTBName(), KubeNameSpaceTBName(), ClusterId, ServiceId, VersionId)

	fmt.Println(sql)
	o.Raw(sql).QueryRows(&data)
	if len(data) == 0 {
		return nil
	}
	return data[0]
}
