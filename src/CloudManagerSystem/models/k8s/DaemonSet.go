package k8s

import (
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	apps "k8s.io/api/apps/v1beta2"
	"CloudManagerSystem/models"
)

// DaemonSetList contains a list of Daemon Sets in the cluster.
type DaemonSetList struct {
	ListMeta          int       `json:"total"`
	DaemonSets        []DaemonSet        `json:"rows"`
	//CumulativeMetrics []metricapi.Metric `json:"cumulativeMetrics"`

	// Basic information about resources status on the list.
	Status ResourceStatus `json:"status"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

// DaemonSet plus zero or more Kubernetes services that target the Daemon Set.
type DaemonSet struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// Aggregate information about pods belonging to this Daemon Set.
	Pods PodInfo `json:"pods"`

	// Container images of the Daemon Set.
	ContainerImages []string `json:"containerImages"`

	// InitContainer images of the Daemon Set.
	InitContainerImages []string `json:"initContainerImages"`
}
type DaemonSetQueryParam struct {
	models.BaseQueryParam
}

// GetDaemonSetList returns a list of all Daemon Set in the cluster.
func GetDaemonSetList(client kubernetes.Interface, namespace string, dsQuery *DaemonSetQueryParam) (*DaemonSetList, error) {
	channels := &ResourceChannels{
		DaemonSetList: GetDaemonSetListChannel(client, namespace, 1),
		ServiceList:   GetServiceListChannel(client, namespace, 1),
		PodList:       GetPodListChannel(client, namespace),
		EventList:     GetEventListChannel(client, namespace),
	}

	return GetDaemonSetListFromChannels(channels, dsQuery)
}

// GetDaemonSetListFromChannels returns a list of all Daemon Set in the cluster
// reading required resource list once from the channels.
func GetDaemonSetListFromChannels(channels *ResourceChannels, dsQuery *DaemonSetQueryParam) (*DaemonSetList, error) {

	daemonSets := <-channels.DaemonSetList.List
	err := <-channels.DaemonSetList.Error
	//nonCriticalErrors, criticalError := errors.HandleError(err)
	//	//if criticalError != nil {
	//	//	return nil, criticalError
	//	//}

	pods := <-channels.PodList.List
	err = <-channels.PodList.Error
	//nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	events := <-channels.EventList.List
	err = <-channels.EventList.Error
	//nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	dsList := toDaemonSetList(daemonSets.Items, pods.Items, events.Items,  dsQuery)
	dsList.Status = getDaemonSetStatus(daemonSets, pods.Items, events.Items)
	dsList.Errors = append(dsList.Errors,err)
	return dsList, nil
}

func toDaemonSetList(daemonSets []apps.DaemonSet, pods []v1.Pod, events []v1.Event, params *DaemonSetQueryParam ) *DaemonSetList {

	daemonSetList := &DaemonSetList{
		DaemonSets: make([]DaemonSet, 0),
		ListMeta:    len(daemonSets),
		//Errors:     nonCriticalErrors,
	}

	//cachedResources := &CachedResources{
	//	Pods: pods,
	//}
	//
	//dsCells, metricPromises, filteredTotal := dataselect.GenericDataSelectWithFilterAndMetrics(ToCells(daemonSets),
	//	dsQuery, cachedResources, metricClient)
	//daemonSets = FromCells(dsCells)
	daemonSetList.ListMeta = len(daemonSets)//api.ListMeta{TotalItems: filteredTotal}

	itemsCount := int64(len(daemonSets))
	//分页索引
	startindex := params.Offset
	endindex := params.Offset + int64(params.Limit)
	if endindex > itemsCount {
		endindex = itemsCount
	}

	if startindex > itemsCount {
		daemonSetList.DaemonSets = []DaemonSet{}
	} else {
		pageList := daemonSets[startindex:endindex]

		for _, daemonSet := range pageList {
			matchingPods := FilterPodsByControllerRef(&daemonSet, pods)
			podInfo := GetPodInfo(daemonSet.Status.CurrentNumberScheduled,
				&daemonSet.Status.DesiredNumberScheduled, matchingPods)
			podInfo.Warnings = GetPodsEventWarnings(events, matchingPods)

			daemonSetList.DaemonSets = append(daemonSetList.DaemonSets, DaemonSet{
				ObjectMeta:          NewObjectMeta(daemonSet.ObjectMeta),
				TypeMeta:            NewTypeMeta(ResourceKindDaemonSet),
				Pods:                podInfo,
				ContainerImages:     GetContainerImages(&daemonSet.Spec.Template.Spec),
				InitContainerImages: GetInitContainerImages(&daemonSet.Spec.Template.Spec),
			})
		}
	}

	//cumulativeMetrics, err := metricPromises.GetMetrics()
	//daemonSetList.CumulativeMetrics = cumulativeMetrics
	//if err != nil {
	//	daemonSetList.CumulativeMetrics = make([]metricapi.Metric, 0)
	//}

	return daemonSetList
}


func getDaemonSetStatus(list *apps.DaemonSetList, pods []v1.Pod, events []v1.Event) ResourceStatus {
	info := ResourceStatus{}
	if list == nil {
		return info
	}

	for _, daemonSet := range list.Items {
		matchingPods := FilterPodsByControllerRef(&daemonSet, pods)
		podInfo := GetPodInfo(daemonSet.Status.CurrentNumberScheduled,
			&daemonSet.Status.DesiredNumberScheduled, matchingPods)
		warnings := GetPodsEventWarnings(events, matchingPods)

		if len(warnings) > 0 {
			info.Failed++
		} else if podInfo.Pending > 0 {
			info.Pending++
		} else {
			info.Running++
		}
	}

	return info
}