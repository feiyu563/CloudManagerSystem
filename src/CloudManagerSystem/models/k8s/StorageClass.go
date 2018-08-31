package k8s

import (
	storage "k8s.io/api/storage/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"CloudManagerSystem/models"
)

type StorageClassList struct {
	ListMeta       int   `json:"total"`
	StorageClasses []StorageClass `json:"rows"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

type StorageClass struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`
	TypeMeta   TypeMeta   `json:"typeMeta"`

	// provisioner is the driver expected to handle this StorageClass.
	// This is an optionally-prefixed name, like a label key.
	// For example: "kubernetes.io/gce-pd" or "kubernetes.io/aws-ebs".
	// This value may not be empty.
	Provisioner string `json:"provisioner"`

	// parameters holds parameters for the provisioner.
	// These values are opaque to the  system and are passed directly
	// to the provisioner.  The only validation done on keys is that they are
	// not empty.  The maximum number of parameters is
	// 512, with a cumulative max size of 256K
	// +optional
	Parameters map[string]string `json:"parameters"`
}
type StorageClassQueryParam struct {
	models.BaseQueryParam
}

// GetStorageClassList returns a list of all storage class objects in the cluster.
func GetStorageClassList(client kubernetes.Interface, dsQuery *StorageClassQueryParam) *StorageClassList {
	log.Print("Getting list of storage classes in the cluster")

	channels := &ResourceChannels{
		StorageClassList: GetStorageClassListChannel(client, 1),
	}

	return GetStorageClassListFromChannels(channels, dsQuery)
}

// GetStorageClassListFromChannels returns a list of all storage class objects in the cluster.
func GetStorageClassListFromChannels(channels *ResourceChannels,
	dsQuery *StorageClassQueryParam) *StorageClassList {
	storageClasses := <-channels.StorageClassList.List
	/*err := <-channels.StorageClassList.Error
	nonCriticalErrors, criticalError := errors.HandleError(err)
	if criticalError != nil {
		return nil, criticalError
	}*/

	return toStorageClassList(storageClasses.Items, dsQuery)
}

func toStorageClassList(storageClasses []storage.StorageClass,params *StorageClassQueryParam) *StorageClassList {

	storageClassList := &StorageClassList{
		StorageClasses: make([]StorageClass, 0),
		ListMeta:      len(storageClasses),
		//Errors:         nonCriticalErrors,
	}

	/*//dashbarod 搜索及分页
	storageClassCells, filteredTotal := dataselect.GenericDataSelectWithFilter(toCells(storageClasses), dsQuery)
	storageClasses = fromCells(storageClassCells)
	storageClassList.ListMeta = api.ListMeta{TotalItems: filteredTotal}*/

	itemsCount := int64(len(storageClasses))
	//分页索引
	startindex := params.Offset
	endindex := params.Offset + int64(params.Limit)
	if endindex > itemsCount {
		endindex = itemsCount
	}

	if startindex > itemsCount {
		storageClassList.StorageClasses = []StorageClass{}
	} else {
		pageList := storageClasses[startindex:endindex]

		for _, item := range pageList {
			storageClassList.StorageClasses = append(storageClassList.StorageClasses, toStorageClass(&item))
		}
	}

	return storageClassList
}


func toStorageClass(storageClass *storage.StorageClass) StorageClass {
	return StorageClass{
		ObjectMeta:  NewObjectMeta(storageClass.ObjectMeta),
		TypeMeta:    NewTypeMeta(ResourceKindStorageClass),
		Provisioner: storageClass.Provisioner,
		Parameters:  storageClass.Parameters,
	}
}