package network

import (
	"strings"
	"CloudManagerSystem/models"
)

//拷贝网络插件
func CopyFile(master *models.KubeHost)string{
	//  拷贝文件/root/cmd/
	var copyFile string
	copyFile="if [ ! -d /root/calico ]; then scp -r @MASTER_USER@@MASTER_IP:/root/calico /root; fi"
	copyFile = strings.Replace(copyFile, "@MASTER_IP", master.Ip, -1)
	copyFile = strings.Replace(copyFile, "@MASTER_USER", master.User, -1)
	return copyFile
}

func InstallNetWork(node *models.KubeHost,master *models.KubeHost){
	//  拷贝文件/root/cmd/
	//var copyFile string

	//sed -i 's@.*etcd_endpoints:.*@\ \ etcd_endpoints:\ \"192.168.60.29:2379\"@gi' calico.yaml
	//sed 's/__ETCD_ENDPOINTS__/192.168.60.29:2379/g' calico.yaml
	//sed 's/__ETCD_KEY_FILE__/'/etc/etcd/ssl/etcd-key.pem'/g' calico.yaml
	//sed 's/__ETCD_ENDPOINTS__/192.168.60.29:2379/g' calico.yaml
	//sed 's/__ETCD_ENDPOINTS__/192.168.60.29:2379/g' calico.yaml
	//export ETCD_CERT=`cat /etc/etcd/ssl/etcd.pem | base64 | tr -d '\n'`
	//export ETCD_KEY=`cat /etc/etcd/ssl/etcd-key.pem | base64 | tr -d '\n'`
	//export ETCD_CA=`cat /etc/etcd/ssl/etcd-root-ca.pem | base64 | tr -d '\n'`
	//sed -i "s@.*etcd-cert:.*@\ \ etcd-cert:\ ${ETCD_CERT}@gi" calico.yaml
	//sed -i "s@.*etcd-key:.*@\ \ etcd-key:\ ${ETCD_KEY}@gi" calico.yaml
	//sed -i "s@.*etcd-ca:.*@\ \ etcd-ca:\ ${ETCD_CA}@gi" calico.yaml
	//sed -i 's@.*etcd_ca:.*@\ \ etcd_ca:\ "/calico-secrets/etcd-ca"@gi' calico.yaml
	//sed -i 's@.*etcd_cert:.*@\ \ etcd_cert:\ "/calico-secrets/etcd-cert"@gi' calico.yaml
	//sed -i 's@.*etcd_key:.*@\ \ etcd_key:\ "/calico-secrets/etcd-key"@gi' calico.yaml
	//sed -i s/192.168.0.0/172.16.0.0/g calico.yaml
	//kubectl create -f rbac.yaml
	//kubectl create -f calico.yaml
	//sed -i '/--cluster-dns=10.254.0.2/a\              --network-plugin=cni \\' /etc/kubernetes/kubelet
	//systemctl restart kubelet
	//sshUtil.ExeShell(client,cmd_initK8sEnv,"@@")
}
func DockerReloadImages(node *models.KubeHost)string{
	//  拷贝文件/root/cmd/
	var copyFile string
	copyFile="docker load -i /root/calico/calico.tar"
	return copyFile
	//if(strings.Compare(node.Role,"master")==0){
	//	sshUtil.ExeCmd(client,copyFile)
	//}else{
	//	sshUtil.ExeCmd(client,copyFile)
	//}
}


