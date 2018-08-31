package test
//
//import (
//	"k8s.install/host"
//	sshUtil"k8s.install/sshUtil"
//	"fmt"
//	"k8s.install/etcd"
//	"k8s.install/k8s"
//	"strings"
//	"k8s.install/docker"
//	"k8s.install/network"
//	"k8s.install/dns"
//)
//func test_ssh2() {
//	var relations []*host.KubeHost
//	var master *host.KubeHost
//	kmaster1 := &host.KubeHost{
//		Ip:       "192.168.60.43",
//		HostName: "master1",
//		User:"root",
//		PassWord:"111111",
//		IsDeploy:false,
//		Role:"master",
//		IsInstallNode:true,
//	}
//	relations = append(relations, kmaster1)
//
//	kmaster2 := &host.KubeHost{
//		Ip:       "192.168.60.45",
//		HostName: "master2",
//		User:"root",
//		PassWord:"111111",
//		IsDeploy:false,
//		Role:"master",
//		IsInstallNode:false,
//	}
//	relations = append(relations, kmaster2)
//
//	kmaster3 := &host.KubeHost{
//		Ip:       "192.168.60.46",
//		HostName: "master3",
//		User:"root",
//		PassWord:"111111",
//		IsDeploy:false,
//		Role:"master",
//		IsInstallNode:false,
//	}
//	relations = append(relations, kmaster3)
//
//	knode1 := &host.KubeHost{
//		Ip:       "192.168.60.42",
//		HostName: "node1",
//		User:"root",
//		PassWord:"111111",
//		IsDeploy:false,
//		Role:"node",
//		IsInstallNode:false,
//	}
//	relations = append(relations, knode1)
//	//找到master
//	for _,node := range relations{
//		if(node.IsInstallNode){
//			master=node
//			MasterInit(master,relations)
//			break;
//		}
//	}
//	//for _,node := range relations{
//	//	go hostInit(node,master,relations)
//	//}
//	//if(master.IsInstallNode){
//	//	client:=sshUtil.GetClient(master)
//	//	if(client==nil){
//	//		fmt.Println("当前机器%s无法连接",master.Ip)
//	//		return
//	//	}
//	//	defer client.Close()
//	//	k8s.InitMaster(client)
//	//}
//	//所有机器重启
//	//startAllService(master,relations)
//}
//
//func MasterInit(node *host.KubeHost,list []*host.KubeHost){
//	client:=sshUtil.GetClient(node)
//	if(client==nil){
//		fmt.Println("当前机器%s无法连接",node.Ip)
//		return
//	}
//	defer client.Close()
//	//安装cfssl
//	host.InstallCert(client)
//	//etcd   文件初始化  /root/etcd/etcd_tem/etcd  -->/root/etcd_tem
//	etcd.MasterInit(client)
//	//生成etcd证书  /root/etcd_tem   创建证书
//	etcd.CreateEtcdCert(client,list)
//	//k8s 文件初始化
//	k8s.InitMaster(client)
//	//生成k8s证书  cd /root/kubernets_tem
//	k8s.CreateCert(client,list)
//}
//func AllMasterInit(node *host.KubeHost,master *host.KubeHost,list []*host.KubeHost){
//	//安装etcd
//
//}
//func hostInit(node *host.KubeHost,master *host.KubeHost,list []*host.KubeHost){
//	client:=sshUtil.GetClient(node)
//	if(client==nil){
//		fmt.Println("当前机器%s无法连接",node.Ip)
//		return
//	}
//	defer client.Close()
//	host.InitHost(client,node,list)
//	host.Sskey(client,node,list)
//	///root/kubenetes/server  /root/kubenetes/node
//	host.CopyFile(client,node,master)
//	host.InitConfig(client)
//	//docker
//	docker.CopyDocker(client,master)
//	//etcd
//	//分发etcd证书
//	etcd.CopyEtcdCert(client,master)
//	//分发k8s证书
//	k8s.CopyK8sCert(client,master)
//	//拷贝k8s主程序到  /usr/bin
//	k8s.CopyK8sCmd(client,node,master)
//	//拷贝网络组件
//	network.CopyFile(client,master)
//	//拷贝dns组件
//	dns.CopyFile(client,master)
//}
//
//func startAllService(master *host.KubeHost,list []*host.KubeHost){
//	for _,node := range list{
//		if(strings.Compare(node.Role,"master")==0){
//			startMasterService(node)
//		} else{
//			startNodeService(node)
//		}
//	}
//}
//func afterStartService(){
//
//}
//func startMasterService(node *host.KubeHost){
//	client:=sshUtil.GetClient(node)
//	if(client==nil){
//		fmt.Println("当前机器%s无法连接",node.Ip)
//		return
//	}
//	defer client.Close()
//	//启动docker
//	startdocker :="systemctl restart docker"
//	sshUtil.ExeCmd(client,startdocker)
//	//加载镜像
//	reloadImages:="if [ -z `docker images|grep nginx|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i nginx_*.tar; fi@@" +
//		"if [ -z `docker images|grep mritd/demo|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i mritd_demo.tar; fi@@"+
//		"if [ -z `docker images|grep gcr.io/google_containers/pause-amd64|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i gcr.io_google_containers_pause-amd64_*.tar; fi"
//	sshUtil.ExeCmdList(client,reloadImages,"@@")
//	//加载网络镜像
//	network.DockerReloadImages(client,node)
//	//启动程序 etcd
//	startCmd :="systemctl restart etcd"
//	sshUtil.ExeCmd(client,startCmd)
//	//启动k8s
//	k8s.RestartService(client,node)
//	//加载dns镜像
//	dns.DockerReloadImages(client)
//	//创建nds
//	if(node.IsInstallNode){
//		dns.K8sCreate(client)
//	}
//
//}
//func startNodeService(node *host.KubeHost){
//	client:=sshUtil.GetClient(node)
//	if(client==nil){
//		fmt.Println("当前机器%s无法连接",node.Ip)
//		return
//	}
//	defer client.Close()
//	//启动docker
//	startdocker :="systemctl restart docker"
//	sshUtil.ExeCmd(client,startdocker)
//	//加载镜像
//	reloadImages:="if [ -z `docker images|grep nginx|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i nginx_*.tar; fi@@" +
//		"if [ -z `docker images|grep mritd/demo|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i mritd_demo.tar; fi@@"+
//		"if [ -z `docker images|grep gcr.io/google_containers/pause-amd64|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i gcr.io_google_containers_pause-amd64_*.tar; fi"
//	sshUtil.ExeCmdList(client,reloadImages,"@@")
//	//加载网络镜像
//	network.DockerReloadImages(client,node)
//	//启动k8s
//	k8s.RestartService(client,node)
//	//加载dns镜像
//	dns.DockerReloadImages(client)
//}
//
//
