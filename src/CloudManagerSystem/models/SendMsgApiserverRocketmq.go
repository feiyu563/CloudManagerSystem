package models


import (
	"github.com/sevenNt/rocketmq"
	//"time"
	"fmt"
	//"strconv"
	"github.com/gin-gonic/gin/json"
)

type UserVo struct {
	UID      string `json:"uid"`
	Name     string `json:"userName"`
	PassWord string `json:"passWord"`
	Oper     string `json:"oper"`
	Groups   string `json:"groups"`
}

func SendUserModifyMsgToApiserver(user *UserVo) error {
	//rocketmq   config
	group := "KUBE_TOPIC_CONSUMER"
	topic := "kube_topic"

	conf := &rocketmq.Config{
		Namesrv:   "rocketmq-cs.zxbike.top:32075",
		InstanceName: "DEFAULT",
	}
	//

	producer, err := rocketmq.NewDefaultProducer(group, conf)
	producer.Start()
	if err != nil {
		return err
	}

	//user.UID="123123"
	//user.Name="daolin"
	//user.PassWord="dadfasdf"
	//user.Oper="ADD"
	//user.Groups="1,2,3,4,5"
	data,_ :=json.Marshal(user)
	msg := rocketmq.NewMessage(topic, data)
	if sendResult, err := producer.Send(msg); err != nil {
		//fmt.Println("-----------------",err)
		return err
	} else {
		fmt.Println("Sync sending success!, ", sendResult)
	}

	return nil
}
