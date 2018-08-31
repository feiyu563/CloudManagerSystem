package k8s

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	restclient "k8s.io/client-go/rest"

	client "k8s.io/client-go/kubernetes"
)

type ResourceVerberQueryParam struct {
	Name      string
	Namespace string
	Kind      string
	PutSpec   *runtime.Unknown
}

// ResourceVerber is responsible for performing generic CRUD operations on all supported resources.
type ResourceVerber interface {
	Put(kind string, namespaceSet bool, namespace string, name string,
		object *runtime.Unknown) error
	Get(kind string, namespaceSet bool, namespace string, name string) (runtime.Object, error)
	Delete(kind string, namespaceSet bool, namespace string, name string) error
}

// RESTClient is an interface for REST operations used in this file.
type RESTClient interface {
	Delete() *restclient.Request
	Put() *restclient.Request
	Get() *restclient.Request
}

func VerberClientHandle(clienthandle client.Interface) ResourceVerber {
	return NewResourceVerber(clienthandle.CoreV1().RESTClient(),
		clienthandle.ExtensionsV1beta1().RESTClient(), clienthandle.AppsV1beta2().RESTClient(),
		clienthandle.BatchV1().RESTClient(), clienthandle.BatchV1beta1().RESTClient(), clienthandle.AutoscalingV1().RESTClient(),
		clienthandle.StorageV1().RESTClient())
}

// NewResourceVerber creates a new resource verber that uses the given client for performing operations.
func NewResourceVerber(client, extensionsClient, appsClient,
batchClient, betaBatchClient, autoscalingClient, storageClient RESTClient) ResourceVerber {
	return &resourceVerber{client, extensionsClient, appsClient,
		batchClient, betaBatchClient, autoscalingClient, storageClient}
}

// resourceVerber is a struct responsible for doing common verb operations on resources, like
// DELETE, PUT, UPDATE.
type resourceVerber struct {
	client            RESTClient
	extensionsClient  RESTClient
	appsClient        RESTClient
	batchClient       RESTClient
	betaBatchClient   RESTClient
	autoscalingClient RESTClient
	storageClient     RESTClient
}

func (verber *resourceVerber) getRESTClientByType(clientType ClientType) RESTClient {
	switch clientType {
	case ClientTypeExtensionClient:
		return verber.extensionsClient
	case ClientTypeAppsClient:
		return verber.appsClient
	case ClientTypeBatchClient:
		return verber.batchClient
	case ClientTypeBetaBatchClient:
		return verber.betaBatchClient
	case ClientTypeAutoscalingClient:
		return verber.autoscalingClient
	case ClientTypeStorageClient:
		return verber.storageClient
	default:
		return verber.client
	}
}

// Get gets the resource of the given kind in the given namespace with the given name.
func (verber *resourceVerber) Get(kind string, namespaceSet bool, namespace string, name string) (runtime.Object, error) {
	resourceSpec, ok := KindToAPIMapping[kind]
	if !ok {
		return nil, fmt.Errorf("Unknown resource kind: %s", kind)
	}

	if namespaceSet != resourceSpec.Namespaced {
		if namespaceSet {
			return nil, fmt.Errorf("Set namespace for not-namespaced resource kind: %s", kind)
		} else {
			return nil, fmt.Errorf("Set no namespace for namespaced resource kind: %s", kind)
		}
	}

	client := verber.getRESTClientByType(resourceSpec.ClientType)
	result := &runtime.Unknown{}
	req := client.Get().Resource(resourceSpec.Resource).Name(name).SetHeader("Accept", "application/json")

	if resourceSpec.Namespaced {
		req.Namespace(namespace)
	}

	err := req.Do().Into(result)
	fmt.Printf("%+v\n",result)
	fmt.Println("###################################",string(result.Raw))
	return result, err
}

// Put puts new resource version of the given kind in the given namespace with the given name.
func (verber *resourceVerber) Put(kind string, namespaceSet bool, namespace string, name string,
	object *runtime.Unknown) error {

	resourceSpec, ok := KindToAPIMapping[kind]
	if !ok {
		return fmt.Errorf("Unknown resource kind: %s", kind)
	}

	if namespaceSet != resourceSpec.Namespaced {
		if namespaceSet {
			return fmt.Errorf("Set namespace for not-namespaced resource kind: %s", kind)
		} else {
			return fmt.Errorf("Set no namespace for namespaced resource kind: %s", kind)
		}
	}

	client := verber.getRESTClientByType(resourceSpec.ClientType)

	req := client.Put().
		Resource(resourceSpec.Resource).
		Name(name).
		SetHeader("Content-Type", "application/json").
		Body([]byte(object.Raw))

	if resourceSpec.Namespaced {
		req.Namespace(namespace)
	}

	return req.Do().Error()
}

// Delete deletes the resource of the given kind in the given namespace with the given name.
func (verber *resourceVerber) Delete(kind string, namespaceSet bool, namespace string, name string) error {
	resourceSpec, ok := KindToAPIMapping[kind]
	if !ok {
		return fmt.Errorf("Unknown resource kind: %s", kind)
	}

	if namespaceSet != resourceSpec.Namespaced {
		if namespaceSet {
			return fmt.Errorf("Set namespace for not-namespaced resource kind: %s", kind)
		} else {
			return fmt.Errorf("Set no namespace for namespaced resource kind: %s", kind)
		}
	}

	client := verber.getRESTClientByType(resourceSpec.ClientType)

	// Do cascade delete by default, as this is what users typically expect.
	defaultPropagationPolicy := v1.DeletePropagationForeground
	defaultDeleteOptions := &v1.DeleteOptions{
		PropagationPolicy: &defaultPropagationPolicy,
	}

	req := client.Delete().Resource(resourceSpec.Resource).Name(name).Body(defaultDeleteOptions)

	if resourceSpec.Namespaced {
		req.Namespace(namespace)
	}

	return req.Do().Error()
}
