package etcd

import (
	"bytes"
	"strings"
	"fmt"
	"CloudManagerSystem/models"
)
//所有集群
func copyFile(){
	//   拷贝/root/etcd_tem
}

const ETCD_SSL_CREATE_PATH  = "/root/install/ssl/etcd"
const ETCD_SSL_PATH  = "/etc/etcd/ssl"

func MasterInit()string{
	//判断etcd包是否解压  service
	//解压文件/root/etcd/cmd/etcd   /root/etcd/cmd/etcdctl
	var etcdService string
	etcdService ="if [ ! -e /root/etcd/cmd/etcd ]; then cd /root/etcd/ &&tar -zxvf  etcd-v*-linux-amd64.tar.gz &&mv etcd-v*-linux-amd64 service &&mv /root/etcd/service/etcd /root/etcd/cmd/ &&mv /root/etcd/service/etcdctl /root/etcd/cmd/ &&rm -rf /root/etcd/service; fi"
	//拷贝/root/etcd_tem
	etcdService1 :="if [ ! -d /root/etcd_tem ]; then mkdir -p /root/etcd_tem &&cp /root/etcd/cmd/* /root/etcd_tem/ &&cp /root/etcd/config/* /root/etcd_tem/ &&cp /root/etcd/ssl/* /root/etcd_tem/ &&chmod +x /root/etcd_tem/etcd.service; fi"
	return  etcdService+"@@"+etcdService1
}

//幂等性
func CreateEtcdCert(list []*models.KubeHost)string{
	//chmod +x /root/install/cfssl/cfssl_linux-amd64
	var cmd_createEtcdCert string
	cmd_createEtcdCert="if [ ! -e /root/etcd_tem/etcd.pem ]; then " +
		//"sed -i 's/@MASTER_IPS/@MasterIP/g' /root/etcd_tem/etcd-csr.json "+
		"cd  /root/etcd_tem " +
		"&&cfssl gencert -initca etcd-root-ca-csr.json | cfssljson -bare etcd-root-ca " +
		"&&cfssl gencert -ca=etcd-root-ca.pem -ca-key=etcd-root-ca-key.pem -config=etcd-gencert.json -profile=etcd -hostname='127.0.0.1,@HOST_LISTlocalhost' etcd-csr.json | cfssljson -bare etcd; fi"
	//修改证书中的master  ip 地址
	var buf,hostList bytes.Buffer
	for _,node:=range list{
		if(strings.Compare(node.Role,"master")==0){
			//"192.168.60.39",
			buf.WriteString("\""+node.Ip+"\",")
			hostList.WriteString(node.Ip+",")
		}
	}
	//判断是否需要创建证书
	//创建Etcd证书 /etc/etcd/ssl
	createEtcdCert := strings.Replace(cmd_createEtcdCert, "@MasterIP", buf.String(), -1)
	createEtcdCert = strings.Replace(createEtcdCert, "@HOST_LIST", hostList.String(), -1)
	fmt.Println(createEtcdCert)
	return createEtcdCert
}

//幂等性
func CopyEtcdCert()string{
	var cmd_copyEtcdCert string
	//需要校验是否有问题
	cmd_copyEtcdCert="if [ ! -e /etc/etcd/ssl/etcd.pem ]; then mkdir -p @ETCD_SSL_PATH &&cp /root/etcd_tem/*.pem @ETCD_SSL_PATH/ &&chmod -R 644 @ETCD_SSL_PATH/* &&chmod 755 @ETCD_SSL_PATH; fi"
	//分发所有的证书
	copyEtcdCert := strings.Replace(cmd_copyEtcdCert, "@ETCD_SSL_PATH", ETCD_SSL_PATH, -1)
	//copyEtcdCert = strings.Replace(copyEtcdCert, "@MASTER_USER", master.User, -1)
	//copyEtcdCert = strings.Replace(copyEtcdCert, "@MASTER_IP", master.Ip, -1)
	return copyEtcdCert
}

func CopyEtcd(master *models.KubeHost)string{
	//拷贝etcd安装包
	copyEtcdPackage:="if [ ! -d /root/etcd_tem ]; then scp -r @MASTER_USER@@MASTER_IP:/root/etcd_tem /root/; fi"
	copyEtcdPackage = strings.Replace(copyEtcdPackage, "@MASTER_IP", master.Ip, -1)
	copyEtcdPackage = strings.Replace(copyEtcdPackage, "@MASTER_USER", master.User, -1)
	return copyEtcdPackage
}
func InitEtcdEnv()string{
	setEnv:="systemctl stop firewalld &&systemctl disable firewalld &&export ETCDCTL_API=3 &&setenforce 0"
	return setEnv
}
func AddEtcdUser()string{
	addUser:="useradd etcd -p etcd"
	return addUser
}

