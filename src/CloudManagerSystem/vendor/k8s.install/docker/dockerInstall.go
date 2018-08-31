package docker

import (
	"strings"
	"CloudManagerSystem/models"
)

//幂等性
//func CopyDocker(client *ssh.Client,master *host.KubeHost){
//	//拷贝安装文件
//	var copyPacker,copyImages,installDocker,addSysBin string
//
//	copyPacker="if [ ! -d /root/docker ]; then scp -r @MASTER_USER@@MASTER_IP:/root/docker /root/; fi"
//	copyPacker = strings.Replace(copyPacker, "@MASTER_USER", master.User, -1)
//	copyPacker = strings.Replace(copyPacker, "@MASTER_IP", master.Ip, -1)
//	sshUtil.ExeCmd(client,copyPacker)
//
//	//拷贝安装镜像
//	copyImages ="if [ ! -d /root/dockerImages ]; then scp -r @MASTER_USER@@MASTER_IP:/root/dockerImages /root/; fi"
//	copyImages = strings.Replace(copyImages, "@MASTER_USER", master.User, -1)
//	copyImages = strings.Replace(copyImages, "@MASTER_IP", master.Ip, -1)
//	sshUtil.ExeCmd(client,copyImages)
//
//	//安装doker
//	installDocker ="if [ ! -e /usr/bin/dockerd ]; then cd /root/docker &&yum install -y yum-utils device-mapper-persistent-data lvm2 &&yum install -y docker-ce-*.ce-1.el7.centos.x86_64.rpm;fi"
//	sshUtil.ExeCmd(client,installDocker)
//	//添加启动
//	addSysBin ="if [ ! -e /usr/lib/systemd/system/docker.service ]; then cp /root/docker/docker.service /usr/lib/systemd/system/;fi"
//	//&&systemctl daemon-reload && systemctl restart docker
//	sshUtil.ExeCmd(client,addSysBin)
//}


func CopyDocker(master *models.KubeHost)string{
	//拷贝安装文件
	var copyPacker,copyImages,installDocker,addSysBin string
	copyPacker="if [ ! -d /root/docker ]; then scp -r @MASTER_USER@@MASTER_IP:/root/docker /root/; fi"
	copyPacker = strings.Replace(copyPacker, "@MASTER_USER", master.User, -1)
	copyPacker = strings.Replace(copyPacker, "@MASTER_IP", master.Ip, -1)

	//拷贝安装镜像
	copyImages ="if [ ! -d /root/dockerImages ]; then scp -r @MASTER_USER@@MASTER_IP:/root/dockerImages /root/; fi"
	copyImages = strings.Replace(copyImages, "@MASTER_USER", master.User, -1)
	copyImages = strings.Replace(copyImages, "@MASTER_IP", master.Ip, -1)

	//安装doker
	installDocker ="if [ ! -e /usr/bin/dockerd ]; then cd /root/docker &&yum install -y yum-utils device-mapper-persistent-data lvm2 &&yum install -y docker-ce-*.ce-1.el7.centos.x86_64.rpm;fi"

	//添加启动
	addSysBin ="rm -rf /usr/lib/systemd/system/docker.service &&cp /root/docker/docker.service /usr/lib/systemd/system/"
	//&&systemctl daemon-reload && systemctl restart docker
	return copyPacker+" &&"+copyImages+" &&"+installDocker+" &&"+addSysBin
}

