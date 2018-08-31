package k8s

import api "k8s.io/api/core/v1"

type ServicePort struct {
	// Positive port number.
	Port int32 `json:"port"`

	// Protocol name, e.g., TCP or UDP.
	Protocol api.Protocol `json:"protocol"`

	// The port on each node on which service is exposed.
	NodePort int32 `json:"nodePort"`
}

type EndpointService struct {
	// Hostname, either as a domain name or IP address.
	Host string `json:"host"`

	// List of ports opened for this endpoint on the hostname.
	Ports []ServicePort `json:"ports"`
}