package host

import (
	"strings"
	"CloudManagerSystem/models"
)
//func CopyFile(client *ssh.Client,node *KubeHost,master *KubeHost){
//	//  拷贝文件/root/cmd/
//	var copyFile,cpoyK8s string
//	copyFile="if [ ! -e /root/cmd/install_init_config.sh ]; then scp -r @MASTER_USER@@MASTER_IP:/root/cmd /root; fi"
//	copyFile = strings.Replace(copyFile, "@MASTER_IP", master.Ip, -1)
//	copyFile = strings.Replace(copyFile, "@MASTER_USER", master.User, -1)
//	sshUtil.ExeCmd(client,copyFile)
//	if(strings.Compare(node.Role,"master")==0){
//		cpoyK8s="if [ ! -d /root/kubenetes/server ]; then mkdir -p /root/kubenetes &&scp -r @MASTER_USER@@MASTER_IP:/root/kubenetes/server /root/kubenetes/;fi"
//		cpoyK8s = strings.Replace(cpoyK8s, "@MASTER_IP", master.Ip, -1)
//		cpoyK8s = strings.Replace(cpoyK8s, "@MASTER_USER", master.User, -1)
//		sshUtil.ExeCmd(client,cpoyK8s)
//	}else{
//		cpoyK8s="if [ ! -d /root/kubenetes/node ]; then mkdir -p /root/kubenetes &&scp -r @MASTER_USER@@MASTER_IP:/root/kubenetes/node /root/kubenetes/;fi"
//		cpoyK8s = strings.Replace(cpoyK8s, "@MASTER_IP", master.Ip, -1)
//		cpoyK8s = strings.Replace(cpoyK8s, "@MASTER_USER", master.User, -1)
//		sshUtil.ExeCmd(client,cpoyK8s)
//	}
//}

func CopyFile(node *models.KubeHost,master *models.KubeHost)string{
	//  拷贝文件/root/cmd/
	var copyFile,cpoyK8s string
	copyFile="if [ ! -e /root/cmd/install_init_config.sh ]; then scp -r @MASTER_USER@@MASTER_IP:/root/cmd /root; fi"
	copyFile = strings.Replace(copyFile, "@MASTER_IP", master.Ip, -1)
	copyFile = strings.Replace(copyFile, "@MASTER_USER", master.User, -1)
	//sshUtil.ExeCmd(client,copyFile)
	if(strings.Compare(node.Role,"master")==0){
		cpoyK8s=" if [ ! -d /root/kubenetes/server ]; then mkdir -p /root/kubenetes &&scp -r @MASTER_USER@@MASTER_IP:/root/kubenetes/server /root/kubenetes/;fi"
		cpoyK8s = strings.Replace(cpoyK8s, "@MASTER_IP", master.Ip, -1)
		cpoyK8s = strings.Replace(cpoyK8s, "@MASTER_USER", master.User, -1)
		//sshUtil.ExeCmd(client,cpoyK8s)
		return copyFile+"@@"+cpoyK8s
	}else{
		cpoyK8s="if [ ! -d /root/kubenetes/node ]; then mkdir -p /root/kubenetes &&scp -r @MASTER_USER@@MASTER_IP:/root/kubenetes/node /root/kubenetes/;fi"
		cpoyK8s = strings.Replace(cpoyK8s, "@MASTER_IP", master.Ip, -1)
		cpoyK8s = strings.Replace(cpoyK8s, "@MASTER_USER", master.User, -1)
		//sshUtil.ExeCmd(client,cpoyK8s)
		return copyFile+"@@"+cpoyK8s
	}
}




//满足幂等性 ok
//func InitHost(host *KubeHost, list []*KubeHost)string {
//	var cmd_cat, cmd_echo, cmd_hostName string
//	cmd_cat = "cat /etc/hosts|grep '@IP'"
//	cmd_echo = ""
//	cmd_hostName = "hostnamectl set-hostname @HostName"
//
//	www:="if [ -z `cat /etc/hosts|grep '@IP'` ]; then echo '@IP @HostName'>>/etc/hosts;fi"
//
//	for _, value := range list {
//		echo := strings.Replace(cmd_echo, "@IP", value.Ip, -1)
//		echo = strings.Replace(echo, "@HostName", value.HostName, -1)
//		cat := strings.Replace(cmd_cat, "@IP", value.Ip, -1)
//		var flg bool
//		flg = true
//		if session, err := client.NewSession(); err == nil {
//			defer session.Close()
//			var b bytes.Buffer
//			session.Stdout = &b
//			session.Run(cat)
//			//fmt.Println(b.String())
//			if b.Len() > 0 {
//				flg = false
//			}
//		}
//		if session, err := client.NewSession(); flg && err == nil {
//			defer session.Close()
//			session.Run(echo)
//		}
//	}
//	//设置本机hostName
//	hostName := strings.Replace(cmd_hostName, "@HostName", host.HostName, -1)
//	if session, err := client.NewSession(); err == nil {
//		defer session.Close()
//		session.Run(hostName)
//	}
//}


