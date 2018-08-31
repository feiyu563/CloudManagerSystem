package k8s

import (
	"CloudManagerSystem/models"
	//"fmt"
	api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apps "k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	client "k8s.io/client-go/kubernetes"
	"errors"
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
)

type DeploymentQueryParam struct {
	models.BaseQueryParam
	Namespace string
}

const (
	// DescriptionAnnotationKey is annotation key for a description.
	DescriptionAnnotationKey = "description"
)

// PortMapping is a specification of port mapping for an application deployment.
type PortMapping struct {
	// Port that will be exposed on the service.
	Port int32 `json:"port"`

	// Docker image path for the application.
	TargetPort int32 `json:"targetPort"`

	// IP protocol for the mapping, e.g., "TCP" or "UDP".
	Protocol api.Protocol `json:"protocol"`
}

// EnvironmentVariable represents a named variable accessible for containers.
type EnvironmentVariable struct {
	// Name of the variable. Must be a C_IDENTIFIER.
	Name string `json:"name"`

	// Value of the variable, as defined in Kubernetes core API.
	Value string `json:"value"`
}

// Label is a structure representing label assignable to Pod/RC/Service
type Label struct {
	// Label key
	Key string `json:"key"`

	// Label value
	Value string `json:"value"`
}

// AppDeploymentSpec is a specification for an app deployment.
type AppDeploymentSpec struct {
	// Name of the application.
	Name string `json:"name"`

	// Docker image path for the application.
	ContainerImage string `json:"containerImage"`

	// The name of an image pull secret in case of a private docker repository.
	ImagePullSecret *string `json:"imagePullSecret"`

	// Command that is executed instead of container entrypoint, if specified.
	ContainerCommand *string `json:"containerCommand"`

	// Arguments for the specified container command or container entrypoint (if command is not
	// specified here).
	ContainerCommandArgs *string `json:"containerCommandArgs"`

	// Number of replicas of the image to maintain.
	Replicas int32 `json:"replicas"`

	// Port mappings for the service that is created. The service is created if there is at least
	// one port mapping.
	PortMappings []PortMapping `json:"portMappings"`

	// List of user-defined environment variables.
	Variables []EnvironmentVariable `json:"variables"`

	// Whether the created service is external.
	IsExternal bool `json:"isExternal"`

	// Description of the deployment.
	Description *string `json:"description"`

	// Target namespace of the application.
	Namespace string `json:"namespace"`

	// Optional memory requirement for the container.
	MemoryRequirement *resource.Quantity `json:"memoryRequirement"`

	// Optional CPU requirement for the container.
	CpuRequirement *resource.Quantity `json:"cpuRequirement"`

	// Labels that will be defined on Pods/RCs/Services
	Labels []Label `json:"labels"`

	// Whether to run the container as privileged user (essentially equivalent to root on the host).
	RunAsPrivileged bool `json:"runAsPrivileged"`
}

// ObjectMeta is metadata about an instance of a resource.
type ObjectMeta struct {
	// Name is unique within a namespace. Name is primarily intended for creation
	// idempotence and configuration definition.
	Name string `json:"name,omitempty"`

	// Namespace defines the space within which name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	Namespace string `json:"namespace,omitempty"`

	// Labels are key value pairs that may be used to scope and select individual resources.
	// Label keys are of the form:
	//     label-key ::= prefixed-name | name
	//     prefixed-name ::= prefix '/' name
	//     prefix ::= DNS_SUBDOMAIN
	//     name ::= DNS_LABEL
	// The prefix is optional.  If the prefix is not specified, the key is assumed to be private
	// to the user.  Other system components that wish to use labels must specify a prefix.
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations are unstructured key value data stored with a resource that may be set by
	// external tooling. They are not queryable and should be preserved when modifying
	// objects.  Annotation keys have the same formatting restrictions as Label keys. See the
	// comments on Labels for details.
	Annotations map[string]string `json:"annotations,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was
	// created. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	CreationTimestamp metaV1.Time `json:"creationTimestamp,omitempty"`
}

type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// In smalllettercase.
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`
	//Kind ResourceKind `json:"kind,omitempty"`
}

// PodInfo represents aggregate information about controller's pods.
type PodInfo struct {
	// Number of pods that are created.
	Current int32 `json:"current"`

	// Number of pods that are desired.
	Desired *int32 `json:"desired,omitempty"`

	// Number of pods that are currently running.
	Running int32 `json:"running"`

	// Number of pods that are currently waiting.
	Pending int32 `json:"pending"`

	// Number of pods that are failed.
	Failed int32 `json:"failed"`

	// Number of pods that are succeeded.
	Succeeded int32 `json:"succeeded"`

	// Unique warning messages related to pods in this resource.
	Warnings []Event `json:"warnings"`
}

