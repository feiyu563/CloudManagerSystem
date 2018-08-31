package controllers
import (
	"fmt"
	"CloudManagerSystem/models"
	"CloudManagerSystem/enums"
	"k8s.io/api/storage/v1"
	"bytes"
	"strings"
)
type InstallClusterController struct {
	BaseController
}

/**
创建基础环境
1.选择集群ip
2.查询集群ip下所有的节点
*/
func (c *InstallClusterController) CreateBaseCluster() {
	strs := c.GetString("id")
	list,_:=models.GetAllKubeHostWithNOCluster(strs)
	CreateCluster(list)
	c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("创建成功"), 0)
}

type StorageClass struct {
	ClusterId string `json:"clusterId"`
	PoolName   string `json:"poolName"`
	PoolSize  string `json:"poolSize"`
	User  string `json:"user"`
	K8sNameSpace  string `json:"k8sNameSpace"`
	K8sSecretName  string `json:"k8sSecretName"`
	K8sStorageClassName  string `json:"k8sStorageClassName"`
}

func (c *InstallClusterController) CreateStorageClass() {
	u:=c.GetSessionUser()
	if (  len(u.Id) == 0) {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("id 不能为空"), 0)
	}
	data,_ := models.EnvUserCluster(u.Id)
	m := StorageClass{}
	c.ParseForm(&m)
	m.ClusterId=data.ClusterId
	var master *models.KubeHost
	list,_:=models.GetAllKubeHostWithNOCluster(m.ClusterId)
	var buf bytes.Buffer
	//连接安装节点
	for _,node := range list{
		if(node.IsInstallNode){
			master=node
		}
		//如果是主节点
		if(strings.Compare(node.Role,"master")==0){
			if(buf.Len()>0){
				buf.WriteString(",")
			}
			//172.16.20.2:6789
			buf.WriteString(node.Ip)
			buf.WriteString(":6789")
		}
	}
	client1:=GetClient(master)
	if(client1==nil){
		fmt.Println("当前机器%s无法连接",master.Ip)
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("当前机器无法连接"), 0)
		return
	}
	defer client1.Close()
	fmt.Println("==========%s========",m.PoolName)
	//创建pool池
	createPool:="ceph osd pool create "+m.PoolName+" 512"
	ExeCmd(client1,createPool)
	//ceph创建用户和权限秘钥
	createUser:="ceph auth get-or-create client."+m.User+" mon 'allow r' osd 'allow rwx pool="+m.PoolName+"'"
	ExeCmd(client1,createUser)
	//k8s创建admin秘钥
	createAdminK8sSecret:="kubectl create secret generic ceph-secret-admin --from-literal=key=`ceph auth get-key client.admin` --type=kubernetes.io/rbd -n kube-system"
	ExeCmd(client1,createAdminK8sSecret)
	//k8s创建用户秘钥
	createUserK8sSecret:="kubectl create secret generic "+m.K8sSecretName+" --from-literal=key=`ceph auth get-key client."+m.User+"` --type=kubernetes.io/rbd -n "+m.K8sNameSpace
	ExeCmd(client1,createUserK8sSecret)
	//k8s创建StorageClass
	clienthandle, err := models.GetApiServerHandle(m.ClusterId, false)
	fmt.Println(err)
	sc:=&v1.StorageClass{}
	sc.APIVersion="storage.k8s.io/v1"
	sc.Kind="StorageClass"
	sc.Name=m.K8sStorageClassName
	sc.Namespace=m.K8sNameSpace
	sc.Provisioner="kubernetes.io/rbd"
	sc.Parameters=make(map[string]string)
	sc.Parameters [ "monitors" ] = buf.String()
	sc.Parameters [ "adminId" ] ="admin"
	sc.Parameters [ "adminSecretName" ] ="ceph-secret-admin"
	sc.Parameters [ "adminSecretNamespace" ] ="kube-system"
	sc.Parameters [ "pool" ] =m.PoolName
	sc.Parameters [ "userId" ] =m.User
	sc.Parameters [ "userId" ] =m.K8sSecretName
	clienthandle.StorageV1().StorageClasses().Create(sc)
	c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("创建成功"), 0)
}