//满足幂等性   执行ok
//func Sskey(client *ssh.Client, host *KubeHost, list []*KubeHost) {
//	var cmd_yum, cmd_createKey, cmd_keyCopy string
//	//cmd_keygen =`ssh-keygen -t rsa -P "" -f /root/.ssh/id_rsa`
//	cmd_createKey = `if [ ! -e /root/.ssh/id_rsa ]; then ssh-keygen -t rsa -P "" -f /root/.ssh/id_rsa; fi`
//	if session, err := client.NewSession(); err == nil {
//		defer session.Close()
//		//var b bytes.Buffer
//		//session.Stdout = &b
//		session.Run(cmd_createKey)
//	}
//	//当前节点安装sshpass
//	cmd_yum = "yum install -y sshpass"
//	if session, err := client.NewSession(); err == nil {
//		defer session.Close()
//		session.Run(cmd_yum)
//	}
//	//分发key
//	cmd_keyCopy = "sshpass -p '@PassWord' ssh-copy-id -i ~/.ssh/id_rsa.pub -o StrictHostKeyChecking=no @User@@IP"
//	for _, node := range list {
//		keyCopy := strings.Replace(cmd_keyCopy, "@PassWord", node.PassWord, -1)
//		keyCopy = strings.Replace(keyCopy, "@User", node.User, -1)
//		keyCopy = strings.Replace(keyCopy, "@IP", node.Ip, -1)
//		//fmt.Println(keyCopy)
//		if session, err := client.NewSession(); err == nil {
//			defer session.Close()
//			session.Run(keyCopy)
//		}
//	}
//}

func CreateSskey()string {
	return  `if [ ! -e /root/.ssh/id_rsa ]; then ssh-keygen -t rsa -P "" -f /root/.ssh/id_rsa &&yum install -y sshpass; fi`
}

func InstallCert()string {
	var cmd string
	cmd = `chmod +x /root/cfssl/cfssl_linux-amd64;
chmod +x /root/cfssl/cfssljson_linux-amd64;
chmod +x /root/cfssl/cfssl-certinfo_linux-amd64;
cp /root/cfssl/cfssl_linux-amd64 /root/cfssl/cfssl;
cp /root/cfssl/cfssljson_linux-amd64 /root/cfssl/cfssljson;
cp /root/cfssl/cfssl-certinfo_linux-amd64 /root/cfssl/cfssl-certinfo;
rm -rf /usr/bin/cfssl;
rm -rf /usr/bin/cfssljson;
rm -rf /usr/bin/cfssl-certinfo;
sudo mv /root/cfssl/cfssl /usr/bin/;
sudo mv /root/cfssl/cfssljson /usr/bin/;
sudo mv /root/cfssl/cfssl-certinfo /usr/bin/;`
	//创建证书命令
	return cmd
}

//幂等性   执行ok
func InitConfig(master *models.KubeHost)string{
	var cmd_initConfig string
	cmd_initConfig="if [ ! -e /root/cmd/install_init_config.sh ]; then scp -r @MASTER_USER@@MASTER_IP:/root/cmd /root &&sh /root/cmd/install_init_config.sh; fi"+
	" &&if [ `grep -c 'net.ipv4.tcp_tw_reuse = 1' /etc/sysctl.conf` -eq '0' ]; then cd /root/cmd &&sudo sh /root/cmd/install_init_config.sh; fi"+
	" &&if [ `grep -c 'export ETCDCTL_API=3' /etc/profile` -eq '0' ]; then export ETCDCTL_API=3 &&echo 'export ETCDCTL_API=3' >>/etc/profile; fi"
	cmd_initConfig = strings.Replace(cmd_initConfig, "@MASTER_IP", master.Ip, -1)
	cmd_initConfig = strings.Replace(cmd_initConfig, "@MASTER_USER", master.User, -1)
	return cmd_initConfig
}