type Deployment struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// Aggregate information about pods belonging to this Deployment.
	Pods PodInfo `json:"pods"`

	// Container images of the Deployment.
	ContainerImages []string `json:"containerImages"`

	// Init Container images of the Deployment.
	InitContainerImages []string `json:"initContainerImages"`
}

type ListMeta struct {
	// Total number of items on the list. Used for pagination.
	TotalItems int `json:"totalItems"`
}

// ResourceStatus provides basic information about resources status on the list.
type ResourceStatus struct {
	// Number of resources that are currently in running state.
	Running int `json:"running"`

	// Number of resources that are currently in pending state.
	Pending int `json:"pending"`

	// Number of resources that are in failed state.
	Failed int `json:"failed"`

	// Number of resources that are in succeeded state.
	Succeeded int `json:"succeeded"`
}

// DeploymentList contains a list of Deployments in the cluster.
type DeploymentList struct {
	ListMeta int `json:"total"`
	//CumulativeMetrics []metricapi.Metric `json:"cumulativeMetrics"`

	// Basic information about resources status on the list.
	//Status ResourceStatus `json:"status"`

	// Unordered list of Deployments.
	Deployments []Deployment `json:"rows"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

func GetPodInfo(current int32, desired *int32, pods []v1.Pod) PodInfo {
	result := PodInfo{
		Current: current,
		Desired: desired,
		//Warnings: make([]Event, 0),
	}

	for _, pod := range pods {
		switch pod.Status.Phase {
		case v1.PodRunning:
			result.Running++
		case v1.PodPending:
			result.Pending++
		case v1.PodFailed:
			result.Failed++
		case v1.PodSucceeded:
			result.Succeeded++
		}
	}

	return result
}

func GetDeploymentList(params *DeploymentQueryParam, client client.Interface, namespace string) *DeploymentList {
	channels := &ResourceChannels{
		DeploymentList: GetDeploymentListChannel(client, namespace),
		PodList:        GetPodListChannel(client, namespace),
		EventList:      GetEventListChannel(client, namespace),
		ReplicaSetList: GetReplicaSetListChannel(client, namespace),
	}
	return GetDeploymentListFromChannels(params, channels)
}

func GetDeploymentListFromChannels(params *DeploymentQueryParam, channels *ResourceChannels) *DeploymentList {
	var deployment Deployment

	deployments := <-channels.DeploymentList.List
	pods := <-channels.PodList.List
	rs := <-channels.ReplicaSetList.List

	deploymentList := &DeploymentList{
		Deployments: make([]Deployment, 0),
		ListMeta:    len(deployments.Items),
		//ListMeta:    ListMeta{TotalItems: len(deployments.Items)},
		//Errors:      nonCriticalErrors,
	}
	total := len(deployments.Items)
	if params.Limit == 0 {
		params.Limit = 10
	} else if params.Limit > total {
		params.Limit = total
	}

	startindex := params.Offset
	endindex := params.Offset + int64(params.Limit)

	if endindex > int64(total) {
		endindex = int64(total)
	}
	paginate := deployments.Items[startindex:endindex]

	for _, item1 := range paginate {

		deployment.ObjectMeta.Name = item1.Name
		deployment.ObjectMeta.Namespace = item1.Namespace
		deployment.ObjectMeta.Labels = item1.Labels
		deployment.ObjectMeta.CreationTimestamp = item1.CreationTimestamp

		var containerImages []string
		for _, container := range item1.Spec.Template.Spec.Containers {
			containerImages = append(containerImages, container.Image)
		}
		deployment.ContainerImages = containerImages
		var initContainerImages []string
		for _, initContainer := range item1.Spec.Template.Spec.InitContainers {
			initContainerImages = append(initContainerImages, initContainer.Image)
		}
		matchingPods := FilterDeploymentPodsByOwnerReference(item1, rs.Items, pods.Items)
		deployment.InitContainerImages = initContainerImages
		podInfo := GetPodInfo(item1.Status.Replicas, item1.Spec.Replicas, matchingPods)
		deployment.Pods = podInfo

		deploymentList.Deployments = append(deploymentList.Deployments, deployment)
	}

	return deploymentList
}

func FilterDeploymentPodsByOwnerReference(deployment apps.Deployment, allRS []apps.ReplicaSet,
	allPods []v1.Pod) []v1.Pod {
	var matchingPods []v1.Pod

	rsTemplate := v1.PodTemplateSpec{
		ObjectMeta: deployment.Spec.Template.ObjectMeta,
		Spec:       deployment.Spec.Template.Spec,
	}

	for _, rs := range allRS {
		if EqualIgnoreHash(rs.Spec.Template, rsTemplate) {
			matchingPods = FilterPodsByControllerRef(&rs, allPods)
		}
	}

	return matchingPods
}

func EqualIgnoreHash(template1, template2 v1.PodTemplateSpec) bool {
	// First, compare template.Labels (ignoring hash)
	labels1, labels2 := template1.Labels, template2.Labels
	if len(labels1) > len(labels2) {
		labels1, labels2 = labels2, labels1
	}
	// We make sure len(labels2) >= len(labels1)
	for k, v := range labels2 {
		if labels1[k] != v && k != apps.DefaultDeploymentUniqueLabelKey {
			return false
		}
	}
	// Then, compare the templates without comparing their labels
	template1.Labels, template2.Labels = nil, nil
	return equality.Semantic.DeepEqual(template1, template2)
}

// FilterPodsByControllerRef returns a subset of pods controlled by given controller resource, excluding deployments.
func FilterPodsByControllerRef(owner metaV1.Object, allPods []v1.Pod) []v1.Pod {
	var matchingPods []v1.Pod
	for _, pod := range allPods {
		if metaV1.IsControlledBy(&pod, owner) {
			matchingPods = append(matchingPods, pod)
		}
	}
	return matchingPods
}

//
func GetDeploymentPods(client client.Interface, namespace, deploymentName string) (*PodList, error) {

	deployment, err := client.AppsV1beta2().Deployments(namespace).Get(deploymentName, metaV1.GetOptions{})
	if err != nil {
		return EmptyPodList, err
	}

	channels := &ResourceChannels{
		PodList:        GetPodListChannel(client, namespace),
		ReplicaSetList: GetReplicaSetListChannel(client, namespace),
	}

	rawPods := <-channels.PodList.List
	if err := <-channels.PodList.Error; err != nil {
		return EmptyPodList, err
	}

	rawRs := <-channels.ReplicaSetList.List
	err = <-channels.ReplicaSetList.Error
	//nonCriticalErrors, criticalError := errors.HandleError(err)
	//if criticalError != nil {
	//	return EmptyPodList, criticalError
	//}

	pods := FilterDeploymentPodsByOwnerReference(*deployment, rawRs.Items, rawPods.Items)
	events, err := GetPodsEvents(client, namespace, pods)
	//nonCriticalErrors, criticalError = AppendError(err, nonCriticalErrors)
	//if criticalError != nil {
	//	return pod.EmptyPodList, criticalError
	//}

	podList := ToPodList(pods, events)
	return &podList, nil
}

func GetNewReplicaSetTemplate(deployment *apps.Deployment) v1.PodTemplateSpec {
	// newRS will have the same template as in deployment spec.
	return v1.PodTemplateSpec{
		ObjectMeta: deployment.Spec.Template.ObjectMeta,
		Spec:       deployment.Spec.Template.Spec,
	}
}

/*

	deployment1 := &apps.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: service.Name,
		},
		Spec: apps.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": service.Name,
				},
			},
			Template: api.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: map[string]string{
						"app": service.Name,
					},
				},
				Spec: api.PodSpec{
					Containers: []api.Container{
						{
							Name:            "node-service",
							Image:           "harbor.zxbike.cn/nodejs/node-alpine:v1.0.1",
							ImagePullPolicy: api.PullIfNotPresent,
							Resources:       r,
							Ports: []api.ContainerPort{
								{
									//Name:          "http",
									Protocol:      api.ProtocolTCP,
									ContainerPort: 3040,
								},
							},
							LivenessProbe: &api.Probe{
								Handler: api.Handler{
									HTTPGet: &api.HTTPGetAction{
										Path:   "",
										Port:   intstr.FromInt(3025),
										Scheme: "",
									},
								},
								InitialDelaySeconds: 15,
								PeriodSeconds:       5,
							},
						},
					},
				},
			},
		},
	}

*/
func int32Ptr(i int32) *int32 { return &i }

// DeployApp deploys an app based on the given configuration. The app is deployed using the given
// client. App deployment consists of a deployment and an optional service. Both of them
// share common labels.
func DeployApp(client client.Interface, puborrollback *models.KubeServicePubORRollback, ClusterId string, Pubflag bool) error {
	//fmt.Printf("Deploying %s application into %s namespace", spec.Name, spec.Namespace)
	service := models.KubeServiceVersionMessageGet(puborrollback.ServiceId, ClusterId, puborrollback.VersionId)
	if service == nil {
		return errors.New("no service")
	}
	serviceport, _ := models.KubeServicePortGet(puborrollback.ServiceId)
	var r api.ResourceRequirements
	//资源分配会遇到无法设置值的问题，故采用json反解析
	j := `{"limits": {"cpu":"%s", "memory": "%sMi"}, "requests": {"cpu":"%s", "memory": "%sMi"}}`
	resource := fmt.Sprintf(j, service.CpuMax, service.MemoryMax, service.CpuNeed, service.MemoryNeed)
	fmt.Println(resource)
	json.Unmarshal([]byte(resource), &r)

	annotations := map[string]string{}
	//if spec.Description != nil {
	//	annotations[DescriptionAnnotationKey] = *spec.Description
	//}
	//labels := getLabelsMap(spec.Labels)
	labels := map[string]string{
		"app": service.Name,
	}
	objectMeta := metaV1.ObjectMeta{
		Annotations: annotations,
		Name:        service.Name,
		Labels:      labels,
	}

	containerSpec := api.Container{
		Name:            service.Name,
		Image:           service.ImageName,
		ImagePullPolicy: api.PullIfNotPresent,
		//Resources: api.ResourceRequirements{
		//	Requests: make(map[api.ResourceName]resource.Quantity),
		//},
		Resources: r,
		//Env:       convertEnvVarsSpec(spec.Variables),
		//LivenessProbe: probe,
	}

	if len(serviceport) != 0 {
		port, _ := strconv.Atoi(serviceport[0].ContainerPort)
		probe := &api.Probe{
			Handler: api.Handler{
				HTTPGet: &api.HTTPGetAction{
					Path: service.Heartbeat,
					Port: intstr.FromInt(port),
					//Scheme: "",
				},
			},
			InitialDelaySeconds: int32(service.RunTime),
			PeriodSeconds:       int32(service.SoketTime),
		}
		containerSpec.LivenessProbe = probe
	}

	if service.Run != "" {
		containerSpec.Command = []string{service.Run}
	}
	//if spec.ContainerCommandArgs != nil {
	//	containerSpec.Args = []string{*spec.ContainerCommandArgs}
	//}

	podSpec := api.PodSpec{
		Containers: []api.Container{containerSpec},
	}
	//if spec.ImagePullSecret != nil {
	//	podSpec.ImagePullSecrets = []api.LocalObjectReference{{Name: *spec.ImagePullSecret}}
	//}

	podTemplate := api.PodTemplateSpec{
		ObjectMeta: objectMeta,
		Spec:       podSpec,
	}

	deployment := &apps.Deployment{
		ObjectMeta: objectMeta,
		Spec: apps.DeploymentSpec{
			Replicas: int32Ptr(int32(service.ServiceNum)),
			Template: podTemplate,
			Selector: &metaV1.LabelSelector{
				// Quoting from https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#selector:
				// In API version apps/v1beta2, .spec.selector and .metadata.labels no longer default to
				// .spec.template.metadata.labels if not set. So they must be set explicitly.
				// Also note that .spec.selector is immutable after creation of the Deployment in apps/v1beta2.
				MatchLabels: labels,
			},
		},
	}
	var err error
	if Pubflag {
		_, err = client.AppsV1beta2().Deployments(puborrollback.Namespace).Create(deployment)

	} else {
		_, err = client.AppsV1beta2().Deployments(puborrollback.Namespace).Update(deployment)
	}
	if err != nil {
		// TODO(bryk): Roll back created resources in case of error.
		return err
	}
	return nil
}

// Converts array of labels to map[string]string
func getLabelsMap(labels []Label) map[string]string {
	result := make(map[string]string)

	for _, label := range labels {
		result[label.Key] = label.Value
	}

	return result
}

func convertEnvVarsSpec(variables []EnvironmentVariable) []api.EnvVar {
	var result []api.EnvVar
	for _, variable := range variables {
		result = append(result, api.EnvVar{Name: variable.Name, Value: variable.Value})
	}
	return result
}
