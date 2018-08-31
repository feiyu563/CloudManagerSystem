package k8s

import (
	"bytes"
	"strings"
	"strconv"
	"CloudManagerSystem/models"
)

const K8S_SSL_CREATE_PATH  = "/root/install/ssl/k8s"
const K8S_SSL_PATH  = "/etc/kubernetes/ssl"
const K8S_BIN_PATH  = "/usr/bin/etcd"


//
func AddK8sUser(node *models.KubeHost,master *models.KubeHost)string{
	return "useradd kube -p kube"
}
//所有节点
func Copyk8sFile(node *models.KubeHost,master *models.KubeHost)string{
	// /root/kubernets_tem
	var cmd_initK8sCmd string
	cmd_initK8sCmd ="if [ ! -d /root/kubernets_tem ]; then scp -r @MASTER_USER@@MASTER_IP:/root/kubernets_tem /root/;fi"
	if(strings.Compare(node.Role,"master")==0){
		cmd_initK8sCmd=cmd_initK8sCmd+" &&if [ ! -d /root/kubenetes/server ]; then mkdir -p /root/kubenetes &&scp -r @MASTER_USER@@MASTER_IP:/root/kubenetes/server /root/kubenetes/;fi"
		cmd_initK8sCmd = strings.Replace(cmd_initK8sCmd, "@MASTER_IP", master.Ip, -1)
		cmd_initK8sCmd = strings.Replace(cmd_initK8sCmd, "@MASTER_USER", master.User, -1)
	}else{
		cmd_initK8sCmd=cmd_initK8sCmd+" &&if [ ! -d /root/kubenetes/node ]; then mkdir -p /root/kubenetes &&scp -r @MASTER_USER@@MASTER_IP:/root/kubenetes/node /root/kubenetes/;fi"
		cmd_initK8sCmd = strings.Replace(cmd_initK8sCmd, "@MASTER_IP", master.Ip, -1)
		cmd_initK8sCmd = strings.Replace(cmd_initK8sCmd, "@MASTER_USER", master.User, -1)
	}
	return cmd_initK8sCmd
}

func InitMaster()string{
	//判断k8s包是否解压
	var cmd_initK8sCmd string
	cmd_initK8sCmd ="if [ ! -d /root/kubenetes/kubernetes/server ]; then cd /root/kubenetes/ &&tar -zxvf kubernetes-server-linux-amd64.tar.gz; fi@@" +
		//"if [ ! -d /root/kubenetes/kubernetes/node ]; then cd /root/kubenetes/ &&tar -zxvf kubernetes-node-linux-amd64.tar.gz; fi@@" +
		"if [ ! -d /root/kubenetes/server ]; then mkdir -p /root/kubenetes/server &&cp /root/kubenetes/kubernetes/server/bin/hyperkube /root/kubenetes/kubernetes/server/bin/kube-apiserver /root/kubenetes/kubernetes/server/bin/kube-scheduler /root/kubenetes/kubernetes/server/bin/kubelet /root/kubenetes/kubernetes/server/bin/kube-controller-manager /root/kubenetes/kubernetes/server/bin/kubectl /root/kubenetes/kubernetes/server/bin/kube-proxy /root/kubenetes/server; fi@@"
		//"if [ ! -d /root/kubenetes/node ]; then cp -r /root/kubenetes/kubernetes/node/bin  /root/kubenetes/node; fi"
	return cmd_initK8sCmd
}

