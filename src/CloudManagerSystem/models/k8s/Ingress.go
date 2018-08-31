package k8s

import (
	"CloudManagerSystem/models"
	"fmt"
	"k8s.io/api/extensions/v1beta1"
)

func CreateIngress(clusterId, namespace string, param *v1beta1.Ingress) (bool, error) {

	clienthandle, err := models.GetApiServerHandle(clusterId, false)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	result := false
	_,err = clienthandle.ExtensionsV1beta1().Ingresses(namespace).Create(param)

	if err == nil {
		result = true
	}

	return result, err
}
