package k8s

import (
	"CloudManagerSystem/models"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
	"k8s.io/api/core/v1"
)

type NodeList struct {
	ListMeta int    `json:"total"`
	Items    []Node `json:"rows"`
	//CumulativeMetrics []metricapi.Metric `json:"cumulativeMetrics"`

	// List of non-critical errors, that occurred during resource retrieval.
	//Errors []error `json:"errors"`
}

type NodeQueryParam struct {
	models.BaseQueryParam
}

type Node struct {
	metaV1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Status NodeStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type NodeStatus struct {
	// Capacity represents the total resources of a node.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	// +optional
	Capacity v1.ResourceList `json:"capacity,omitempty" protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName"`
	// Allocatable represents the resources of a node that are available for scheduling.
	// Defaults to Capacity.
	// +optional
	Allocatable v1.ResourceList `json:"allocatable,omitempty" protobuf:"bytes,2,rep,name=allocatable,casttype=ResourceList,castkey=ResourceName"`
	// List of addresses reachable to the node.
	// Queried from cloud provider, if available.
	// More info: https://kubernetes.io/docs/concepts/nodes/node/#addresses
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Addresses []v1.NodeAddress `json:"addresses,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,5,rep,name=addresses"`
}

//func GetNodeList( clusterId string) (*v1.NodeList, error) {
func GetNodeList(clusterId string, params *NodeQueryParam) (*NodeList, error) {

	//clusterId = "1"
	clientset, err := models.GetApiServerHandle(clusterId, false)
	if err != nil {
		fmt.Println(err)
	}
	v1Nodes, err := clientset.CoreV1().Nodes().List(metaV1.ListOptions{})

	nodes := &NodeList{}
	if err == nil {

		nodes.ListMeta = len(v1Nodes.Items)
		itemsCount := int64(nodes.ListMeta)

		//分页索引
		startindex := params.Offset
		endindex := params.Offset + int64(params.Limit)
		if endindex > itemsCount {
			endindex = itemsCount
		}

		if startindex > itemsCount {
			nodes.Items = []Node{}
		} else {
			pageList := v1Nodes.Items[startindex:endindex]
			for i := 0; i < len(pageList); i++ {
				item := pageList[i]
				node := Node{ObjectMeta: item.ObjectMeta, Status: NodeStatus{Capacity: item.Status.Capacity, Allocatable: item.Status.Allocatable, Addresses: item.Status.Addresses}}
				nodes.Items = append(nodes.Items, node)
			}
		}
	}

	return nodes, err

}
