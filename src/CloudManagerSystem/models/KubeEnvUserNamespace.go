package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func KubeEnvUserNamespaceTBName() string {
	return "kube_env_user_namespace"
}

type KubeEnvUserNamespaceQueryParam struct {
	BaseQueryParam
}

type KubeEnvUserNamespace struct {
	Id          	string `orm:"pk"`
	UserId      	string
	NamespaceId     string
	NamespaceName   string `orm:"-"`
	ClusterId      	string
}

type KubeEnvUserNamespaceMore struct{
	ClusterId string
	Eunsm [] *KubeEnvUserNamespace
}

func (a *KubeEnvUserNamespace) TableName() string {
	return KubeEnvUserNamespaceTBName()
}

func EnvUserNamespace(userId string) (*KubeEnvUserNamespace, error) {
	o := orm.NewOrm()
	m := KubeEnvUserNamespace{UserId: userId}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

type SeleNameSpaceEnvUser_VR struct {
	Id         string
	Name       string
	Remark     string
	Selected   string
}

//获取环境变量的值+
func SeleNameSpaceEnvUser(userId ,clusterId string)([] *SeleNameSpaceEnvUser_VR){
	o := orm.NewOrm()

	data := make([] *SeleNameSpaceEnvUser_VR, 0)
	sql := fmt.Sprintf("CALL proc_EnvUserNameSpace (?,?)")
	o.Raw(sql,userId,clusterId).QueryRows(&data)

	return data
}

func GetEnvNameSpaceByUserIdClusterId(userId ,clusterId string)( *SeleNameSpaceEnvUser_VR){
	o := orm.NewOrm()

	data :=SeleNameSpaceEnvUser_VR {} //make([] *SeleNameSpaceEnvUser_VR, 0)
	sql := fmt.Sprintf("select ns.id,ns.`name`,ns.remark from kube_env_user_namespace euns left join kube_namespace ns on (euns.namespace_id = ns.id)  where euns.user_id =? and ns.cluster_id = ?;")
	o.Raw(sql,userId,clusterId).QueryRow(&data)

	return &data
}


func  KubeEnvUserNameSpaceDeleteByUserId(userId string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeEnvUserNamespaceTBName())
	num, err := query.Filter("user_id", userId).Delete()
	return num, err
}