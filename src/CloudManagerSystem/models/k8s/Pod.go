package k8s

import (
	"CloudManagerSystem/models"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"

	"fmt"
)

var EmptyPodList = &PodList{
	Pods: make([]Pod, 0),
	//Errors: make([]error, 0),
	ListMeta: ListMeta{
		TotalItems: 0,
	},
	//ListMeta: 0,
}

const (
	ResourceKindConfigMap               = "configmap"
	ResourceKindDaemonSet               = "daemonset"
	ResourceKindDeployment              = "deployment"
	ResourceKindEvent                   = "event"
	ResourceKindHorizontalPodAutoscaler = "horizontalpodautoscaler"
	ResourceKindIngress                 = "ingress"
	ResourceKindJob                     = "job"
	ResourceKindCronJob                 = "cronjob"
	ResourceKindLimitRange              = "limitrange"
	ResourceKindNamespace               = "namespace"
	ResourceKindNode                    = "node"
	ResourceKindPersistentVolumeClaim   = "persistentvolumeclaim"
	ResourceKindPersistentVolume        = "persistentvolume"
	ResourceKindPod                     = "pod"
	ResourceKindReplicaSet              = "replicaset"
	ResourceKindReplicationController   = "replicationcontroller"
	ResourceKindResourceQuota           = "resourcequota"
	ResourceKindSecret                  = "secret"
	ResourceKindService                 = "service"
	ResourceKindStatefulSet             = "statefulset"
	ResourceKindStorageClass            = "storageclass"
	ResourceKindRbacRole                = "role"
	ResourceKindRbacClusterRole         = "clusterrole"
	ResourceKindRbacRoleBinding         = "rolebinding"
	ResourceKindRbacClusterRoleBinding  = "clusterrolebinding"
	ResourceKindEndpoint                = "endpoint"
)

