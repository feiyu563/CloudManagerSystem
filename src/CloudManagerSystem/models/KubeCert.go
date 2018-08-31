package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"fmt"
)

func init() {
	orm.RegisterModel(new(KubeCert))
}

func KubeCertTBName() string {
	return "kube_cert"
}

type KubeCertQueryParam struct {
	BaseQueryParam
}

type KubeCert struct {
	Id        string `orm:"pk"`
	KeyName   string `orm:"size(40)"`   //       varchar(10),
	CertValue string `orm:"size(4096)"` //varchar(255) comment '证书',
	ClusterId string `orm:"size(40)"`   //varchar(40) comment '集群id',
}

func (a *KubeCert) TableName() string {
	return KubeCertTBName()
}

func TT() {
	buf, err := ioutil.ReadFile("/root/work/src/admin.pem")
	KubeCertAddtest(string(buf), "admin")
	if err != nil {
		fmt.Println(err)
	}
	bufkey, err := ioutil.ReadFile("/root/work/src/admin-key.pem")
	KubeCertAddtest(string(bufkey), "admin-key")

	bufca, err := ioutil.ReadFile("/root/work/src/k8s-root-ca.pem")
	KubeCertAddtest(string(bufca), "k8s-root-ca")
}

func KubeCertAddtest(certvalue string, keyname string) {
	m := KubeCert{}
	o := orm.NewOrm()

	u4, _ := uuid.NewV4()
	m.Id = u4.String()
	m.CertValue = certvalue
	m.KeyName = keyname
	m.ClusterId = "96724a68-d10a-49f6-b852-b3c1d053e238"

	_, err := o.Insert(&m)
	if err == nil {

	}
}

func GetCertDataByCluster(ClusterId string, KeyName string) string{//[]*KubeCert {

	query := orm.NewOrm().QueryTable(KubeCertTBName())
	data := make([]*KubeCert, 0)
	total ,_ :=query.Filter("cluster_id", ClusterId).Filter("key_name", KeyName).All(&data) //check master
	if total == 0 {
		return ""
	}

	return data[0].CertValue
}
