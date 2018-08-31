package k8s

import (
	"CloudManagerSystem/models"
	"log"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	apps "k8s.io/api/apps/v1beta2"
)

// StatefulSetList contains a list of Stateful Sets in the cluster.
type StatefulSetList struct {
	ListMeta int `json:"total"`

	// Basic information about resources status on the list.
	Status ResourceStatus `json:"status"`

	// Unordered list of Pet Sets.
	StatefulSets      []StatefulSet      `json:"rows"`
	//CumulativeMetrics []metricapi.Metric `json:"cumulativeMetrics"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

// StatefulSet is a presentation layer view of Kubernetes Stateful Set resource. This means it is
// Stateful Set plus additional augmented data we can get from other sources (like services that
// target the same pods).
type StatefulSet struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// Aggregate information about pods belonging to this Pet Set.
	Pods PodInfo `json:"pods"`

	// Container images of the Stateful Set.
	ContainerImages []string `json:"containerImages"`

	// Init container images of the Stateful Set.
	InitContainerImages []string `json:"initContainerImages"`
}

type StatefulSetQueryParam struct {
	models.BaseQueryParam
}

// GetStatefulSetList returns a list of all Stateful Sets in the cluster.
func GetStatefulSetList(client kubernetes.Interface, namespace string, dsQuery *StatefulSetQueryParam) (*StatefulSetList, error) {
	log.Print("Getting list of all pet sets in the cluster")

	channels := &ResourceChannels{
		StatefulSetList: GetStatefulSetListChannel(client, namespace, 1),
		PodList:         GetPodListChannel(client, namespace),
		EventList:       GetEventListChannel(client, namespace),
	}

	return GetStatefulSetListFromChannels(channels, dsQuery)
}

// GetStatefulSetListFromChannels returns a list of all Stateful Sets in the cluster reading
// required resource list once from the channels.
func GetStatefulSetListFromChannels(channels *ResourceChannels,dsQuery *StatefulSetQueryParam) (*StatefulSetList, error) {

	statefulSets := <-channels.StatefulSetList.List
	err := <-channels.StatefulSetList.Error
	//nonCriticalErrors, criticalError := errors.HandleError(err)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

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


	ssList := toStatefulSetList(statefulSets.Items, pods.Items, events.Items,dsQuery)
	ssList.Status = getStatefulStatus(statefulSets, pods.Items, events.Items)
	ssList.Errors = append(ssList.Errors,err)
	return ssList, nil
}

func toStatefulSetList(statefulSets []apps.StatefulSet, pods []v1.Pod, events []v1.Event,params *StatefulSetQueryParam) *StatefulSetList {

	statefulSetList := &StatefulSetList{
		StatefulSets: make([]StatefulSet, 0),
		ListMeta:     len(statefulSets),//api.ListMeta{TotalItems: len(statefulSets)},
		//Errors:       nonCriticalErrors,
	}

	//cachedResources := &metricapi.CachedResources{
	//	Pods: pods,
	//}
	//ssCells, metricPromises, filteredTotal := dataselect.GenericDataSelectWithFilterAndMetrics(
	//	toCells(statefulSets), dsQuery, cachedResources, metricClient)
	//statefulSets = fromCells(ssCells)
	statefulSetList.ListMeta = len(statefulSets) //api.ListMeta{TotalItems: filteredTotal}


	itemsCount := int64(len(statefulSets))
	//分页索引
	startindex := params.Offset
	endindex := params.Offset + int64(params.Limit)
	if endindex > itemsCount {
		endindex = itemsCount
	}

	if startindex > itemsCount {
		statefulSetList.StatefulSets = []StatefulSet{}
	} else {
		pageList := statefulSets[startindex:endindex]

		for _, statefulSet := range pageList {
			///statefulSetList.StatefulSets = append(statefulSetList.StatefulSets, toStorageClass(&item))
			matchingPods := FilterPodsByControllerRef(&statefulSet, pods)
			podInfo := GetPodInfo(statefulSet.Status.Replicas, statefulSet.Spec.Replicas, matchingPods)
			podInfo.Warnings = GetPodsEventWarnings(events, matchingPods)
			statefulSetList.StatefulSets = append(statefulSetList.StatefulSets, toStatefulSet(&statefulSet, &podInfo))
		}
	}


	//cumulativeMetrics, err := metricPromises.GetMetrics()
	//statefulSetList.CumulativeMetrics = cumulativeMetrics
	//if err != nil {
	//	statefulSetList.CumulativeMetrics = make([]metricapi.Metric, 0)
	//}

	return statefulSetList
}

func toStatefulSet(statefulSet *apps.StatefulSet, podInfo *PodInfo) StatefulSet {
	return StatefulSet{
		ObjectMeta:          NewObjectMeta(statefulSet.ObjectMeta),
		TypeMeta:            NewTypeMeta(ResourceKindStatefulSet),
		ContainerImages:     GetContainerImages(&statefulSet.Spec.Template.Spec),
		InitContainerImages: GetInitContainerImages(&statefulSet.Spec.Template.Spec),
		Pods:                *podInfo,
	}
}


func getStatefulStatus(list *apps.StatefulSetList, pods []v1.Pod, events []v1.Event) ResourceStatus {
	info := ResourceStatus{}
	if list == nil {
		return info
	}

	for _, ss := range list.Items {
		matchingPods := FilterPodsByControllerRef(&ss, pods)
		podInfo := GetPodInfo(ss.Status.Replicas, ss.Spec.Replicas, matchingPods)
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
