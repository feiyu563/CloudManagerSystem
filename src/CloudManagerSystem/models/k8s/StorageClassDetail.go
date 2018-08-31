package k8s

import (
	storage "k8s.io/api/storage/v1"
	"k8s.io/client-go/kubernetes"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

type StorageClassDetail struct {
	ObjectMeta           ObjectMeta                        `json:"objectMeta"`
	TypeMeta             TypeMeta                          `json:"typeMeta"`
	Provisioner          string                                `json:"provisioner"`
	Parameters           map[string]string                     `json:"parameters"`
	PersistentVolumeList PersistentVolumeList `json:"persistentVolumeList"`
}

func GetStorageClassDetail(client kubernetes.Interface, name string) (*StorageClassDetail, error) {
	log.Printf("Getting details of %s storage class", name)

	storage, err := client.StorageV1().StorageClasses().Get(name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	persistentVolumeList, err := GetStorageClassPersistentVolumes(client, storage.Name)
	storageClass := toStorageClassDetail(storage, persistentVolumeList)
	return &storageClass, err
}

func toStorageClassDetail(storageClass *storage.StorageClass,
	persistentVolumeList *PersistentVolumeList) StorageClassDetail {
	return StorageClassDetail{
		ObjectMeta:           NewObjectMeta(storageClass.ObjectMeta),
		TypeMeta:             NewTypeMeta(ResourceKindStorageClass),
		Provisioner:          storageClass.Provisioner,
		Parameters:           storageClass.Parameters,
		PersistentVolumeList: *persistentVolumeList,
	}
}