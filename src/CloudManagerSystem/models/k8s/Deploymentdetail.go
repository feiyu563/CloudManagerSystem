package k8s

import (
	"log"

	apps "k8s.io/api/apps/v1beta2"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	client "k8s.io/client-go/kubernetes"
	//"fmt"
)

// RollingUpdateStrategy is behavior of a rolling update. See RollingUpdateDeployment K8s object.
type RollingUpdateStrategy struct {
	MaxSurge       *intstr.IntOrString `json:"maxSurge"`
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable"`
}

type StatusInfo struct {
	// Total number of desired replicas on the deployment
	Replicas int32 `json:"replicas"`

	// Number of non-terminated pods that have the desired template spec
	Updated int32 `json:"updated"`

	// Number of available pods (ready for at least minReadySeconds)
	// targeted by this deployment
	Available int32 `json:"available"`

	// Total number of unavailable pods targeted by this deployment.
	Unavailable int32 `json:"unavailable"`
}

// DeploymentDetail is a presentation layer view of Kubernetes Deployment resource.
type DeploymentDetail struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// Detailed information about Pods belonging to this Deployment.
	PodList PodList `json:"podList"`

	// Label selector of the service.
	Selector map[string]string `json:"selector"`

	// Status information on the deployment
	StatusInfo `json:"statusInfo"`

	// The deployment strategy to use to replace existing pods with new ones.
	// Valid options: Recreate, RollingUpdate
	Strategy apps.DeploymentStrategyType `json:"strategy"`

	// Min ready seconds
	MinReadySeconds int32 `json:"minReadySeconds"`

	// Rolling update strategy containing maxSurge and maxUnavailable
	RollingUpdateStrategy *RollingUpdateStrategy `json:"rollingUpdateStrategy,omitempty"`

	// RepliaSetList containing old replica sets from the deployment
	OldReplicaSetList ReplicaSetList `json:"oldReplicaSetList"`

	// New replica set used by this deployment
	NewReplicaSet ReplicaSet `json:"newReplicaSet"`

	// Optional field that specifies the number of old Replica Sets to retain to allow rollback.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit"`

	// List of events related to this Deployment
	EventList EventList `json:"eventList"`

	// List of Horizontal Pod AutoScalers targeting this Deployment
	HorizontalPodAutoscalerList HorizontalPodAutoscalerList `json:"horizontalPodAutoscalerList"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

type DeploymentDetailQueryParam struct {
	Name   string
	Namespace string
}

// GetDeploymentDetail returns model object of deployment and error, if any.
func GetDeploymentDetail(params *DeploymentDetailQueryParam, client client.Interface, namespace string, deploymentName string) (*DeploymentDetail, error) {

	log.Printf("Getting details of %s deployment in %s namespace", deploymentName, namespace)

	deployment, err := client.AppsV1beta2().Deployments(namespace).Get(deploymentName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	//fmt.Println("#####################",deployment)
	selector, err := metaV1.LabelSelectorAsSelector(deployment.Spec.Selector)
	if err != nil {
		return nil, err
	}
	options := metaV1.ListOptions{LabelSelector: selector.String()}

	channels := &ResourceChannels{
		ReplicaSetList: GetReplicaSetListChannelWithOptions(client, namespace, options),
		PodList:        GetPodListChannelWithOptions(client, namespace, options),
	}

	rawRs := <-channels.ReplicaSetList.List
	err = <-channels.ReplicaSetList.Error
	//nonCriticalErrors, criticalError := errors.HandleError(err)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	rawPods := <-channels.PodList.List
	err = <-channels.PodList.Error
	//nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	podList, err := GetDeploymentPods(client, namespace, deploymentName)
	//nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	eventList, err := GetResourceEvents(client, namespace, deploymentName)
	//nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	hpas, err := GetHorizontalPodAutoscalerListForResource(client, namespace, "Deployment", deploymentName)
	//nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	oldReplicaSetList, err := GetDeploymentOldReplicaSets(client, namespace, deploymentName)
	//nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	rawRepSets := make([]*apps.ReplicaSet, 0)
	for i := range rawRs.Items {
		rawRepSets = append(rawRepSets, &rawRs.Items[i])
	}
	newRs, err := FindNewReplicaSet(deployment, rawRepSets)
	//nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	var newReplicaSet ReplicaSet
	if newRs != nil {
		matchingPods := FilterPodsByControllerRef(newRs, rawPods.Items)
		newRsPodInfo := GetPodInfo(newRs.Status.Replicas, newRs.Spec.Replicas, matchingPods)
		events, err := GetPodsEvents(client, namespace, matchingPods)
		if err != nil {
			return nil, err
		}
		//nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
		//if criticalError != nil {
		//	return nil, criticalError
		//}

		newRsPodInfo.Warnings = GetPodsEventWarnings(events, matchingPods)
		newReplicaSet = ToReplicaSet(newRs, &newRsPodInfo)
	}

	// Extra Info
	var rollingUpdateStrategy *RollingUpdateStrategy
	if deployment.Spec.Strategy.RollingUpdate != nil {
		rollingUpdateStrategy = &RollingUpdateStrategy{
			MaxSurge:       deployment.Spec.Strategy.RollingUpdate.MaxSurge,
			MaxUnavailable: deployment.Spec.Strategy.RollingUpdate.MaxUnavailable,
		}
	}

	return &DeploymentDetail{
		ObjectMeta:                  NewObjectMeta(deployment.ObjectMeta),
		TypeMeta:                    NewTypeMeta(ResourceKindDeployment),
		PodList:                     *podList,
		Selector:                    deployment.Spec.Selector.MatchLabels,
		StatusInfo:                  GetStatusInfo(&deployment.Status),
		Strategy:                    deployment.Spec.Strategy.Type,
		MinReadySeconds:             deployment.Spec.MinReadySeconds,
		RollingUpdateStrategy:       rollingUpdateStrategy,
		OldReplicaSetList:           *oldReplicaSetList,
		NewReplicaSet:               newReplicaSet,
		RevisionHistoryLimit:        deployment.Spec.RevisionHistoryLimit,
		EventList:                   *eventList,
		HorizontalPodAutoscalerList: *hpas,
		Errors:                      nil,
	}, nil

}

func GetStatusInfo(deploymentStatus *apps.DeploymentStatus) StatusInfo {
	return StatusInfo{
		Replicas:    deploymentStatus.Replicas,
		Updated:     deploymentStatus.UpdatedReplicas,
		Available:   deploymentStatus.AvailableReplicas,
		Unavailable: deploymentStatus.UnavailableReplicas,
	}
}
