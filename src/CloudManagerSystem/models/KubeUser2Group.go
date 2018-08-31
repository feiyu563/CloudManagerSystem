package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

func KubeUser2GroupTBName() string {
	return "user2group"
}

type KubeUser2GroupQueryParam struct {
	BaseQueryParam
	UserId      	string
}

type KubeUser2Group struct {
	Id          	string `orm:"pk"`
	UserId      	string
	GroupId      	string
}

func (a *KubeUser2Group) TableName() string {
	return KubeUser2GroupTBName()
}

func KubeUser2GroupAllList(params *KubeUser2GroupQueryParam) ([]*KubeUser2Group, error) {
	query := orm.NewOrm().QueryTable(KubeUser2GroupTBName())
	var data []*KubeUser2Group
	if(strings.Compare(params.UserId,"")!=0){
		query = query.Filter("user_id", params.UserId)
	}
	_,err:=query.All(&data)
	if(err==nil){
		return data, nil
	}
	return nil, err
}

//批量删除
func KubeUser2GroupDeleteByUserId(userId string) (int64, error) {
	query := orm.NewOrm().QueryTable(KubeUser2GroupTBName())
	num, err := query.Filter("userId", userId).Delete()
	return num, err
}