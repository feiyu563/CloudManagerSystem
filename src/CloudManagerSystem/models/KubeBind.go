package models

import "github.com/astaxie/beego/orm"

func KubeBindTBName() string {
	return "kube_bind"
}

type KubeBindQueryParam struct {
	BaseQueryParam
}

type KubeBind struct {
	Id          	string `orm:"pk"`
	Name			string
	UserId      	string
	UserType		int
	ClusterId      	string
	//ClusterName		string	`orm:"-"`
	NamespaceId		string
	//NamespaceName	string	`orm:"-"`
	RoleId      	string
	//RoleName		string	`orm:"-"`
}

func (a *KubeBind) TableName() string {
	return KubeBindTBName()
}

func FindKubeBindsByUserId(userId string) ([]*KubeBind, error) {
	o := orm.NewOrm()
	sql :="SELECT kb.id id,kb.name name,kb.user_id user_id,kb.user_type user_type,kb.cluster_id cluster_id,kc.name ClusterName,kb.namespace_id namespace_id,kn.name namespace_name,kb.role_id role_id,kr.name role_name from kube_bind kb INNER JOIN kube_cluster kc on kc.id=kb.cluster_id INNER JOIN kube_namespace kn on kn.cluster_id=kb.cluster_id and kn.id=kb.namespace_id INNER JOIN kube_role kr on kr.cluster_id=kb.cluster_id and kr.id=kb.role_id where kb.user_id=?"
	//var kubeBinds []KubeBind
	kubeBinds := make([]*KubeBind, 0)
	_, err := o.Raw(sql).SetArgs(userId).QueryRows(&kubeBinds)
	if err != nil {
	    return nil, err
	}
	return kubeBinds, nil
}

func KubeBindDeleteByUserId(userId string) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM kube_bind WHERE user_id =?" ).SetArgs(userId).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num, nil
	}
	return 0, err
}

//批量删除
func KubeBindDeleteByUserIds(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeBindTBName())
	num, err := query.Filter("user_id__in", ids).Delete()
	return num, err
}