func InitK8sRole(node *models.KubeHost,master *models.KubeHost)string{
	//"chown -R kube:kube /etc/kubernetes/ssl &&chmod -R 755 /etc/kubernetes/ssl" +
	//	" &&chown -R kube:kube /etc/kubernetes" +
	//	" &&chown -R kube:kube /var/log/kube-audit /usr/libexec/kubernetes" +
	//	" &&chmod -R 755 /var/log/kube-audit /usr/libexec/kubernetes"+
	//
	//	" &&chown -R kube:kube /var/run/kubernetes"+
	//	" &&chown -R kube:kube /etc/kubernetes/ssl"+
	//	" &&chmod -R 755 /etc/kubernetes/ssl"+
	//	" &&chmod -R 755 /var/lib/kubelet"+
	//	" &&chown -R kube:kube /etc/kubernetes"+
	//	" &&mkdir -p /var/log/kube-audit /usr/libexec/kubernetes"+
	//	" &&chown -R kube:kube /var/log/kube-audit /usr/libexec/kubernetes"+
	//	" &&chmod -R 755 /var/log/kube-audit /usr/libexec/kubernetes"+
	var cmd string
	cmd="mkdir -p /var/run/kubernetes &&chown -R kube:kube /var/run/kubernetes &&chmod -R 755 /var/run/kubernetes"+
		" &&mkdir -p /var/lib/kubelet &&chown -R kube:kube /var/lib/kubelet &&chmod -R 755 /var/lib/kubelet"+
		" &&chown -R kube:kube /etc/kubernetes/ssl &&chmod -R 755 /etc/kubernetes/ssl"+
		" &&chown -R kube:kube /etc/kubernetes &&chmod -R 755 /etc/kubernetes"+
		" &&mkdir -p /var/log/kube-audit /usr/libexec/kubernetes"+
		" &&chown -R kube:kube /var/log/kube-audit /usr/libexec/kubernetes"+
		" &&chmod -R 755 /var/log/kube-audit /usr/libexec/kubernetes"+
	" &&chmod 755 /usr/bin/kubelet &&chmod 755 /usr/bin/kube-proxy &&chmod 755 /usr/bin/kubectl &&chmod 755 /usr/bin/kube-apiserver &&chmod 755 /usr/bin/kube-controller-manager &&chmod 755 /usr/bin/kube-scheduler &&chmod 755 /usr/bin/hyperkube"
	return cmd
}

func CreateK8sPath()string{
	cmd:="mkdir -p /var/lib/kubelet &&mkdir -p /var/log/kube-audit /usr/libexec/kubernetes &&mkdir -p /var/run/kubernetes"
	return cmd
}