func InitEtcdRole()string{
	setEnv:="export ETCDCTL_API=3"+
		" &&mkdir -p /var/lib/etcd &&chown -R etcd:etcd /var/lib/etcd" +
		" &&chown -R etcd:etcd /etc/etcd/ssl &&chmod -R 644 /etc/etcd/ssl/* &&chmod 755 /etc/etcd/ssl"+
	    " &&mkdir -p /var/lib/etcd/wal &&chown -R etcd:etcd /var/lib/etcd &&chown -R etcd:etcd /var/lib/etcd/wal"
	//		chmod 755 /etc/etcd/ssl
	//		chmod -R 644 /etc/etcd/ssl/*
	//chown -R etcd:etcd /var/lib/etcd
	//chown -R etcd:etcd /var/lib/etcd/wal
	return setEnv
}

func InstallEtcd(node *models.KubeHost,master *models.KubeHost,list []*models.KubeHost)string{
	//var cmd_getEtcdFile string
	var buf bytes.Buffer
	for _,node:=range list{
		//etcd1=https://192.168.60.43:2380
		if(strings.Compare(node.Role,"master")==0){
			if(buf.Len()>0){
				buf.WriteString(",")
			}
			etcdname:=strings.Replace(node.HostName,"master","etcd",-1)
			buf.WriteString(etcdname)
			buf.WriteString("=https:\\/\\/")
			//buf.WriteString("=https:\\\\/\\\\/")
			buf.WriteString(node.Ip)
			buf.WriteString(":2380")
		}
	}
	///etc/etcd/
	createEtcdPackage:="if [ ! -d /etc/etcd/ ]; then mkdir -p /etc/etcd/; fi"
	//拷贝config 到/etc/etcd/etcd.conf
	copyConToEnv:=" &&if [ ! -e /etc/etcd/etcd.conf ]; then cp /root/etcd_tem/etcd.conf /etc/etcd/ &&sed -i 's/@HOST_NAME/@ETCD_NAME/g' /etc/etcd/etcd.conf &&sed -i 's/@HOST_IP/@ETCD_IP/g' /etc/etcd/etcd.conf &&sed -i 's/@CLUSTER_URL/@CLUSTER_LIST/g' /etc/etcd/etcd.conf; fi"
	etcdname:=strings.Replace(node.HostName,"master","etcd",-1)
	copyConToEnv = strings.Replace(copyConToEnv, "@ETCD_NAME", etcdname, -1)
	copyConToEnv = strings.Replace(copyConToEnv, "@ETCD_IP", node.Ip, -1)
	copyConToEnv = strings.Replace(copyConToEnv, "@CLUSTER_LIST", buf.String(), -1)

	//拷贝启动程序
	copyCmdBin:=" &&if [ ! -e /usr/bin/etcd ]; then cp /root/etcd_tem/etcd /usr/bin/; fi"
	copyetcdctlCmdBin:=" &&if [ ! -e /usr/bin/etcdctl ]; then cp /root/etcd_tem/etcdctl /usr/bin/; fi"
	//拷贝启动程序if [ ! -e /lib/systemd/system/etcd.service ]; then cp /root/etcd/etcd.service /lib/systemd/system/; fi
	copyStartCmd:=" &&if [ ! -e /lib/systemd/system/etcd.service ]; then cp /root/etcd_tem/etcd.service /lib/systemd/system/; fi"
	return createEtcdPackage+copyConToEnv+copyCmdBin+copyetcdctlCmdBin+copyStartCmd
}



func clearEtcd() {
	//var cmd_clearEtcd string
	//cmd_clearEtcd ="if [ ! -e /usr/bin/etcd ]; then systemctl disable etcd &&systemctl stop etcd &&rm -rf /lib/systemd/system/etcd.service &&systemctl daemon-reload &&rm -rf /etc/etcd &&rm -rf /var/lib/etcd &&rm -rf /usr/bin/etcd &&rm -rf /usr/bin/etcdctl; fi"

}