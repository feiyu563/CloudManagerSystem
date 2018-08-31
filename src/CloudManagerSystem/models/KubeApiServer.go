package models

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
	"github.com/astaxie/beego"
	"fmt"
	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/orm"
	"sync"
	"k8s.io/client-go/tools/clientcmd/api"
	"path/filepath"
	"io/ioutil"
	"errors"
	"strings"
)

var ApiServerHandleMap map[string]*kubernetes.Clientset

func init() {
	ApiServerHandleMap = make(map[string]*kubernetes.Clientset, 1024)
}

func LoginClusterApiserverInit(UserId string) {
	/*del all*/
	for key, _ := range ApiServerHandleMap {
		delete(ApiServerHandleMap, key)
	}
	//delete end
	user, err := BackendUserOne(UserId)
	if err != nil || user == nil {
		beego.Error("No user auth!")
		return
	}
	var mutex sync.Mutex

	if user.IsSuper == true {
		query := orm.NewOrm().QueryTable(KubeClusterTBName())
		data := make([]*KubeCluster, 0)
		query.All(&data)
		mutex.Lock()
		for _, m := range data {
			//fmt.Println(m)
			clienthandle, err := GetApiServerHandle(m.Id, false)
			if err != nil {
				continue
			}
			ApiServerHandleMap[m.Id] = clienthandle
		}
		mutex.Unlock()
	} else {
		datausercluster := make([]*KubeUserAuthCluster, 0)
		queryusercluster := orm.NewOrm().QueryTable(KubeUserAuthClusterTBName())
		queryusercluster.Filter("user_id", UserId).All(&datausercluster)
		mutex.Lock()
		for _, m := range datausercluster {
			//fmt.Println(m)
			clienthandle, err := GetApiServerHandle(m.ClusterId, false)
			if err != nil {
				continue
			}
			ApiServerHandleMap[m.ClusterId] = clienthandle
		}
		mutex.Unlock()

	}
	fmt.Println("########################", ApiServerHandleMap)
}

func GetApiServerHandle(ClusterId string, cacheorlocal bool) (*kubernetes.Clientset, error) {
	var configpath string
	var config *rest.Config
	var err error
	//ClusterId = "test"
	if cacheorlocal {
		configpath = filepath.Join("conf", "kubeconfig", ClusterId, ClusterId)
	} else {
		configpath = filepath.Join("conf", "kubeconfig", ClusterId, "config")
	}
	//fmt.Println(configpath)

	buf, err := ioutil.ReadFile(configpath)
	if err != nil {
		return nil, errors.New("config read error")
	}
	var cflag bool
	content := string(buf)
	//fmt.Println(content)
	query := orm.NewOrm().QueryTable(KubeHostTBName())
	data := make([]*KubeHost, 0)
	query.Filter("cluster_id", ClusterId).Filter("role", 0).All(&data) //check master
	cflag = false
	for _, host := range data {
		if strings.Contains(content, host.Ip) {
			cflag = true
		}
	}
	if !cflag {
		return nil, errors.New("config Ip error")
	}

	if configpath == "" {
		beego.Info("Using in cluster config")
		config, err = rest.InClusterConfig()
		// in cluster access
	} else {
		beego.Info("Using out of cluster config")
		config, err = clientcmd.BuildConfigFromFlags("", configpath)
	}

	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfigOrDie(config), nil
}

//func GetApiServerHandle(ClusterId string, cacheorlocal bool) (*kubernetes.Clientset, error) {
//	var config *rest.Config
//
//	adminbuf := GetCertDataByCluster(ClusterId, "admin")
//	adminkeybuf := GetCertDataByCluster(ClusterId, "admin-key")
//	k8srootcabuf := GetCertDataByCluster(ClusterId, "k8s-root-ca")
//	if adminbuf == "" || adminkeybuf == "" || k8srootcabuf == ""{
//		return nil , errors.New("cert data error")
//	}
//
//	tlsClientConfig := rest.TLSClientConfig{CertData: []byte(adminbuf), KeyData: []byte(adminkeybuf), CAData: []byte(k8srootcabuf)}
//
//	query := orm.NewOrm().QueryTable(KubeHostTBName())
//	data := make([]*KubeHost, 0)
//	query.Filter("cluster_id", ClusterId).Filter("role", 0).All(&data) //check master
//
//	for _, host := range data {
//
//		config = &rest.Config{
//			//Host:            "https://" + net.JoinHostPort(host, port),
//			Host: "https://" + host.Ip + ":6443",
//
//			//Host:"https://192.168.1.153:6443",
//			TLSClientConfig: tlsClientConfig,
//		}
//		clientset, err := kubernetes.NewForConfig(config)
//
//		if err == nil {
//			//fmt.Println()
//			return clientset, nil
//		}
//	}
//
//	return nil, errors.New("Ip machine is error")
//}

// Dashboard UI default values for client configs.
const (
	// High enough QPS to fit all expected use cases. QPS=0 is not set here, because
	// client code is overriding it.
	DefaultQPS = 1e6
	// High enough Burst to fit all expected use cases. Burst=0 is not set here, because
	// client code is overriding it.
	DefaultBurst = 1e6
	// Use kubernetes protobuf as content type by default
	DefaultContentType = "application/vnd.kubernetes.protobuf"
	// Default cluster/context/auth name to be set in clientcmd config
	DefaultCmdConfigName = "kubernetes"
	// Header name that contains token used for authorization. See TokenManager for more information.
	JWETokenHeader = "jweToken"
	// Default http header for user-agent
	DefaultUserAgent = "dashboard"
)

// Based on auth info and rest config creates client cmd config.
func buildCmdConfig(authInfo *api.AuthInfo, cfg *rest.Config) clientcmd.ClientConfig {
	cmdCfg := api.NewConfig()
	cmdCfg.Clusters[DefaultCmdConfigName] = &api.Cluster{
		Server:                   cfg.Host,
		CertificateAuthority:     cfg.TLSClientConfig.CAFile,
		CertificateAuthorityData: cfg.TLSClientConfig.CAData,
		InsecureSkipTLSVerify:    cfg.TLSClientConfig.Insecure,
	}
	cmdCfg.AuthInfos[DefaultCmdConfigName] = authInfo
	cmdCfg.Contexts[DefaultCmdConfigName] = &api.Context{
		Cluster:  DefaultCmdConfigName,
		AuthInfo: DefaultCmdConfigName,
	}
	cmdCfg.CurrentContext = DefaultCmdConfigName
	//一个可选的AutH信息回退阅读器
	return clientcmd.NewDefaultClientConfig(
		*cmdCfg,
		&clientcmd.ConfigOverrides{},
	)
}
