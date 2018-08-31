package k8s

import (
	v1 "k8s.io/api/core/v1"
	"CloudManagerSystem/models"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceQuery struct {
	namespaces []string
}

func (n *NamespaceQuery) ToRequestParam() string {
	if len(n.namespaces) == 1 {
		return n.namespaces[0]
	}
	return v1.NamespaceAll
}

// NewSameNamespaceQuery creates new namespace query that queries single namespace.
func NewSameNamespaceQuery(namespace string) *NamespaceQuery {
	return &NamespaceQuery{[]string{namespace}}
}

// NewNamespaceQuery creates new query for given namespaces.
func NewNamespaceQuery(namespaces []string) *NamespaceQuery {
	return &NamespaceQuery{namespaces}
}

type NamespaceQueryParam struct {
	models.BaseQueryParam
}

type NamespaceList struct {
	ListMeta int `json:"total"`
	// List of services
	Items []v1.Namespace `json:"rows"`
}

func GetNamespaceList(clusterId string, params *NamespaceQueryParam) (*NamespaceList, error) {

	//clusterId = "1"
	clientset, err := models.GetApiServerHandle(clusterId, false)
	if err != nil {
		fmt.Println(err)
	}
	v1Namespaces, err := clientset.CoreV1().Namespaces().List(metaV1.ListOptions{})

	namespaces := &NamespaceList{}
	if err == nil {

		namespaces.ListMeta = len(v1Namespaces.Items)
		itemsCount := int64(namespaces.ListMeta)

		//分页索引
		startindex := params.Offset
		endindex := params.Offset + int64(params.Limit)
		if endindex > itemsCount {
			endindex = itemsCount
		}

		if startindex > itemsCount {
			namespaces.Items = []v1.Namespace{}
		} else {
			pageList := v1Namespaces.Items[startindex:endindex]
			namespaces.Items = pageList
		}
	}

	return namespaces, err
}
