package k8s

import(
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"k8s.io/api/batch/v1beta1"
	"k8s.io/client-go/kubernetes"
	"CloudManagerSystem/models"
	batch2 "k8s.io/api/batch/v1beta1"
)


// CronJobList contains a list of CronJobs in the cluster.
type CronJobList struct {
	ListMeta          int       `json:"total"`
	//CumulativeMetrics []Metric `json:"cumulativeMetrics"`
	Items             []CronJob          `json:"rows"`

	// Basic information about resources status on the list.
	Status ResourceStatus `json:"status"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}

// CronJob is a presentation layer view of Kubernetes Cron Job resource.
type CronJob struct {
	ObjectMeta   ObjectMeta `json:"objectMeta"`
	TypeMeta     TypeMeta   `json:"typeMeta"`
	Schedule     string         `json:"schedule"`
	Suspend      *bool          `json:"suspend"`
	Active       int            `json:"active"`
	LastSchedule *metav1.Time   `json:"lastSchedule"`
}

type CronJobQueryParam struct {
	models.BaseQueryParam
}

// GetCronJobList returns a list of all CronJobs in the cluster.
func GetCronJobList(client kubernetes.Interface, namespace string, dsQuery *CronJobQueryParam) (*CronJobList, error) {
	log.Print("Getting list of all cron jobs in the cluster")

	channels := &ResourceChannels{
		CronJobList: GetCronJobListChannel(client, namespace, 1),
	}

	return GetCronJobListFromChannels(channels, dsQuery)
}

// GetCronJobListFromChannels returns a list of all CronJobs in the cluster reading required resource
// list once from the channels.
func GetCronJobListFromChannels(channels *ResourceChannels, dsQuery *CronJobQueryParam) (*CronJobList, error) {

	cronJobs := <-channels.CronJobList.List
	err := <-channels.CronJobList.Error
	//nonCriticalErrors, criticalError := errors.HandleError(err)
	//if criticalError != nil {
	//	return nil, criticalError
	//}

	cronJobList := toCronJobList(cronJobs.Items ,dsQuery)
	cronJobList.Status = getCronJobStatus(cronJobs)
	cronJobList.Errors = append(cronJobList.Errors,err)
	return cronJobList, nil
}

func toCronJobList(cronJobs []v1beta1.CronJob,  params *CronJobQueryParam) *CronJobList {

	list := &CronJobList{
		Items:    make([]CronJob, 0),
		ListMeta: len(cronJobs), //api.ListMeta{TotalItems: len(cronJobs)},
		//Errors:   nonCriticalErrors,
	}

	//cachedResources := &metricapi.CachedResources{}
	//
	//cronJobCells, metricPromises, filteredTotal := dataselect.GenericDataSelectWithFilterAndMetrics(ToCells(cronJobs),
	//	dsQuery, cachedResources, metricClient)
	//cronJobs = FromCells(cronJobCells)
	list.ListMeta = len(cronJobs) //api.ListMeta{TotalItems: filteredTotal}

	itemsCount := int64(len(cronJobs))
	//分页索引
	startindex := params.Offset
	endindex := params.Offset + int64(params.Limit)
	if endindex > itemsCount {
		endindex = itemsCount
	}

	if startindex > itemsCount {
		list.Items = []CronJob{}
	} else {
		pageList := cronJobs[startindex:endindex]

		for _, cronJob := range pageList {
			list.Items = append(list.Items, toCronJob(&cronJob))
		}
	}




	//cumulativeMetrics, err := metricPromises.GetMetrics()
	//if err != nil {
	//	list.CumulativeMetrics = make([]metricapi.Metric, 0)
	//} else {
	//	list.CumulativeMetrics = cumulativeMetrics
	//}

	return list
}

func toCronJob(cj *v1beta1.CronJob) CronJob {
	return CronJob{
		ObjectMeta:   NewObjectMeta(cj.ObjectMeta),
		TypeMeta:     NewTypeMeta(ResourceKindCronJob),
		Schedule:     cj.Spec.Schedule,
		Suspend:      cj.Spec.Suspend,
		Active:       len(cj.Status.Active),
		LastSchedule: cj.Status.LastScheduleTime,
	}
}


func getCronJobStatus(list *batch2.CronJobList) ResourceStatus {
	info := ResourceStatus{}
	if list == nil {
		return info
	}

	for _, cronJob := range list.Items {
		if cronJob.Spec.Suspend != nil && !(*cronJob.Spec.Suspend) {
			info.Running++
		} else {
			info.Failed++
		}
	}

	return info
}