func InitK8sNodeEnv(node *models.KubeHost,master *models.KubeHost)string{
	// /root/kubernets_tem
	cmd_initK8sCmd :="if [ ! -e /etc/kubernetes/kubelet ]; then cp /root/kubernets_tem/config/kubelet /etc/kubernetes/ &&sed -i s/@HOST_IP/@IP/g /etc/kubernetes/kubelet &&sed -i s/@HOST_NAME/@NAME/g /etc/kubernetes/kubelet;fi@@"+
		"if [ ! -e /etc/kubernetes/proxy ]; then cp /root/kubernets_tem/config/proxy /etc/kubernetes/ &&sed -i s/@HOST_IP/@IP/g /etc/kubernetes/proxy &&sed -i s/@HOST_NAME/@NAME/g /etc/kubernetes/proxy;fi@@"+
		"if [ ! -e /lib/systemd/system/kubelet.service ]; then cp /root/kubernets_tem/cmd/kubelet.service /lib/systemd/system/;fi@@"+
		"if [ ! -e /lib/systemd/system/kube-proxy.service ]; then cp /root/kubernets_tem/cmd/kube-proxy.service /lib/systemd/system/;fi@@"
		//"sed -i s/@HOST_IP/@IP/g /etc/kubernetes/kubelet &&sed -i s/@HOST_NAME/@NAME/g /etc/kubernetes/kubelet@@"+
		//"sed -i s/@HOST_IP/@IP/g /etc/kubernetes/proxy &&sed -i s/@HOST_NAME/@NAME/g /etc/kubernetes/proxy@@"+
		//"mkdir -p /var/lib/kubelet"
	//cmd_initK8sCmd = strings.Replace(cmd_initK8sCmd, "@MASTER_IP", master.Ip, -1)
	//cmd_initK8sCmd = strings.Replace(cmd_initK8sCmd, "@MASTER_USER", master.User, -1)

	cmd_initK8sCmd = strings.Replace(cmd_initK8sCmd, "@IP", node.Ip, -1)
	cmd_initK8sCmd = strings.Replace(cmd_initK8sCmd, "@NAME", node.HostName, -1)
	return cmd_initK8sCmd
}
func InitK8sEnv()string {
	var cmd_initK8sEnv string
	cmd_initK8sEnv=`
if [ ! -e /usr/bin/kubectl ]; then cp /root/kubenetes/server/kubectl /usr/bin/ &&chmod +x /usr/bin/kubectl; fi@@
cd /root/kubernets_tem@@
export KUBE_APISERVER="https://127.0.0.1:6443"@@
export BOOTSTRAP_TOKEN=$(head -c 16 /dev/urandom | od -An -t x | tr -d ' ')@@
cat > token.csv <<EOF@@
${BOOTSTRAP_TOKEN},kubelet-bootstrap,10001,"system:kubelet-bootstrap"@@
EOF@@

kubectl config set-cluster kubernetes --certificate-authority=k8s-root-ca.pem --embed-certs=true --server=${KUBE_APISERVER} --kubeconfig=bootstrap.kubeconfig@@
kubectl config set-credentials kubelet-bootstrap --token=${BOOTSTRAP_TOKEN} --kubeconfig=bootstrap.kubeconfig@@
kubectl config set-context default --cluster=kubernetes --user=kubelet-bootstrap --kubeconfig=bootstrap.kubeconfig@@
kubectl config use-context default --kubeconfig=bootstrap.kubeconfig@@

kubectl config set-cluster kubernetes --certificate-authority=k8s-root-ca.pem --embed-certs=true --server=${KUBE_APISERVER} --kubeconfig=kube-proxy.kubeconfig@@
kubectl config set-credentials kube-proxy --client-certificate=kube-proxy.pem --client-key=kube-proxy-key.pem --embed-certs=true --kubeconfig=kube-proxy.kubeconfig@@
kubectl config set-context default --cluster=kubernetes --user=kube-proxy --kubeconfig=kube-proxy.kubeconfig@@
kubectl config use-context default --kubeconfig=kube-proxy.kubeconfig@@

kubectl config set-cluster kubernetes --certificate-authority=k8s-root-ca.pem --embed-certs=true --server=https://127.0.0.1:6443@@
kubectl config set-credentials admin --client-certificate=admin.pem --embed-certs=true --client-key=admin-key.pem@@
kubectl config set-context kubernetes --cluster=kubernetes --user=admin@@
kubectl config use-context kubernetes@@

cp /root/.kube/config /root/kubernets_tem/config_tem@@
rm -rf /root/.kube/config
`
return cmd_initK8sEnv
}
func CopyK8sCmd(node *models.KubeHost)string {
	var cmd_copyK8sCmdService,cmd_copyK8sCmdNode,copyK8sCmd string
	cmd_copyK8sCmdService =`
if [ ! -e /usr/bin/kube-proxy ]; then cp /root/kubenetes/server/kube-proxy /usr/bin/;fi@@
if [ ! -e /usr/bin/kubectl ]; then cp /root/kubenetes/server/kubectl /usr/bin/;fi@@
if [ ! -e /usr/bin/kubelet ]; then cp /root/kubenetes/server/kubelet /usr/bin/;fi@@
if [ ! -e /usr/bin/hyperkube ]; then cp /root/kubenetes/server/hyperkube /usr/bin/;fi@@
if [ ! -e /usr/bin/kube-apiserver ]; then cp /root/kubenetes/server/kube-apiserver /usr/bin/;fi@@
if [ ! -e /usr/bin/kube-controller-manager ]; then cp /root/kubenetes/server/kube-controller-manager /usr/bin/;fi@@
if [ ! -e /usr/bin/kube-scheduler ]; then cp /root/kubenetes/server/kube-scheduler /usr/bin/;fi@@
if [ ! -e /lib/systemd/system/kubelet.service ]; then cp /root/kubenetes/cmd/kubelet.service /lib/systemd/system/;fi@@
if [ ! -e /lib/systemd/system/kube-proxy.service ]; then cp /root/kubenetes/cmd/kube-proxy.service /lib/systemd/system/;fi@@
if [ ! -e /lib/systemd/system/kube-apiserver.service ]; then cp /root/kubenetes/cmd/kube-apiserver.service /lib/systemd/system/;fi@@
if [ ! -e /lib/systemd/system/kube-controller-manager.service ]; then cp /root/kubenetes/cmd/kube-controller-manager.service /lib/systemd/system/;fi@@
if [ ! -e /lib/systemd/system/kube-scheduler.service ]; then cp /root/kubenetes/cmd/kube-scheduler.service /lib/systemd/system/;fi@@
chmod 755 /usr/bin/kubelet &&chmod 755 /usr/bin/kube-proxy &&chmod 755 /usr/bin/kubectl &&chmod 755 /usr/bin/kube-apiserver &&chmod 755 /usr/bin/kube-controller-manager &&chmod 755 /usr/bin/kube-scheduler &&chmod 755 /usr/bin/hyperkube
`
	cmd_copyK8sCmdNode=`
if [ ! -e /usr/bin/kube-proxy ]; then cp /root/kubenetes/server/kube-proxy /usr/bin/;fi@@
if [ ! -e /usr/bin/kubectl ]; then cp /root/kubenetes/server/kubectl /usr/bin/;fi@@
if [ ! -e /usr/bin/kubelet ]; then cp /root/kubenetes/server/kubelet /usr/bin/;fi@@
if [ ! -e /usr/bin/kubeadm ]; then cp /root/kubenetes/server/kubeadm /usr/bin/;fi`
	if(strings.Compare(node.Role,"master")==0){
		copyK8sCmd=cmd_copyK8sCmdService
	}else{
		copyK8sCmd=cmd_copyK8sCmdNode
	}
	return copyK8sCmd
}

