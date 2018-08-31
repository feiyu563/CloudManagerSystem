package k8s

import (
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

type PersistentVolumeDetail struct {
	ObjectMeta             ObjectMeta                   `json:"objectMeta"`
	TypeMeta               TypeMeta                     `json:"typeMeta"`
	Status                 v1.PersistentVolumePhase         `json:"status"`
	Claim                  string                           `json:"claim"`
	ReclaimPolicy          v1.PersistentVolumeReclaimPolicy `json:"reclaimPolicy"`
	AccessModes            []v1.PersistentVolumeAccessMode  `json:"accessModes"`
	StorageClass           string                           `json:"storageClass"`
	Capacity               v1.ResourceList                  `json:"capacity"`
	Message                string                           `json:"message"`
	PersistentVolumeSource v1.PersistentVolumeSource        `json:"persistentVolumeSource"`
	Reason                 string                           `json:"reason"`
}


// GetPersistentVolumeDetail returns detailed information about a persistent volume
func GetPersistentVolumeDetail(client kubernetes.Interface, name string) (*PersistentVolumeDetail, error) {
	log.Printf("Getting details of %s persistent volume", name)

	rawPersistentVolume, err := client.CoreV1().PersistentVolumes().Get(name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return getPersistentVolumeDetail(rawPersistentVolume), nil
}

func getPersistentVolumeDetail(persistentVolume *v1.PersistentVolume) *PersistentVolumeDetail {
	return &PersistentVolumeDetail{
		ObjectMeta:             NewObjectMeta(persistentVolume.ObjectMeta),
		TypeMeta:               NewTypeMeta(ResourceKindPersistentVolume),
		Status:                 persistentVolume.Status.Phase,
		Claim:                  getPersistentVolumeClaim(persistentVolume),
		ReclaimPolicy:          persistentVolume.Spec.PersistentVolumeReclaimPolicy,
		AccessModes:            persistentVolume.Spec.AccessModes,
		StorageClass:           persistentVolume.Spec.StorageClassName,
		Capacity:               persistentVolume.Spec.Capacity,
		Message:                persistentVolume.Status.Message,
		PersistentVolumeSource: persistentVolume.Spec.PersistentVolumeSource,
		Reason:                 persistentVolume.Status.Reason,
	}
}
