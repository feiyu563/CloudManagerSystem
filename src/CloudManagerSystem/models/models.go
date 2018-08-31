package models

import (
	"github.com/astaxie/beego/orm"
)
//初始化
func init() {
	orm.RegisterModel(new(BackendUser), new(Resource), new(Role), new(RoleResourceRel), new(RoleBackendUserRel),new(KubeHost),new(KubeNameSpace),new(KubeAuthUserNameSpace),new(ClusterRole),new(ClusterResource),new(KubeServiceAccounts),new(KubeBind),new(KubeUserGroup),new(KubeUser2Group),new(KubeCluster),new (KubeEnvUserCluster),new (KubeEnvUserNamespace),new(KubePublishService),new(KubePublishServicePath))
}
