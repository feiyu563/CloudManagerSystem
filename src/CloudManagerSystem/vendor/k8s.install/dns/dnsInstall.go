package dns

import (
	"strings"
	"CloudManagerSystem/models"
)

func CopyFile(master *models.KubeHost)string{
	//  拷贝文件/root/cmd/
	var copyFile string
	copyFile="if [ ! -d /root/dns ]; then scp -r @MASTER_USER@@MASTER_IP:/root/dns /root; fi"
	copyFile = strings.Replace(copyFile, "@MASTER_IP", master.Ip, -1)
	copyFile = strings.Replace(copyFile, "@MASTER_USER", master.User, -1)
	return copyFile
	//sshUtil.ExeCmd(client,copyFile)
}

func K8sCreate()string{
	var copyFile string
	copyFile="kubectl create -f kube-dns.yaml"
	return copyFile
}

func DockerReloadImages()string{
	var reloadImages string
	reloadImages="cd /root/dns @@" +
		"for i in `ls |grep tar`;do docker load -i $i;done "
		return reloadImages
}