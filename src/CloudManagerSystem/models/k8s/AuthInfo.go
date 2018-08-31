package k8s

import (
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/rest"

	"errors"
	"github.com/astaxie/beego/orm"
	"CloudManagerSystem/models"
)

type Authenticator interface {
	// GetAuthInfo returns filled AuthInfo structure that can be used for K8S api client creation.
	GetAuthInfo() (api.AuthInfo, error)
}

// LoginSpec is extracted from request coming from Dashboard frontend during login request. It contains all the
// information required to authenticate user.
type LoginSpec struct {
	// Username is the username for basic authentication to the kubernetes cluster.
	Username string `json:"username"`
	// Password is the password for basic authentication to the kubernetes cluster.
	Password string `json:"password"`
	// Token is the bearer token for authentication to the kubernetes cluster.
	//Token string `json:"token"`
	// KubeConfig is the content of users' kubeconfig file. It will be parsed and auth data will be extracted.
	// Kubeconfig can not contain any paths. All data has to be provided within the file.
	//KubeConfig string `json:"kubeConfig"`
}

// Implements Authenticator interface
type basicAuthenticator struct {
	username string
	password string
}

// GetAuthInfo implements Authenticator interface. See Authenticator for more information.
func (self *basicAuthenticator) GetAuthInfo() (api.AuthInfo, error) {
	return api.AuthInfo{
		Username: self.username,
		Password: self.password,
	}, nil
}

// NewBasicAuthenticator returns Authenticator based on LoginSpec.
func NewBasicAuthenticator(spec *LoginSpec) Authenticator {
	return &basicAuthenticator{
		username: spec.Username,
		password: spec.Password,
	}
}

// Returns authenticator based on provided LoginSpec.
func getAuthenticator(spec *LoginSpec) Authenticator {
	return NewBasicAuthenticator(spec)
}

func buildConfigFromFlags(ClusterId string) (*rest.Config, error) {

	//clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
	//	&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
	//	&clientcmd.ConfigOverrides{ClusterInfo: api.Cluster{Server: apiserverHost}}).ClientConfig()

	var config *rest.Config

	adminbuf := models.GetCertDataByCluster(ClusterId, "admin")
	adminkeybuf := models.GetCertDataByCluster(ClusterId, "admin-key")
	k8srootcabuf := models.GetCertDataByCluster(ClusterId, "k8s-root-ca")
	if adminbuf == "" || adminkeybuf == "" || k8srootcabuf == "" {
		return nil, errors.New("cert data error")
	}

	tlsClientConfig := rest.TLSClientConfig{CertData: []byte(adminbuf), KeyData: []byte(adminkeybuf), CAData: []byte(k8srootcabuf)}

	query := orm.NewOrm().QueryTable(models.KubeHostTBName())
	data := make([]*models.KubeHost, 0)
	query.Filter("cluster_id", ClusterId).Filter("role", 0).All(&data) //check master

	for _, host := range data {

		config = &rest.Config{
			//Host:            "https://" + net.JoinHostPort(host, port),
			Host: "https://" + host.Ip + ":6443",
			//Host:"https://192.168.1.153:6443",
			TLSClientConfig: tlsClientConfig,
		}
		_, err := kubernetes.NewForConfig(config)

		if err == nil {
			//fmt.Println()
			return config, nil
		}
	}
	return nil, errors.New("error config build")
}

// Based on auth info and rest config creates client cmd config.
func buildCmdConfig(authInfo *api.AuthInfo, cfg *rest.Config) clientcmd.ClientConfig {
	cmdCfg := api.NewConfig()
	cmdCfg.Clusters[models.DefaultCmdConfigName] = &api.Cluster{
		Server:                   cfg.Host,
		CertificateAuthority:     cfg.TLSClientConfig.CAFile,
		CertificateAuthorityData: cfg.TLSClientConfig.CAData,
		InsecureSkipTLSVerify:    cfg.TLSClientConfig.Insecure,
	}
	cmdCfg.AuthInfos[models.DefaultCmdConfigName] = authInfo
	cmdCfg.Contexts[models.DefaultCmdConfigName] = &api.Context{
		Cluster:  models.DefaultCmdConfigName,
		AuthInfo: models.DefaultCmdConfigName,
	}
	cmdCfg.CurrentContext = models.DefaultCmdConfigName

	return clientcmd.NewDefaultClientConfig(
		*cmdCfg,
		&clientcmd.ConfigOverrides{},
	)
}

func HasAccess(authInfo api.AuthInfo, ClusterId string) error {
	cfg, err := buildConfigFromFlags(ClusterId)
	if err != nil {
		return err
	}

	clientConfig := buildCmdConfig(&authInfo, cfg)
	cfg, err = clientConfig.ClientConfig()
	if err != nil {
		return err
	}

	_, k8serr := kubernetes.NewForConfig(cfg)
	if k8serr != nil {
		return k8serr
	}

	return nil
}

// Login implements auth manager. See AuthManager interface for more information.
func K8sLogin(username string, passwd string, ClusterId string) error {
	loginSpec := new(LoginSpec)
	loginSpec.Username = username
	loginSpec.Password = passwd

	authenticator := getAuthenticator(loginSpec)
	authInfo, err := authenticator.GetAuthInfo()
	if err != nil {
		return err
	}

	err = HasAccess(authInfo, ClusterId)
	if err != nil {
		return err
	}

	return nil
}


