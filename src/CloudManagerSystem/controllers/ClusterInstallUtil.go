package controllers


import (
	"fmt"
	"k8s.install/etcd"
	"k8s.install/k8s"
	"strings"
	"k8s.install/docker"
	"k8s.install/network"
	"k8s.install/dns"
	"golang.org/x/crypto/ssh"
	"bytes"
	"time"
	"net"
	"CloudManagerSystem/models"
	"k8s.install/host"
)

func exeFun(ch chan int,end_ch chan string,node *models.KubeHost,master *models.KubeHost,list []*models.KubeHost) {
	var isStart int
	isStart=0
	select {
	case num := <-ch :
		if(isStart==0){
			hostInit(node,master,list)
			num=num+1
			ch<-num
			isStart=1
		}
		if(num==10){
			end_ch<-"e"
		}
	default :
	}
}
func GetClient(node *models.KubeHost)*ssh.Client{
	var addr string
	addr="@IP:22"
	var client *ssh.Client
	config := &ssh.ClientConfig{
		User: node.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(node.PassWord),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	_add := strings.Replace(addr, "@IP", node.Ip, -1)
	client, err := ssh.Dial("tcp", _add, config)
	if err != nil {
		return nil
		//panic("Failed to dial: " + err.Error())
	}
	return client
}

func ExeCmd(client *ssh.Client,cmd string){
	if session, err := client.NewSession(); err == nil {
		defer session.Close()
		session.Run(cmd)
	}
}

func ExeCmdBack(client *ssh.Client,cmd string)(string,error){
	if session, err := client.NewSession(); err == nil {
		defer session.Close()
		//session.Run(cmd)
		buf, err := session.CombinedOutput(cmd)
		return string(buf), err
	}
	return "",nil
}

func ExeCmdList(client *ssh.Client,cmd string,sep string){
	cmdlist := strings.Split(cmd, sep)
	for _, cmd := range cmdlist {
		if session, err := client.NewSession(); err == nil {
			defer session.Close()
			session.Run(cmd)
		}
	}
}

func ExeShell(client *ssh.Client,cmd string,sep string)(*bytes.Buffer,error) {
	if session, err := client.NewSession(); err == nil {
		defer session.Close()
		var b bytes.Buffer
		session.Stdout = &b
		cmdlist := strings.Split(cmd, sep)
		stdinBuf, err := session.StdinPipe()
		if err != nil {
			//t.Error(err)
			return nil ,err
		}
		var outbt, errbt bytes.Buffer
		session.Stdout = &outbt
		session.Stderr = &errbt
		err = session.Shell()
		if err != nil {
			//t.Error(err)
			return nil ,err
		}
		for _, c := range cmdlist {
			c = c + "\n"
			stdinBuf.Write([]byte(c))
			time.Sleep(time.Duration(1)*time.Second)
		}
		stdinBuf.Close()
		session.Wait()
		//if(errbt.Len()>0){
		//	fmt.Println(errbt)
		//}
		return &errbt ,nil
	}
	return nil ,nil
}

//幂等性
func restartNode(list []*models.KubeHost){
	var cmd_initConfig string
	cmd_initConfig="shutdown -r now"
	for _,node := range list{
		client:=GetClient(node)
		if(client==nil){
			fmt.Println("当前机器%s无法连接",node.Ip)
			continue
		}
		defer client.Close()
		if session, err := client.NewSession(); err == nil {
			defer session.Close()
			session.Start(cmd_initConfig)
		}
	}
}

func CreateCluster(relations []*models.KubeHost) {
	startTime := time.Now()
	//var relations []*host.KubeHost
	var master *models.KubeHost
	//设置互信
	//设置系统环境变量
	//基础文件拷贝
	//找到master
	for _,node := range relations{
		if(node.IsInstallNode){
			master=node
			break;
		}
	}
	for _,node := range relations{
		//设置互信
		//设置环境变量
		setsTrust(node,master,relations)
	}
	//找到master
	for _,node := range relations{
		if(node.IsInstallNode){
			master=node
			//k8s 启动文件初始化
			//etcd启动文件初始化
			//k8s证书
			//etcd证书
			MasterInit(master,relations)
			break;
		}
	}

	for _,node := range relations{
		if(strings.Compare(node.Role,"master")==0){
			//分发etcd  并安装
			//分发k8s   并安装
			AllMasterInit(node,master,relations)
		}
	}
	for _,node := range relations{
		if(strings.Compare(node.Role,"node")==0){
			hostInit(node,master,relations)
		}
	}
	//重启所有master
	if(INSTALL_RESTARTE){
		restartNode(relations)
		fmt.Println("等待30秒	开始")
		time.Sleep(time.Duration(30)*time.Second)
		fmt.Println("等待30秒	结束")
	}
	//启动所有master服务
	for _,node := range relations{
		client:=GetClient(node)
		if(client==nil){
			fmt.Println("当前机器%s无法连接",node.Ip)
			return
		}
		defer client.Close()
		if(strings.Compare(node.Role,"master")==0){
			startMasterService(node)
		}
	}
	endTime := time.Now()
	cha:=endTime.Minute()-startTime.Minute()
	fmt.Println("===send====%d==========",cha)
}
const INSTALL_RESTARTE  = false
const INSTALL_BASE  = false
const INSTALL_ETCD  = false
const INSTALL_K8S  = false
const INSTALL_K8S_RUN  = false
const INSTALL_DOCKER  = true

func MasterInit(node *models.KubeHost,list []*models.KubeHost){
	client:=GetClient(node)
	if(client==nil){
		fmt.Println("当前机器%s无法连接",node.Ip)
		return
	}
	defer client.Close()
	//安装cfssl  ok
	host_InstallCert:=host.InstallCert()
	ExeShell(client, host_InstallCert, ";")
	if(INSTALL_ETCD){
		//etcd   文件初始化  /root/etcd/cmd/etcd     ok
		etcd_MasterInit:=etcd.MasterInit()
		ExeShell(client, etcd_MasterInit, "@@")

		//生成etcd证书  /root/etcd_tem   创建证书   ok
		etcd_CreateEtcdCert:=etcd.CreateEtcdCert(list)
		ExeCmd(client,etcd_CreateEtcdCert)
	}
	if(INSTALL_K8S){
		////k8s 文件初始化  ok  /root/kubenetes/server  /root/kubenetes/node
		k8s_InitMaster:=k8s.InitMaster()
		ExeCmdList(client,k8s_InitMaster,"@@")
		//生成k8s证书  cd /root/kubernets_tem   ok   证书  config
		k8s_CreateCert:=k8s.CreateCert(list)
		ExeCmd(client,k8s_CreateCert)
		//生成config文件		ok
		k8s_InitK8sEnv:=k8s.InitK8sEnv()
		ExeShell(client,k8s_InitK8sEnv,"@@")
	}
}

func AllMasterInit(node *models.KubeHost,master *models.KubeHost,list []*models.KubeHost){
	client:=GetClient(node)
	if(client==nil){
		fmt.Println("当前机器%s无法连接",node.Ip)
		return
	}
	defer client.Close()
	if(INSTALL_DOCKER){
		//docker  ok
		docker_CopyDocker:=docker.CopyDocker(master)
		ExeCmd(client,docker_CopyDocker)
		//拷贝网络组件  ok
		network_CopyFile:=network.CopyFile(master)
		ExeCmd(client,network_CopyFile)
		//拷贝dns组件  ok
		dns_CopyFile:=dns.CopyFile(master)
		ExeCmd(client,dns_CopyFile)
	}
	if(INSTALL_ETCD){
		//拷贝etcd   /root/etcd_tem
		etcd_CopyEtcd:=etcd.CopyEtcd(master)
		ExeCmd(client,etcd_CopyEtcd)

		//添加用户
		etcd_AddEtcdUser:=etcd.AddEtcdUser()
		ExeCmd(client,etcd_AddEtcdUser)

		//设置etc环境变量
		etcd_InitEtcdEnv:=etcd.InitEtcdEnv()
		ExeCmd(client,etcd_InitEtcdEnv)

		//安装etcd
		//node *host.KubeHost,master *host.KubeHost,list []*host.KubeHost
		etcd_InstallEtcd:=etcd.InstallEtcd(node,master,list)
		fmt.Println(etcd_InstallEtcd)
		ExeCmd(client,etcd_InstallEtcd)

		//拷贝etcd证书   ok
		etcd_CopyEtcdCert:=etcd.CopyEtcdCert()
		ExeShell(client,etcd_CopyEtcdCert,"@@")

		//etcd授权
		etcd_InitEtcdRole:=etcd.InitEtcdRole()
		fmt.Println(etcd_InitEtcdRole)
		ExeCmd(client,etcd_InitEtcdRole)
	}
	if(INSTALL_K8S){
		//添加用户
		k8s_AddK8sUser:=k8s.AddK8sUser(node,master)
		ExeCmd(client,k8s_AddK8sUser)

		//拷贝/root/kubernets_tem   /root/kubenetes/server   /root/kubenetes/node
		k8s_Copyk8sFile:=k8s.Copyk8sFile(node,master)
		ExeCmd(client,k8s_Copyk8sFile)

		//拷贝k8s主程序到  /usr/bin    @todo   拷贝/usr/lib/systemd/system/
		k8s_CopyK8sCmd:=k8s.CopyK8sCmd(node)
		ExeShell(client,k8s_CopyK8sCmd,"@@")

		//分发k8s证书  /etc/kubernetes/ssl
		k8s_CopyK8sCert:=k8s.CopyK8sCert(master)
		ExeCmd(client,k8s_CopyK8sCert)

		//拷贝k8s  /etc/kubernetes/  下的所有配置文件
		k8s_CopyK8sEnv:=k8s.CopyK8sEnv(node,master,list)
		ExeCmdList(client,k8s_CopyK8sEnv,"@@")

		//所有的node节点都拷贝    /etc/kubernetes/kubelet    /etc/kubernetes/proxy
		k8s_InitK8sNodeEnv:=k8s.InitK8sNodeEnv(node,master)
		ExeCmdList(client,k8s_InitK8sNodeEnv,"@@")

		//创建k8s目录
		k8s_CreateK8sPath:=k8s.CreateK8sPath()
		ExeCmd(client,k8s_CreateK8sPath)
		//授权
		k8s_InitK8sRole:=k8s.InitK8sRole(node,master)
		ExeCmd(client,k8s_InitK8sRole)
	}
}
func setsTrust(node *models.KubeHost,master *models.KubeHost,list []*models.KubeHost){
	if(!INSTALL_BASE){
		return
	}
	client:=GetClient(node)
	if(client==nil){
		fmt.Println("当前机器%s无法连接",node.Ip)
		return
	}
	defer client.Close()
	//设置hosts
	var cmd_setHosts, cmd_hostName string
	cmd_setHosts="if [ -z `cat /etc/hosts|grep '@IP'` ]; then echo '@IP @HostName'>>/etc/hosts;fi"
	for _, value := range list {
		setHosts := strings.Replace(cmd_setHosts, "@IP", value.Ip, -1)
		setHosts = strings.Replace(setHosts, "@HostName", value.HostName, -1)
		ExeCmd(client,setHosts)
	}
	//设置hostnams
	cmd_hostName = "hostnamectl set-hostname @HostName"
	hostName := strings.Replace(cmd_hostName, "@HostName", node.HostName, -1)
	ExeCmd(client,hostName)
	//生成秘钥
	host_CreateSskey:=host.CreateSskey()
	ExeCmd(client,host_CreateSskey)
	//分发key
	var cmd_keyCopy string
	cmd_keyCopy = "sshpass -p '@PassWord' ssh-copy-id -i ~/.ssh/id_rsa.pub -o StrictHostKeyChecking=no @User@@IP"
	for _, value := range list {
		keyCopy := strings.Replace(cmd_keyCopy, "@PassWord", value.PassWord, -1)
		keyCopy = strings.Replace(keyCopy, "@User", value.User, -1)
		keyCopy = strings.Replace(keyCopy, "@IP", value.Ip, -1)
		if session, err := client.NewSession(); err == nil {
			defer session.Close()
			session.Run(keyCopy)
		}
	}
	//执行环境变量  ok
	host_InitConfig :=host.InitConfig(master)
	ExeCmd(client,host_InitConfig)
}

func hostInit(node *models.KubeHost,master *models.KubeHost,list []*models.KubeHost){
	client:=GetClient(node)
	if(client==nil){
		fmt.Println("当前机器%s无法连接",node.Ip)
		return
	}
	defer client.Close()
	if(INSTALL_BASE){
		//设置hosts
		var cmd_setHosts, cmd_hostName string
		cmd_setHosts="if [ -z `cat /etc/hosts|grep '@IP'` ]; then echo '@IP @HostName'>>/etc/hosts;fi"
		for _, value := range list {
			setHosts := strings.Replace(cmd_setHosts, "@IP", value.Ip, -1)
			setHosts = strings.Replace(setHosts, "@HostName", value.HostName, -1)
			ExeCmd(client,setHosts)
		}
		//设置hostnams
		cmd_hostName = "hostnamectl set-hostname @HostName"
		hostName := strings.Replace(cmd_hostName, "@HostName", node.HostName, -1)
		ExeCmd(client,hostName)
		//生成秘钥
		host_CreateSskey:=host.CreateSskey()
		ExeCmd(client,host_CreateSskey)
		//分发key
		var cmd_keyCopy string
		cmd_keyCopy = "sshpass -p '@PassWord' ssh-copy-id -i ~/.ssh/id_rsa.pub -o StrictHostKeyChecking=no @User@@IP"
		for _, value := range list {
			keyCopy := strings.Replace(cmd_keyCopy, "@PassWord", value.PassWord, -1)
			keyCopy = strings.Replace(keyCopy, "@User", value.User, -1)
			keyCopy = strings.Replace(keyCopy, "@IP", value.Ip, -1)
			if session, err := client.NewSession(); err == nil {
				defer session.Close()
				session.Run(keyCopy)
			}
		}
		//执行环节设置  ok
		host_InitConfig :=host.InitConfig(master)
		ExeCmd(client,host_InitConfig)
	}
	if(INSTALL_DOCKER){
		//docker  ok
		docker_CopyDocker:=docker.CopyDocker(master)
		ExeCmdList(client,docker_CopyDocker,"@@")
		//拷贝网络组件  ok
		network_CopyFile:=network.CopyFile(master)
		ExeCmd(client,network_CopyFile)
		//拷贝dns组件  ok
		dns_CopyFile:=dns.CopyFile(master)
		ExeCmd(client,dns_CopyFile)
	}
	if(INSTALL_ETCD){
		//etcd
		//拷贝etcd
		etcd_CopyEtcd:=etcd.CopyEtcd(master)
		ExeCmdList(client,etcd_CopyEtcd,"@@")
		//拷贝etcd证书   ok
		etcd_CopyEtcdCert:=etcd.CopyEtcdCert()
		ExeShell(client,etcd_CopyEtcdCert,"@@")
	}
	if(INSTALL_K8S){
		///root/kubenetes/server  /root/kubenetes/node   ok
		//拷贝k8s主程序
		host_CopyFile:=host.CopyFile(node,master)
		ExeCmdList(client,host_CopyFile,"@@")

		//拷贝k8s主程序到  /usr/bin    ok
		k8s_CopyK8sCmd:=k8s.CopyK8sCmd(node)
		ExeShell(client,k8s_CopyK8sCmd,"@@")

		//分发k8s证书  ok
		k8s_CopyK8sCert:=k8s.CopyK8sCert(master)
		ExeCmd(client,k8s_CopyK8sCert)
		//拷贝k8s  env  ok
		k8s_CopyK8sEnv:=k8s.CopyK8sEnv(node,master,list)
		ExeCmdList(client,k8s_CopyK8sEnv,"@@")
		//所有的node节点都拷贝
		k8s_InitK8sNodeEnv:=k8s.InitK8sNodeEnv(node,master)
		ExeCmdList(client,k8s_InitK8sNodeEnv,"@@")
	}
}

func startMasterService(node *models.KubeHost){
	client:=GetClient(node)
	if(client==nil){
		fmt.Println("当前机器%s无法连接",node.Ip)
		return
	}
	defer client.Close()
	daemon_reload :="systemctl daemon-reload"
	ExeCmd(client,daemon_reload)
	if(INSTALL_DOCKER){
		//启动docker
		startdocker :="systemctl restart docker"
		ExeCmd(client,startdocker)
		//加载镜像
		reloadImages:="if [ -z `docker images|grep nginx|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i nginx_*.tar; fi@@" +
			"if [ -z `docker images|grep mritd/demo|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i mritd_demo.tar; fi@@"+
			"if [ -z `docker images|grep gcr.io/google_containers/pause-amd64|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i gcr.io_google_containers_pause-amd64_*.tar; fi"
		ExeCmdList(client,reloadImages,"@@")
		//加载网络镜像
		network_DockerReloadImages:=network.DockerReloadImages(node)
		ExeCmd(client,network_DockerReloadImages)
		//dns
		dns_DockerReloadImages:=dns.DockerReloadImages()
		ExeShell(client,dns_DockerReloadImages,"@@")
	}
	//启动程序 etcd
	if(INSTALL_ETCD){
		startCmd :="systemctl restart etcd"
		ExeCmd(client,startCmd)
	}
	//启动k8s
	if(INSTALL_K8S_RUN){
		k8s_RestartService:=k8s.RestartMasterService(node)
		ExeCmdList(client,k8s_RestartService,"&&")
		//创建
		createbootstrap:="kubectl create clusterrolebinding kubelet-bootstrap --clusterrole=system:node-bootstrapper --user=kubelet-bootstrap"
		ExeCmd(client,createbootstrap)
		//if(node.IsInstallNode){
		//	//创建calico
		//	//创建nds
		//	dns_K8sCreate:=dns.K8sCreate()
		//	ExeCmd(client,dns_K8sCreate)
		//}
	}
}
func startNodeService(node *models.KubeHost,list []*models.KubeHost){
	client:=GetClient(node)
	if(client==nil){
		fmt.Println("当前机器%s无法连接",node.Ip)
		return
	}
	defer client.Close()
	if(INSTALL_DOCKER){
		//启动docker
		startdocker :="systemctl restart docker"
		ExeCmd(client,startdocker)
		//加载镜像
		reloadImages:="if [ -z `docker images|grep nginx|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i nginx_*.tar; fi@@" +
			"if [ -z `docker images|grep mritd/demo|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i mritd_demo.tar; fi@@"+
			"if [ -z `docker images|grep gcr.io/google_containers/pause-amd64|awk 'NR >0{print $1}'` ]; then cd /root/dockerImages &&docker load -i gcr.io_google_containers_pause-amd64_*.tar; fi"
		ExeCmdList(client,reloadImages,"@@")
		//加载网络镜像
		network_DockerReloadImages:=network.DockerReloadImages(node)
		ExeCmd(client,network_DockerReloadImages)
		//dns
		dns_DockerReloadImages:=dns.DockerReloadImages()
		ExeShell(client,dns_DockerReloadImages,"@@")
	}
	//启动k8s
	if(INSTALL_K8S){
		k8s_RestartService:=k8s.RestartNodeService(node)
		ExeCmd(client,k8s_RestartService)
	}
	////启动nginx
	////server 172.16.20.2:6443 weight=20 max_fails=1 fail_timeout=10s;
	//var buf bytes.Buffer
	//for _,node:=range list{
	//	if(strings.Compare(node.Role,"master")==0){
	//		//"192.168.60.39",
	//		buf.WriteString("server "+node.Ip+":6443 weight=20 max_fails=1 fail_timeout=10s;\n")
	//	}
	//}
	//starnginx:="rm -rf /etc/nginx @@"+
	//	"mkdir -p /etc/nginx @@"+
	//	"cp /root/kubernets_tem/config/nginx.conf /etc/nginx @@"+
	//	"sed -i s/@CLUSTER_URL/@CLUSTER_IPS/g /etc/kubernetes/proxy @@"+
	//	"chmod +r /etc/nginx/nginx.conf"
	//starnginx = strings.Replace(starnginx, "@CLUSTER_IPS", buf.String(), -1)
	//ExeCmdList(client,starnginx,"@@")
	////docker 启动nginx
	//startNginx:="docker run -it -d -p 127.0.0.1:6443:6443 -v /etc/nginx:/etc/nginx  --name nginx-proxy --net=host --restart=on-failure:5 --memory=512M  nginx:1.13.5-alpine"
	//ExeCmd(client,startNginx)
}
