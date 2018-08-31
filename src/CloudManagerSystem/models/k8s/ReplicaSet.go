package k8s

import (
	//"sort"
	//"CloudManagerSystem/models"
	//"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apps "k8s.io/api/apps/v1beta2"
	//"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
	client "k8s.io/client-go/kubernetes"
)

// ReplicaSetList contains a list of Replica Sets in the cluster.
type ReplicaSetList struct {
	ListMeta ListMeta `json:"listMeta"`
	//CumulativeMetrics []Metric `json:"cumulativeMetrics"`

	// Basic information about resources status on the list.
	Status ResourceStatus `json:"status"`

	// Unordered list of Replica Sets.
	ReplicaSets []ReplicaSet `json:"replicaSets"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

// ReplicaSet is a presentation layer view of Kubernetes Replica Set resource. This means
// it is Replica Set plus additional augmented data we can get from other sources
// (like services that target the same pods).
type ReplicaSet struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// Aggregate information about pods belonging to this Replica Set.
	Pods PodInfo `json:"pods"`

	// Container images of the Replica Set.
	ContainerImages []string `json:"containerImages"`

	// Init Container images of the Replica Set.
	InitContainerImages []string `json:"initContainerImages"`
}

// ToReplicaSet converts replica set api object to replica set model object.
func ToReplicaSet(replicaSet *apps.ReplicaSet, podInfo *PodInfo) ReplicaSet {
	return ReplicaSet{
		ObjectMeta:          NewObjectMeta(replicaSet.ObjectMeta),
		TypeMeta:            NewTypeMeta(ResourceKindReplicaSet),
		ContainerImages:     GetContainerImages(&replicaSet.Spec.Template.Spec),
		InitContainerImages: GetInitContainerImages(&replicaSet.Spec.Template.Spec),
		Pods:                *podInfo,
	}
}

// GetContainerImages returns container image strings from the given pod spec.
func GetContainerImages(podTemplate *v1.PodSpec) []string {
	var containerImages []string
	for _, container := range podTemplate.Containers {
		containerImages = append(containerImages, container.Image)
	}
	return containerImages
}

// GetInitContainerImages returns init container image strings from the given pod spec.
func GetInitContainerImages(podTemplate *v1.PodSpec) []string {
	var initContainerImages []string
	for _, initContainer := range podTemplate.InitContainers {
		initContainerImages = append(initContainerImages, initContainer.Image)
	}
	return initContainerImages
}

//GetDeploymentOldReplicaSets returns old replica sets targeting Deployment with given name
func GetDeploymentOldReplicaSets(client client.Interface, namespace string, deploymentName string) (*ReplicaSetList, error) {

	oldReplicaSetList := &ReplicaSetList{
		ReplicaSets: make([]ReplicaSet, 0),
		ListMeta:    ListMeta{TotalItems: 0},
	}

	deployment, err := client.AppsV1beta2().Deployments(namespace).Get(deploymentName, metaV1.GetOptions{})
	if err != nil {
		return oldReplicaSetList, err
	}

	selector, err := metaV1.LabelSelectorAsSelector(deployment.Spec.Selector)
	if err != nil {
		return oldReplicaSetList, err
	}
	options := metaV1.ListOptions{LabelSelector: selector.String()}

	channels := &ResourceChannels{
		ReplicaSetList: GetReplicaSetListChannelWithOptions(client,
			namespace, options),
		PodList: GetPodListChannelWithOptions(client,
			namespace, options),
		EventList: GetEventListChannelWithOptions(client,
			namespace, options),
	}

	rawRs := <-channels.ReplicaSetList.List
	if err := <-channels.ReplicaSetList.Error; err != nil {
		return oldReplicaSetList, err
	}

	rawPods := <-channels.PodList.List
	if err := <-channels.PodList.Error; err != nil {
		return oldReplicaSetList, err
	}

	rawEvents := <-channels.EventList.List
	err = <-channels.EventList.Error
	//nonCriticalErrors, criticalError := errors.HandleError(err)
	//if criticalError != nil {
	//	return oldReplicaSetList, criticalError
	//}

	rawRepSets := make([]*apps.ReplicaSet, 0)
	for i := range rawRs.Items {
		rawRepSets = append(rawRepSets, &rawRs.Items[i])
	}
	oldRs, _, err := FindOldReplicaSets(deployment, rawRepSets)
	if err != nil {
		return oldReplicaSetList, err
	}

	oldReplicaSets := make([]apps.ReplicaSet, len(oldRs))
	for i, replicaSet := range oldRs {
		oldReplicaSets[i] = *replicaSet
	}

	oldReplicaSetList = ToReplicaSetList(oldReplicaSets, rawPods.Items, rawEvents.Items,
		nil)
	return oldReplicaSetList, nil
}

// ToReplicaSetList creates paginated list of Replica Set model
// objects based on Kubernetes Replica Set objects array and related resources arrays.
func ToReplicaSetList(replicaSets []apps.ReplicaSet, pods []v1.Pod, events []v1.Event, nonCriticalErrors []error,
) *ReplicaSetList {

	replicaSetList := &ReplicaSetList{
		ReplicaSets: make([]ReplicaSet, 0),
		ListMeta:    ListMeta{TotalItems: len(replicaSets)},
		Errors:      nonCriticalErrors,
	}

	replicaSetList.ListMeta = ListMeta{TotalItems: len(replicaSets)}

	for _, replicaSet := range replicaSets {
		matchingPods := FilterPodsByControllerRef(&replicaSet, pods)
		podInfo := GetPodInfo(replicaSet.Status.Replicas, replicaSet.Spec.Replicas,
			matchingPods)
		podInfo.Warnings = GetPodsEventWarnings(events, matchingPods)
		replicaSetList.ReplicaSets = append(replicaSetList.ReplicaSets,
			ToReplicaSet(&replicaSet, &podInfo))
	}

	return replicaSetList
}

// FindNewReplicaSet returns the new RS this given deployment targets (the one with the same pod template).
func FindNewReplicaSet(deployment *apps.Deployment, rsList []*apps.ReplicaSet) (*apps.ReplicaSet, error) {
	newRSTemplate := GetNewReplicaSetTemplate(deployment)
	for i := range rsList {
		if EqualIgnoreHash(rsList[i].Spec.Template, newRSTemplate) {
			// This is the new ReplicaSet.
			return rsList[i], nil
		}
	}
	// new ReplicaSet does not exist.
	return nil, nil
}

// FindOldReplicaSets returns the old replica sets targeted by the given Deployment, with the given slice of RSes.
// Note that the first set of old replica sets doesn't include the ones with no pods, and the second set of old replica
// sets include all old replica sets.
func FindOldReplicaSets(deployment *apps.Deployment, rsList []*apps.ReplicaSet) ([]*apps.ReplicaSet,
	[]*apps.ReplicaSet, error) {
	var requiredRSs []*apps.ReplicaSet
	var allRSs []*apps.ReplicaSet
	newRS, err := FindNewReplicaSet(deployment, rsList)
	if err != nil {
		return nil, nil, err
	}
	for _, rs := range rsList {
		// Filter out new replica set
		if newRS != nil && rs.UID == newRS.UID {
			continue
		}
		allRSs = append(allRSs, rs)
		if *(rs.Spec.Replicas) != 0 {
			requiredRSs = append(requiredRSs, rs)
		}
	}
	return requiredRSs, allRSs, nil
}
