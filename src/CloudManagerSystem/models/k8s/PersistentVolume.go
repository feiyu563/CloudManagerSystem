package k8s

import (
	"k8s.io/api/core/v1"
	"log"
	"k8s.io/client-go/kubernetes"
	"CloudManagerSystem/models"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type PersistentVolumeList struct {
	//ListMeta ListMeta       `json:"listMeta"`
	ListMeta int                `json:"total"`
	Items    []PersistentVolume `json:"rows"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

type PersistentVolume struct {
	ObjectMeta    ObjectMeta                       `json:"objectMeta"`
	TypeMeta      TypeMeta                         `json:"typeMeta"`
	Capacity      v1.ResourceList                  `json:"capacity"`
	AccessModes   []v1.PersistentVolumeAccessMode  `json:"accessModes"`
	ReclaimPolicy v1.PersistentVolumeReclaimPolicy `json:"reclaimPolicy"`
	StorageClass  string                           `json:"storageClass"`
	Status        v1.PersistentVolumePhase         `json:"status"`
	Claim         string                           `json:"claim"`
	Reason        string                           `json:"reason"`
}

type PersistentVolumeQueryParam struct {
	models.BaseQueryParam
}

func GetPersistentVolumeList(client kubernetes.Interface, param *PersistentVolumeQueryParam) (*PersistentVolumeList) {
	log.Print("Getting list persistent volumes")
	channels := &ResourceChannels{
		PersistentVolumeList: GetPersistentVolumeListChannel(client, 1),
	}

	return GetPersistentVolumeListFromChannels(channels, param)
}

//func GetPersistentVolumeListFromChannels(channels *common.ResourceChannels, dsQuery *dataselect.DataSelectQuery) (*PersistentVolumeList, error) {
func GetPersistentVolumeListFromChannels(channels *ResourceChannels, param *PersistentVolumeQueryParam) *PersistentVolumeList {
	persistentVolumes := <-channels.PersistentVolumeList.List
	//err := <-channels.PersistentVolumeList.Error

	//nonCriticalErrors, criticalError := errors.HandleError(err)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	return toPersistentVolumeList(persistentVolumes.Items, param)
}

//func toPersistentVolumeList(persistentVolumes []v1.PersistentVolume, nonCriticalErrors []error,
func toPersistentVolumeList(persistentVolumes []v1.PersistentVolume, params *PersistentVolumeQueryParam) *PersistentVolumeList {

	result := &PersistentVolumeList{
		Items:    make([]PersistentVolume, 0),
		ListMeta: len(persistentVolumes),
		//Errors:   nonCriticalErrors,
	}

	//pvCells, filteredTotal := dataselect.GenericDataSelectWithFilter(toCells(persistentVolumes), dsQuery)
	//persistentVolumes = fromCells(pvCells)
	//result.ListMeta = api.ListMeta{TotalItems: filteredTotal}

	itemsCount := int64(result.ListMeta)
	//分页索引
	startindex := params.Offset
	endindex := params.Offset + int64(params.Limit)
	if endindex > itemsCount {
		endindex = itemsCount
	}

	if startindex > itemsCount {
		result.Items = []PersistentVolume{}
	} else {
		pageList := persistentVolumes[startindex:endindex]

		for _, item := range pageList {
			result.Items = append(result.Items,
				PersistentVolume{
					ObjectMeta:    NewObjectMeta(item.ObjectMeta),
					TypeMeta:      NewTypeMeta(ResourceKindPersistentVolume),
					Capacity:      item.Spec.Capacity,
					AccessModes:   item.Spec.AccessModes,
					ReclaimPolicy: item.Spec.PersistentVolumeReclaimPolicy,
					StorageClass:  item.Spec.StorageClassName,
					Status:        item.Status.Phase,
					Claim:         getPersistentVolumeClaim(&item),
					Reason:        item.Status.Reason,
				})
		}
	}
	return result
}

func getPersistentVolumeClaim(pv *v1.PersistentVolume) string {
	var claim string

	if pv.Spec.ClaimRef != nil {
		claim = pv.Spec.ClaimRef.Namespace + "/" + pv.Spec.ClaimRef.Name
	}

	return claim
}

func GetStorageClassPersistentVolumes(client kubernetes.Interface, storageClassName string) (*PersistentVolumeList, error) {

	storageClass, err := client.StorageV1().StorageClasses().Get(storageClassName, metaV1.GetOptions{})

	if err != nil {
		return nil, err
	}

	channels := &ResourceChannels{
		PersistentVolumeList: GetPersistentVolumeListChannel(
			client, 1),
	}

	persistentVolumeList := <-channels.PersistentVolumeList.List

	//err = <-channels.PersistentVolumeList.Error
	//nonCriticalErrors, criticalError := errors.HandleError(err)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	storagePersistentVolumes := make([]v1.PersistentVolume, 0)
	for _, pv := range persistentVolumeList.Items {
		if strings.Compare(pv.Spec.StorageClassName, storageClass.Name) == 0 {
			storagePersistentVolumes = append(storagePersistentVolumes, pv)
		}
	}

	log.Printf("Found %d persistentvolumes related to %s storageclass",
		len(storagePersistentVolumes), storageClassName)

	pars := PersistentVolumeQueryParam{models.BaseQueryParam{Offset:int64(0),Limit:int(4294967295)}}

	return toPersistentVolumeList(storagePersistentVolumes, &pars), nil
}
