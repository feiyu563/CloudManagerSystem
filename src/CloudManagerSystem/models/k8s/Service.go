package k8s

import (
	"CloudManagerSystem/models"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

type ServiceQueryParam struct {
	models.BaseQueryParam
}

type ServiceList struct {
	ListMeta int `json:"total"`
	// List of services
	Items []v1.Service `json:"rows" protobuf:"bytes,2,rep,name=items"`
}

func GetServiceList(clusterId, namespace string, params *ServiceQueryParam) (*ServiceList, error) {

	clientset, err := models.GetApiServerHandle(clusterId, false)
	if err != nil {
		fmt.Println(err)
	}
	v1Services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
	//return v1Services, err

	services := &ServiceList{}
	if err == nil {
		services.ListMeta = len(v1Services.Items)
		itemsCount := int64(services.ListMeta)
		//分页索引
		startindex := params.Offset
		endindex := params.Offset + int64(params.Limit)
		if endindex > itemsCount {
			endindex = itemsCount
		}

		if startindex > itemsCount {
			services.Items = []v1.Service{}
		} else {
			pageList := v1Services.Items[startindex:endindex]
			services.Items = pageList //v1Services.Items
		}
	}

	return services, err
}

func CreateService(clusterId, namespace string, paras []*v1.Service) (bool, error) {

	clienthandle, err := models.GetApiServerHandle(clusterId, false)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	svcCreateChannel := GetServiceCreateChannel(clienthandle, namespace, paras)

	result := <-svcCreateChannel.Result
	if !result {
		err = <-svcCreateChannel.Eorror
	}
	return result, err
}
