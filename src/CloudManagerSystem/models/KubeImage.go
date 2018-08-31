package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

func KubeImageTBName() string {
	return "kube_image"
}

type KubeImage struct {
	Id        string `orm:"pk;size(40)"`
	Name      string `orm:"size(100)"`
	Tag       string `orm:"size(40)"`
	Env       string `orm:"size(100)"` //varchar(100) comment '环境变量',
	Runcmd    string `orm:"size(100)"` //varchar(100) comment '启动命令',
	Heartbeat string `orm:"size(100)"` //varchar(100) comment '监控检测',
}

func (a *KubeImage) TableName() string {
	return KubeImageTBName()
}

type KubeImageQueryParam struct {
	BaseQueryParam
	Id   string
	Name string
	Tag  string
}

func KubeImagePageList(params *KubeImageQueryParam) ([]*KubeImage, int64) {
	query := orm.NewOrm().QueryTable(KubeImageTBName())

	data := make([]*KubeImage, 0)
	//默认排序
	sortorder := "id"

	if strings.Compare(params.Sort, "") != 0 {
		sortorder = params.Sort
	}

	if params.Order == "desc" || params.Order == "asc" {
		sortorder = "-" + sortorder
	}

	//
	if strings.Compare(params.Name, "") != 0 {
		query = query.Filter("name__icontains", params.Name)
	}

	if strings.Compare(params.Id, "") != 0 {
		query = query.Filter("id", params.Id)
	}

	if strings.Compare(params.Tag, "") != 0 {
		query = query.Filter("tag__icontains", params.Tag)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

func KubeImageList() ([]*KubeImage, int64) {
	data := make([]*KubeImage, 0)

	query := orm.NewOrm().QueryTable(KubeImageTBName())
	total, _ := query.Count()
	query.All(&data)

	return data, total
}

//批量删除
func KubeImageDelete(ids []string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeImageTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