//幂等性		主节点执行一次
func CreateCert(list []*models.KubeHost)string{
	//chmod +x /root/install/cfssl/cfssl_linux-amd64
	var cmd_createK8sCert string
	//cmd = "echo '192.168.60.41 master1'>>/etc/hosts;echo '192.168.60.41 master1'>>/etc/hosts;echo '192.168.60.41 master1'>>/etc/hosts"
	//kubernetes-csr.json
	cmd_createK8sCert="if [ ! -e /root/kubernets_tem/k8s-root-ca.pem ]; then " +
		"mkdir -p /root/kubernets_tem " +
		"&&cp /root/kubenetes/ssl/* /root/kubernets_tem/ " +
			"&&cp -r /root/kubenetes/cmd /root/kubernets_tem/ " +
			"&&cp -r /root/kubenetes/config /root/kubernets_tem/ " +
			//"&&sed -i 's/@MASTER_IPS/@MasterIP/g' /root/kubernets_tem/kubernetes-csr.json " +
			//"&&sed -i 's/@APISERVER_COUNT/@COUNT/g' /root/kubernets_tem/config/apiserver "+
			"&&cd /root/kubernets_tem " +
			"&&cfssl gencert -initca k8s-root-ca-csr.json | cfssljson -bare k8s-root-ca " +
			"&&cfssl gencert -ca=k8s-root-ca.pem -ca-key=k8s-root-ca-key.pem -config=k8s-gencert.json -profile=kubernetes -hostname='127.0.0.1,10.254.0.1,@HOST_LISTlocalhost,kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster,kubernetes.default.svc.cluster.local' kubernetes-csr.json | cfssljson -bare kubernetes " +
			"&&cfssl gencert -ca=k8s-root-ca.pem -ca-key=k8s-root-ca-key.pem -config=k8s-gencert.json -profile=kubernetes admin-csr.json | cfssljson -bare admin " +
			"&&cfssl gencert -ca=k8s-root-ca.pem -ca-key=k8s-root-ca-key.pem -config=k8s-gencert.json -profile=kubernetes kube-proxy-csr.json | cfssljson -bare kube-proxy; fi"
		//rm -rf /root/kubernets_tem/*.pem
		//  &&rm -rf /root/kubernets_tem/*.csr
		//	&&rm -rf /root/kubernets_tem/*.json
		//&&rm -rf *.csr &&rm -rf kubernetes-csr.json
		//修改证书中的master  ip 地址
		var buf,hostList bytes.Buffer
		var len int
		len=0
		for _,node:=range list{
			if(strings.Compare(node.Role,"master")==0){
				//"192.168.60.39",
				buf.WriteString("\""+node.Ip+"\",")
				hostList.WriteString(node.Ip+",")
				len=len+1
			}
		}
		lenStr:=strconv.Itoa(len)
		//创建k8s证书  /etc/kubernetes/ssl
		createK8sCert := strings.Replace(cmd_createK8sCert, "@MasterIP", buf.String(), -1)
		createK8sCert = strings.Replace(createK8sCert, "@COUNT", lenStr, -1)
		createK8sCert = strings.Replace(createK8sCert, "@HOST_LIST", hostList.String(), -1)

		//fmt.Println("创建k8s证书======开始========")
		//	//if session, err := client.NewSession(); err == nil {
		//	//	defer session.Close()
		//	//	session.Run(createK8sCert)
		//	//}
		//	//fmt.Println("创建k8s证书======结束========")
		return createK8sCert
	}
	//幂等性
	func CopyK8sCert(master *models.KubeHost)string{
		var cmd_copyK8sCert string
		//需要校验是否有问题
		cmd_copyK8sCert="if [ ! -d @K8S_SSL_PATH ]; then mkdir -p @K8S_SSL_PATH &&cp /root/kubernets_tem/*.pem @K8S_SSL_PATH/ &&chown -R kube:kube @K8S_SSL_PATH &&chown -R kube:kube @K8S_SSL_PATH/* &&mkdir -p /var/run/kubernetes &&chown -R kube:kube /var/run/kubernetes; fi"
		//分发所有的证书
		//copyK8sCert := strings.Replace(cmd_copyK8sCert, "@MASTER_USER", master.User, -1)
		//copyK8sCert = strings.Replace(copyK8sCert, "@MASTER_IP", master.Ip, -1)
		cmd_copyK8sCert = strings.Replace(cmd_copyK8sCert, "@K8S_SSL_PATH", K8S_SSL_PATH, -1)
		return cmd_copyK8sCert
	}


	func CopyK8sEnv(node *models.KubeHost,master *models.KubeHost,list []*models.KubeHost)string{
		var cmd_copyK8sBase,cmd_copyK8sCmdService string
		cmd_copyK8sBase=`
	if [ ! -e /etc/kubernetes/token.csv ]; then cp /root/kubernets_tem/token.csv /etc/kubernetes/;fi@@
	if [ ! -e /etc/kubernetes/bootstrap.kubeconfig ]; then cp /root/kubernets_tem/bootstrap.kubeconfig /etc/kubernetes/ &&sed -i s/127.0.0.1/@IP/g /etc/kubernetes/bootstrap.kubeconfig;fi@@
	if [ ! -e /etc/kubernetes/kube-proxy.kubeconfig ]; then cp /root/kubernets_tem/kube-proxy.kubeconfig /etc/kubernetes/ &&sed -i s/127.0.0.1/@IP/g /etc/kubernetes/kube-proxy.kubeconfig;fi@@
	if [ ! -e /etc/kubernetes/audit-policy.yaml ]; then cp /root/kubernets_tem/config/audit-policy.yaml /etc/kubernetes/;fi@@
	if [ ! -e /etc/kubernetes/config ]; then cp /root/kubernets_tem/config/config /etc/kubernetes/ &&sed -i s/127.0.0.1/@IP/g /etc/kubernetes/config;fi
	`
		cmd_copyK8sCmdService=`
	if [ ! -e /root/.kube/config ]; then mkdir -p /root/.kube &&cp /root/kubernets_tem/config_tem /root/.kube/config &&sed -i s/127.0.0.1/@IP/g /root/.kube/config;fi@@
	if [ ! -e /etc/kubernetes/apiserver ]; then cp /root/kubernets_tem/config/apiserver /etc/kubernetes/ &&sed -i s/@HOST_IP/@IP/g /etc/kubernetes/apiserver &&sed -i s/@ETCD_CLUSTER_URL/@CLUSTER_LIST/g /etc/kubernetes/apiserver;fi@@
	if [ ! -e /etc/kubernetes/scheduler ]; then cp /root/kubernets_tem/config/scheduler /etc/kubernetes/;fi@@
	if [ ! -e /etc/kubernetes/controller-manager ]; then cp /root/kubernets_tem/config/controller-manager /etc/kubernetes/;fi@@
		`
		///lib/systemd/system/kube-apiserver.service      scp -r @MASTER_USER@@MASTER_IP:/root/kubenetes/kube-apiserver.service /lib/systemd/system/
		///lib/systemd/system/kube-controller-manager.service   scp -r @MASTER_USER@@MASTER_IP:/root/kubenetes/kube-controller-manager.service /lib/systemd/system/
		///lib/systemd/system/kube-scheduler.service     scp -r @MASTER_USER@@MASTER_IP:/root/kubenetes/kube-scheduler.service /lib/systemd/system/
		var buf bytes.Buffer
		for _,value:=range list{
			//,etcd3=https://`cat /etc/hosts|grep master3 |awk '{print $1}'`:2380
			if(strings.Compare(value.Role,"master")==0){
				if(buf.Len()>0){
					buf.WriteString(",")
				}
				buf.WriteString("https:\\\\/\\\\/")
				buf.WriteString(value.Ip)
				buf.WriteString(":2379")
			}
		}
		//fmt.Println(buf.String())

		if(strings.Compare(node.Role,"master")==0){
			cmd_copyK8sCmdService = cmd_copyK8sBase+"@@"+cmd_copyK8sCmdService
			cmd_copyK8sCmdService = strings.Replace(cmd_copyK8sCmdService, "@IP", node.Ip, -1)
			cmd_copyK8sCmdService = strings.Replace(cmd_copyK8sCmdService, "@MASTER_IP", master.Ip, -1)
			cmd_copyK8sCmdService = strings.Replace(cmd_copyK8sCmdService, "@MASTER_USER", master.User, -1)
			cmd_copyK8sCmdService = strings.Replace(cmd_copyK8sCmdService, "@CLUSTER_LIST", buf.String(), -1)
			return cmd_copyK8sCmdService
		}else{
			cmd_copyK8sBase = strings.Replace(cmd_copyK8sBase, "@IP", node.Ip, -1)
			cmd_copyK8sBase = strings.Replace(cmd_copyK8sBase, "@MASTER_IP", master.Ip, -1)
			cmd_copyK8sBase = strings.Replace(cmd_copyK8sBase, "@MASTER_USER", master.User, -1)
			return cmd_copyK8sBase
		}
	}


func RestartNodeService(node *models.KubeHost)string{
	var restartNodeService string
	//需要校验是否有问题
	restartNodeService="systemctl restart kubelet &&systemctl restart kube-proxy"
	return restartNodeService
}
func RestartMasterService(node *models.KubeHost)string{
	var restartMasterService string
	restartMasterService="systemctl restart kube-apiserver"+
		" &&systemctl restart kube-controller-manager"+
		" &&systemctl restart kube-scheduler"
	return restartMasterService
}
