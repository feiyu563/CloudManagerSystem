package k8s

import (
	"k8s.io/api/core/v1"
)

type EndpointList struct {
	ListMeta ListMeta `json:"listMeta"`
	// List of endpoints
	Endpoints []Endpoint `json:"endpoints"`
}

type EndpointShort struct {
	// Hostname, either as a domain name or IP address.
	Host string `json:"host"`

	// List of ports opened for this endpoint on the hostname.
	Ports []ServicePort `json:"ports"`
}
// toEndpointList converts array of api events to endpoint List structure
func toEndpointList(endpoints []v1.Endpoints) *EndpointList {
	endpointList := EndpointList{
		Endpoints: make([]Endpoint, 0),
		ListMeta:  ListMeta{TotalItems: len(endpoints)},
	}

	for _, endpoint := range endpoints {
		for _, subSets := range endpoint.Subsets {
			for _, address := range subSets.Addresses {
				endpointList.Endpoints = append(endpointList.Endpoints, *toEndpoint(address, subSets.Ports, true))
			}
			for _, notReadyAddress := range subSets.NotReadyAddresses {
				endpointList.Endpoints = append(endpointList.Endpoints, *toEndpoint(notReadyAddress, subSets.Ports, false))
			}
		}
	}

	return &endpointList
}

// GetInternalEndpoint returns internal endpoint name for the given service properties, e.g.,
// "my-service.namespace 80/TCP" or "my-service 53/TCP,53/UDP".
func GetInternalEndpoint(serviceName, namespace string, ports [] v1.ServicePort) EndpointShort {
	name := serviceName

	/*if namespace != NamespaceDefault && len(namespace) > 0 && len(serviceName) > 0 {
		bufferName := bytes.NewBufferString(name)
		bufferName.WriteString(".")
		bufferName.WriteString(namespace)
		name = bufferName.String()
	}*/

	return EndpointShort{
		Host:  name,
		Ports: GetServicePorts(ports),
	}
}

// GetExternalEndpoints returns endpoints that are externally reachable for a service.
func GetExternalEndpoints(service *v1.Service) []EndpointShort {
	var externalEndpoints []EndpointShort
	if service.Spec.Type == v1.ServiceTypeLoadBalancer {
		for _, ingress := range service.Status.LoadBalancer.Ingress {
			externalEndpoints = append(externalEndpoints, getExternalEndpoint(ingress, service.Spec.Ports))
		}
	}

	for _, ip := range service.Spec.ExternalIPs {
		externalEndpoints = append(externalEndpoints, EndpointShort{
			Host:  ip,
			Ports: GetServicePorts(service.Spec.Ports),
		})
	}

	return externalEndpoints
}

func GetServicePorts(apiPorts []v1.ServicePort) []ServicePort {
	var ports []ServicePort
	for _, port := range apiPorts {
		ports = append(ports, ServicePort{port.Port, port.Protocol, port.NodePort})
	}
	return ports
}

// Returns external endpoint name for the given service properties.
func getExternalEndpoint(ingress v1.LoadBalancerIngress, ports []v1.ServicePort) EndpointShort {
	var host string
	if ingress.Hostname != "" {
		host = ingress.Hostname
	} else {
		host = ingress.IP
	}
	return EndpointShort{
		Host:  host,
		Ports: GetServicePorts(ports),
	}
}
