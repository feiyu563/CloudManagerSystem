package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
)

type KubePrometheusController struct {
	BaseController
}

func (c *KubePrometheusController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("Namespaces", "Nodes")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}

type PrometheusResp struct {
	Status string `json:"status"`
	Data Data `json:"data"`
}
type Data struct {
	Result []Result `json:"result"`
}
type Result struct {
	Value []interface{} `json:"value"`
}
type ReplayChannel struct {
	statefulset chan string
	daemonset chan string
	cronjob chan string
	deployment chan string
	nodenum chan string
	podnum chan string
	netin chan string
	netout chan string
	cpu chan string
	mem chan string
	disk chan string
}
//集群汇总信息展示2s读取一次,cpu/mem/disk
func (c *KubePrometheusController) ProCpuMemDisk() {
	Prometheus_url := beego.AppConfig.String("Prometheus::prometheus_url")
	channels:=&ReplayChannel{
		cpu:GetPrometheusData(Prometheus_url+"api/v1/query?query=avg(100-(avg%20by%20(cpu)%20(irate(node_cpu%7Bmode%3D%22idle%22%7D%5B5m%5D))%20*100))"),
		mem:GetPrometheusData(Prometheus_url+"api/v1/query?query=((sum(node_memory_MemTotal)-sum(node_memory_MemFree)-sum(node_memory_Buffers)-sum(node_memory_Cached))%2Fsum(node_memory_MemTotal))*100"),
		disk:GetPrometheusData(Prometheus_url+"api/v1/query?query=((sum(node_filesystem_size%7Bdevice!%3D%22rootfs%22%7D)-sum(node_filesystem_free%7Bdevice!%3D%22rootfs%22%7D))%2Fsum(node_filesystem_size%7Bdevice!%3D%22rootfs%22%7D))*100"),
	}
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["cpu"]=<-channels.cpu
	result["mem"]=<-channels.mem
	result["disk"]=<-channels.disk
	c.Data["json"] = result
	c.ServeJSON()
}

//静态数据展示
func (c *KubePrometheusController) ProCount() {
	Prometheus_url := beego.AppConfig.String("Prometheus::prometheus_url")
	channels:=&ReplayChannel{
		statefulset:GetPrometheusData(Prometheus_url+"api/v1/query?query=count(kube_statefulset_created)"),
		daemonset:GetPrometheusData(Prometheus_url+"api/v1/query?query=count(kube_daemonset_created)"),
		cronjob:GetPrometheusData(Prometheus_url+"api/v1/query?query=count(kube_cronjob_created)"),
		deployment:GetPrometheusData(Prometheus_url+"api/v1/query?query=count(kube_deployment_created)"),
		nodenum:GetPrometheusData(Prometheus_url+"api/v1/query?query=count(kube_node_created)"),
		podnum:GetPrometheusData(Prometheus_url+"api/v1/query?query=sum(kubelet_running_pod_count)"),
		netin:GetPrometheusData(Prometheus_url+"api/v1/query?query=sum(instance%3Anode_network_receive_bytes%3Arate%3Asum)"),
		netout:GetPrometheusData(Prometheus_url+"api/v1/query?query=sum(instance%3Anode_network_transmit_bytes%3Arate%3Asum)"),
	}
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["statefulset"]=<-channels.statefulset
	result["daemonset"]=<-channels.daemonset
	result["cronjob"]=<-channels.cronjob
	result["deployment"]=<-channels.deployment
	result["nodenum"]=<-channels.nodenum
	result["podnum"]=<-channels.podnum
	result["netin"]=<-channels.netin
	result["netout"]=<-channels.netout
	c.Data["json"] = result
	c.ServeJSON()
}

func GetPrometheusData(url string) chan string{
	channel:=make(chan string)
	go func() {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, strings.NewReader(""))
		//req, err := http.NewRequest("GET", url, strings.NewReader(""))
		if err != nil {
			beego.Error(err)
			return
		}
		//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		//req.Header.Set("Cookie", "rem-username=admin; beegosessionID=3a7f697427caed8eb17f14368bb8c832")
		resp, err := client.Do(req)
		if err != nil {
			beego.Error(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			beego.Error(err)
			return
		}
		//fmt.Println(url,string(body))
		prometheusResp:=PrometheusResp{}
		json.Unmarshal(body,&prometheusResp)
		channel<-prometheusResp.Data.Result[0].Value[1].(string)
	}()
	return channel
}

