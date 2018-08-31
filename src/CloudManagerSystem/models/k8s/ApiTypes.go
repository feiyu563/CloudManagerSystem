package k8s




// ClientType represents type of client that is used to perform generic operations on resources.
// Different resources belong to different client, i.e. Deployments belongs to extension client
// and StatefulSets to apps client.
type ClientType string

// List of client types supported by the UI.
const (
	ClientTypeDefault           = "restclient"
	ClientTypeExtensionClient   = "extensionclient"
	ClientTypeAppsClient        = "appsclient"
	ClientTypeBatchClient       = "batchclient"
	ClientTypeBetaBatchClient   = "betabatchclient"
	ClientTypeAutoscalingClient = "autoscalingclient"
	ClientTypeStorageClient     = "storageclient"
)

var KindToAPIMapping = map[string]struct {
	// Kubernetes resource name.
	Resource string
	// Client type used by given resource, i.e. deployments are using extension client.
	ClientType ClientType
	// Is this object global scoped (not below a namespace).
	Namespaced bool
}{
	ResourceKindConfigMap:               {"configmaps", ClientTypeDefault, true},
	ResourceKindDaemonSet:               {"daemonsets", ClientTypeExtensionClient, true},
	ResourceKindDeployment:              {"deployments", ClientTypeExtensionClient, true},
	ResourceKindEvent:                   {"events", ClientTypeDefault, true},
	ResourceKindHorizontalPodAutoscaler: {"horizontalpodautoscalers", ClientTypeAutoscalingClient, true},
	ResourceKindIngress:                 {"ingresses", ClientTypeExtensionClient, true},
	ResourceKindJob:                     {"jobs", ClientTypeBatchClient, true},
	ResourceKindCronJob:                 {"cronjobs", ClientTypeBetaBatchClient, true},
	ResourceKindLimitRange:              {"limitrange", ClientTypeDefault, true},
	ResourceKindNamespace:               {"namespaces", ClientTypeDefault, false},
	ResourceKindNode:                    {"nodes", ClientTypeDefault, false},
	ResourceKindPersistentVolumeClaim:   {"persistentvolumeclaims", ClientTypeDefault, true},
	ResourceKindPersistentVolume:        {"persistentvolumes", ClientTypeDefault, false},
	ResourceKindPod:                     {"pods", ClientTypeDefault, true},
	ResourceKindReplicaSet:              {"replicasets", ClientTypeExtensionClient, true},
	ResourceKindReplicationController:   {"replicationcontrollers", ClientTypeDefault, true},
	ResourceKindResourceQuota:           {"resourcequotas", ClientTypeDefault, true},
	ResourceKindSecret:                  {"secrets", ClientTypeDefault, true},
	ResourceKindService:                 {"services", ClientTypeDefault, true},
	ResourceKindStatefulSet:             {"statefulsets", ClientTypeAppsClient, true},
	ResourceKindStorageClass:            {"storageclasses", ClientTypeStorageClient, false},
	ResourceKindEndpoint:                {"endpoints", ClientTypeDefault, true},
}

