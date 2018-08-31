package k8s

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

type Endpoint struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// Hostname, either as a domain name or IP address.
	Host string `json:"host"`

	// Name of the node the endpoint is located
	NodeName *string `json:"nodeName"`

	// Status of the endpoint
	Ready bool `json:"ready"`

	// Array of endpoint ports
	Ports []v1.EndpointPort `json:"ports"`
}

// GetServiceEndpoints gets list of endpoints targeted by given label selector in given namespace.
func GetServiceEndpoints(client kubernetes.Interface, namespace, name string) (*EndpointList, error) {
	endpointList := &EndpointList{
		Endpoints: make([]Endpoint, 0),
		ListMeta:  ListMeta{TotalItems: 0},
	}

	serviceEndpoints, err := GetEndpoints(client, namespace, name)
	if err != nil {
		return endpointList, err
	}

	endpointList = toEndpointList(serviceEndpoints)
	log.Printf("Found %d endpoints related to %s service in %s namespace", len(endpointList.Endpoints), name, namespace)
	return endpointList, nil
}

// GetEndpoints gets endpoints associated to resource with given name.
func GetEndpoints(client kubernetes.Interface, namespace, name string) ([]v1.Endpoints, error) {
	fieldSelector, err := fields.ParseSelector("metadata.name" + "=" + name)
	if err != nil {
		return nil, err
	}

	channels := &ResourceChannels{
		EndpointList: GetEndpointListChannelWithOptions(client,
			NewSameNamespaceQuery(namespace),
			metaV1.ListOptions{
				LabelSelector: labels.Everything().String(),
				FieldSelector: fieldSelector.String(),
			},
			1),
	}

	endpointList := <-channels.EndpointList.List
	if err := <-channels.EndpointList.Error; err != nil {
		return nil, err
	}

	return endpointList.Items, nil
}

// toEndpoint converts endpoint api Endpoint to Endpoint model object.
func toEndpoint(address v1.EndpointAddress, ports []v1.EndpointPort, ready bool) *Endpoint {
	return &Endpoint{
		TypeMeta: NewTypeMeta(ResourceKindEndpoint),
		Host:     address.IP,
		Ports:    ports,
		Ready:    ready,
		NodeName: address.NodeName,
	}
}