package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func KubeEnvUserClusterTBName() string {
	return "kube_env_user_cluster"
}

type KubeEnvUserClusterQueryParam struct {
	BaseQueryParam
}

type KubeEnvUserCluster struct {
	Id          	string `orm:"pk"`
	UserId      	string
	ClusterId      	string
	ClusterName     string `orm:"-"`
}

func (a *KubeEnvUserCluster) TableName() string {
	return KubeEnvUserClusterTBName()
}

func EnvUserCluster(userId string) (*KubeEnvUserCluster, error) {
	o := orm.NewOrm()
	var u KubeEnvUserCluster
	sql := fmt.Sprintf(`SELECT euc.id,euc.cluster_id,kc.name cluster_name 
		FROM %s euc 
		INNER JOIN %s kc on kc.id=euc.cluster_id 
		WHERE euc.user_id=%s`, KubeEnvUserClusterTBName(), KubeClusterTBName(),userId)
	err := o.Raw(sql).QueryRow(&u)
	if(err!=nil){
		return nil, err
	}
	return &u, nil
}

type SeleClusterEnvUser_VR struct {
	Id         string
	Name       string
	Remark     string
	Selected   string
}

//获取环境变量的值+
func SeleClusterEnvUser(userId string)([] *SeleClusterEnvUser_VR){
	o := orm.NewOrm()

	data := make([] *SeleClusterEnvUser_VR, 0)
	sql := fmt.Sprintf("CALL proc_EnvUserCluster (?)")
	o.Raw(sql,userId).QueryRows(&data)

	return data
}

func  KubeEnvUserClusterDeleteByUserId(userId string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeEnvUserClusterTBName())
	num, err := query.Filter("user_id", userId).Delete()
	return num, err
}