// PodList contains a list of Pods in the cluster.
type PodList struct {
	ListMeta          ListMeta       `json:"listMeta"`
	//CumulativeMetrics []Metric `json:"cumulativeMetrics"`

	// Basic information about resources status on the list.
	Status ResourceStatus `json:"status"`

	// Unordered list of Pods.
	Pods []Pod `json:"pods"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

// Pod is a presentation layer view of Kubernetes Pod resource. This means it is Pod plus additional augmented data
// we can get from other sources (like services that target it).
type Pod struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// More info on pod status
	PodStatus PodStatus `json:"podStatus"`

	// Count of containers restarts.
	RestartCount int32 `json:"restartCount"`

	// Pod metrics.
	//Metrics *PodMetrics `json:"metrics"`

	// Pod warning events
	Warnings []Event `json:"warnings"`

	// Name of the Node this Pod runs on.
	NodeName string `json:"nodeName"`
}

//
type PodListPage struct {
	ListMeta int   `json:"total"`
	Pods     []PodPage `json:"rows"`
	//CumulativeMetrics []metricapi.Metric `json:"cumulativeMetrics"`

	// List of non-critical errors, that occurred during resource retrieval.
	//Errors []error `json:"errors"`
}

type PodPage struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metaV1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	//所在节点名称
	NodeName string `json:"nodeName"`
	//phase
	Phase string `json:"phase"`
	//hostIP
	HostIP string `json:"hostIP"`
	//podIP
	PodIP string `json:"podIP"`
	//startTime
	StartTime string `json:"startTime"`
	//restartCount
	RestartCount int32 `json:"restartCount"`

}

//
type PodQueryParam struct {
	models.BaseQueryParam
}

type PodStatus struct {
	Status          string              `json:"status"`
	PodPhase        v1.PodPhase         `json:"podPhase"`
	ContainerStates []v1.ContainerState `json:"containerStates"`
}

//func GetPodList( clusterId ,namespace string) (*v1.PodList, error) {
func GetPodList(clusterId, namespace string, params *PodQueryParam) (*PodListPage, error) {

	clientset, err := models.GetApiServerHandle(clusterId, false)
	if err != nil {
		fmt.Println(err)
	}
	v1Pods, err := clientset.CoreV1().Pods(namespace).List(metaV1.ListOptions{})

	//	return v1Pods,err

	pods := &PodListPage{}
	if err == nil {
		pods.ListMeta = len(v1Pods.Items)
		itemsCount := int64(len(v1Pods.Items))

		//分页索引
		startindex := params.Offset
		endindex := params.Offset + int64(params.Limit)

		if endindex > itemsCount {
			endindex = itemsCount
		}

		if startindex > itemsCount {
			pods.Pods = []PodPage{}
		} else {
			pageList := v1Pods.Items[startindex:endindex]
			for i := 0; i < len(pageList); i++ {
				item := pageList[i]

				nodeName := item.Spec.NodeName
				phase := string(item.Status.Phase)
				hostIP := item.Status.HostIP
				podIP := item.Status.PodIP

				startTime := ""
				if item.Status.StartTime != nil {
					startTime = item.Status.StartTime.Format("")
				}
				restartCount := int32(0)
				for j := 0; j < len(item.Status.ContainerStatuses); j++ {
					if restartCount < item.Status.ContainerStatuses[j].RestartCount {
						restartCount = item.Status.ContainerStatuses[j].RestartCount
					}
				}

				pod := PodPage{ObjectMeta: item.ObjectMeta, NodeName: nodeName, Phase: phase, HostIP: hostIP, PodIP: podIP, StartTime: startTime, RestartCount: restartCount}
				pods.Pods = append(pods.Pods, pod)
			}
		}
	}

	return pods, err
}

func ToPodList(pods []v1.Pod, events []v1.Event) PodList {
	podList := PodList{
		Pods: make([]Pod, 0),
		//Errors: nonCriticalErrors,
	}

	//podCells, cumulativeMetricsPromises, filteredTotal := dataselect.
	//	GenericDataSelectWithFilterAndMetrics(toCells(pods), dsQuery, metricapi.NoResourceCache, metricClient)
	//pods = fromCells(podCells)
	//podList.ListMeta = len(pods)
	podList.ListMeta = ListMeta{TotalItems: len(pods)}

	//metrics, err := getMetricsPerPod(pods, metricClient, dsQuery)
	//if err != nil {
	//	log.Printf("Skipping Heapster metrics because of error: %s\n", err)
	//}

	for _, pod := range pods {
		warnings := GetPodsEventWarnings(events, []v1.Pod{pod})
		podDetail := toPod(&pod, warnings)
		podList.Pods = append(podList.Pods, podDetail)
	}

	//cumulativeMetrics, err := cumulativeMetricsPromises.GetMetrics()
	//podList.CumulativeMetrics = cumulativeMetrics
	//if err != nil {
	//	podList.CumulativeMetrics = make([]metricapi.Metric, 0)
	//}

	return podList
}

// NewObjectMeta returns internal endpoint name for the given service properties, e.g.,
// NewObjectMeta creates a new instance of ObjectMeta struct based on K8s object meta.
func NewObjectMeta(k8SObjectMeta metaV1.ObjectMeta) ObjectMeta {
	return ObjectMeta{
		Name:              k8SObjectMeta.Name,
		Namespace:         k8SObjectMeta.Namespace,
		Labels:            k8SObjectMeta.Labels,
		CreationTimestamp: k8SObjectMeta.CreationTimestamp,
		Annotations:       k8SObjectMeta.Annotations,
	}
}

// NewTypeMeta creates new type mete for the resource kind.
func NewTypeMeta(kind string) TypeMeta {
	return TypeMeta{
		Kind: kind,
	}
}

func toPod(pod *v1.Pod, warnings []Event) Pod {
	podDetail := Pod{
		ObjectMeta:   NewObjectMeta(pod.ObjectMeta),
		TypeMeta:     NewTypeMeta(ResourceKindPod),
		Warnings:     warnings,
		PodStatus:    getPodStatus(*pod, warnings),
		RestartCount: getRestartCount(*pod),
		NodeName:     pod.Spec.NodeName,
	}

	//if m, exists := metrics.MetricsMap[pod.UID]; exists {
	//	podDetail.Metrics = &m
	//}

	return podDetail
}

// getPodStatus returns a PodStatus object containing a summary of the pod's status.
func getPodStatus(pod v1.Pod, warnings []Event) PodStatus {
	var states []v1.ContainerState
	for _, containerStatus := range pod.Status.ContainerStatuses {
		states = append(states, containerStatus.State)
	}

	return PodStatus{
		Status:          string(getPodStatusPhase(pod, warnings)),
		PodPhase:        pod.Status.Phase,
		ContainerStates: states,
	}
}

// getPodStatus returns one of three pod statuses (pending, success, failed)
func getPodStatusPhase(pod v1.Pod, warnings []Event) v1.PodPhase {
	// For terminated pods that failed
	if pod.Status.Phase == v1.PodFailed {
		return v1.PodFailed
	}

	// For successfully terminated pods
	if pod.Status.Phase == v1.PodSucceeded {
		return v1.PodSucceeded
	}

	ready := false
	initialized := false
	for _, c := range pod.Status.Conditions {
		if c.Type == v1.PodReady {
			ready = c.Status == v1.ConditionTrue
		}
		if c.Type == v1.PodInitialized {
			initialized = c.Status == v1.ConditionTrue
		}
	}

	if initialized && ready && pod.Status.Phase == v1.PodRunning {
		return v1.PodRunning
	}

	// If the pod would otherwise be pending but has warning then label it as
	// failed and show and error to the user.
	if len(warnings) > 0 {
		return v1.PodFailed
	}

	// Unknown?
	return v1.PodPending
}

// Gets restart count of given pod (total number of its containers restarts).
func getRestartCount(pod v1.Pod) int32 {
	var restartCount int32 = 0
	for _, containerStatus := range pod.Status.ContainerStatuses {
		restartCount += containerStatus.RestartCount
	}
	return restartCount
}
