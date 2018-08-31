package k8s

import (
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/fields"
)

type ServiceDetail struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// InternalEndpoint of all Kubernetes services that have the same label selector as connected Replication
	// Controller. Endpoints is DNS name merged with ports.
	InternalEndpoint EndpointShort `json:"internalEndpoint"`

	// ExternalEndpoints of all Kubernetes services that have the same label selector as connected Replication
	// Controller. Endpoints is external IP address name merged with ports.
	ExternalEndpoints []EndpointShort `json:"externalEndpoints"`

	// List of Endpoint obj. that are endpoints of this Service.
	EndpointList EndpointList `json:"endpointList"`

	// Label selector of the service.
	Selector map[string]string `json:"selector"`

	// Type determines how the service will be exposed.  Valid options: ClusterIP, NodePort, LoadBalancer
	Type v1.ServiceType `json:"type"`

	// ClusterIP is usually assigned by the master. Valid values are None, empty string (""), or
	// a valid IP address. None can be specified for headless services when proxying is not required
	ClusterIP string `json:"clusterIP"`

	// List of events related to this Service
	EventList EventList `json:"eventList"`

	// PodList represents list of pods targeted by same label selector as this service.
	PodList PodList `json:"podList"`

	// Show the value of the SessionAffinity of the Service.
	SessionAffinity v1.ServiceAffinity `json:"sessionAffinity"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

//func GetServiceDetail(client kubernetes.Interface, metricClient metricapi.MetricClient, namespace, name string,
func GetServiceDetail(client kubernetes.Interface, namespace, name string) (*ServiceDetail, error) {

	log.Printf("Getting details of %s service in %s namespace", name, namespace)
	serviceData, err := client.CoreV1().Services(namespace).Get(name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	endpointList, err := GetServiceEndpoints(client, namespace, name)
	/*nonCriticalErrors, criticalError := errors.HandleError(err)
	if criticalError != nil {
		return nil, criticalError
	}*/

	podList, err := GetServicePods(client, namespace, name)
	/*nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	if criticalError != nil {
		return nil, criticalError
	}*/

	eventList, err := GetServiceEvents(client, namespace, name)
	/*nonCriticalErrors, criticalError = errors.AppendError(err, nonCriticalErrors)
	if criticalError != nil {
		return nil, criticalError
	}*/

	service := ToServiceDetail(serviceData, *eventList, *podList, *endpointList)
	return &service, nil
}

// GetServicePods gets list of pods targeted by given label selector in given namespace.
//func GetServicePods(client kubernetes.Interface, metricClient metricapi.MetricClient, namespace,
func GetServicePods(client kubernetes.Interface, namespace, name string) (*PodList, error) {
	podList := PodList{
		Pods:              []Pod{},
		//CumulativeMetrics: []metricapi.Metric{},
	}

	service, err := client.CoreV1().Services(namespace).Get(name, metaV1.GetOptions{})
	if err != nil {
		return &podList, err
	}

	if service.Spec.Selector == nil {
		return &podList, nil
	}

	labelSelector := labels.SelectorFromSet(service.Spec.Selector)
	channels := &ResourceChannels{
		PodList: GetPodListChannelWithOptions(client, namespace,
			metaV1.ListOptions{
				LabelSelector: labelSelector.String(),
				FieldSelector: fields.Everything().String(),
			}),
	}

	apiPodList := <-channels.PodList.List
	if err := <-channels.PodList.Error; err != nil {
		return &podList, err
	}

	events, err := GetPodsEvents(client, namespace, apiPodList.Items)
	/*nonCriticalErrors, criticalError := errors.HandleError(err)
	if criticalError != nil {
		return &podList, criticalError
	}*/

	podList = ToPodList(apiPodList.Items, events)
	return &podList, nil
}

// ToServiceDetail returns api service object based on kubernetes service object
//func ToServiceDetail(service *v1.Service, events EventList, pods PodList, endpointList EndpointList, nonCriticalErrors []error) ServiceDetail {
func ToServiceDetail(service *v1.Service, events EventList, pods PodList, endpointList EndpointList) ServiceDetail {
	return ServiceDetail{
		ObjectMeta:        NewObjectMeta(service.ObjectMeta),
		TypeMeta:          NewTypeMeta(ResourceKindService),
		InternalEndpoint:  GetInternalEndpoint(service.Name, service.Namespace, service.Spec.Ports),
		ExternalEndpoints: GetExternalEndpoints(service),
		EndpointList:      endpointList,
		Selector:          service.Spec.Selector,
		ClusterIP:         service.Spec.ClusterIP,
		Type:              service.Spec.Type,
		EventList:         events,
		PodList:           pods,
		SessionAffinity:   service.Spec.SessionAffinity,
		//Errors:            nonCriticalErrors,
	}
}