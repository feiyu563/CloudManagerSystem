package k8s

import (
	autoscaling "k8s.io/api/autoscaling/v1"
	k8sClient "k8s.io/client-go/kubernetes"
)

type HorizontalPodAutoscalerList struct {
	ListMeta ListMeta `json:"listMeta"`

	// Unordered list of Horizontal Pod Autoscalers.
	HorizontalPodAutoscalers []HorizontalPodAutoscaler `json:"horizontalpodautoscalers"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

type ScaleTargetRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}
// HorizontalPodAutoscaler (aka. Horizontal Pod Autoscaler)
type HorizontalPodAutoscaler struct {
	ObjectMeta                      ObjectMeta `json:"objectMeta"`
	TypeMeta                        TypeMeta   `json:"typeMeta"`
	ScaleTargetRef                  ScaleTargetRef `json:"scaleTargetRef"`
	MinReplicas                     *int32         `json:"minReplicas"`
	MaxReplicas                     int32          `json:"maxReplicas"`
	CurrentCPUUtilizationPercentage *int32         `json:"currentCPUUtilizationPercentage"`
	TargetCPUUtilizationPercentage  *int32         `json:"targetCPUUtilizationPercentage"`
}



func GetHorizontalPodAutoscalerListForResource(client k8sClient.Interface, namespace, kind, name string) (*HorizontalPodAutoscalerList, error) {
	//nsQuery := NewSameNamespaceQuery(namespace)
	channel := GetHorizontalPodAutoscalerListChannel(client, namespace)
	hpaList := <-channel.List
	//err := <-channel.Error

	//nonCriticalErrors, criticalError := errors.HandleError(err)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	filteredHpaList := make([]autoscaling.HorizontalPodAutoscaler, 0)
	for _, hpa := range hpaList.Items {
		if hpa.Spec.ScaleTargetRef.Kind == kind && hpa.Spec.ScaleTargetRef.Name == name {
			filteredHpaList = append(filteredHpaList, hpa)
		}
	}

	return toHorizontalPodAutoscalerList(filteredHpaList, nil), nil
}

func toHorizontalPodAutoscalerList(hpas []autoscaling.HorizontalPodAutoscaler, nonCriticalErrors []error) *HorizontalPodAutoscalerList {
	hpaList := &HorizontalPodAutoscalerList{
		HorizontalPodAutoscalers: make([]HorizontalPodAutoscaler, 0),
		ListMeta:                 ListMeta{TotalItems: len(hpas)},
		Errors:                   nonCriticalErrors,
	}

	for _, hpa := range hpas {
		horizontalPodAutoscaler := toHorizontalPodAutoScaler(&hpa)
		hpaList.HorizontalPodAutoscalers = append(hpaList.HorizontalPodAutoscalers, horizontalPodAutoscaler)
	}
	return hpaList
}

func toHorizontalPodAutoScaler(hpa *autoscaling.HorizontalPodAutoscaler) HorizontalPodAutoscaler {
	return HorizontalPodAutoscaler{
		ObjectMeta: NewObjectMeta(hpa.ObjectMeta),
		TypeMeta:   NewTypeMeta(ResourceKindHorizontalPodAutoscaler),
		ScaleTargetRef: ScaleTargetRef{
			Kind: hpa.Spec.ScaleTargetRef.Kind,
			Name: hpa.Spec.ScaleTargetRef.Name,
		},
		MinReplicas:                     hpa.Spec.MinReplicas,
		MaxReplicas:                     hpa.Spec.MaxReplicas,
		CurrentCPUUtilizationPercentage: hpa.Status.CurrentCPUUtilizationPercentage,
		TargetCPUUtilizationPercentage:  hpa.Spec.TargetCPUUtilizationPercentage,
	}

}
