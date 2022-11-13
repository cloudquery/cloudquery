package recipes

import (
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

func BatchResources() []*Resource {
	resources := []*Resource{
		{
			SubService:   "jobs",
			Struct:       &batchv1.Job{},
			ResourceFunc: v1.JobsGetter.Jobs,
		},
		{
			SubService:   "cron_jobs",
			Struct:       &batchv1.CronJob{},
			Multiplex:    `client.APIFilterContextMultiplex("/apis/batch/v1/cronjobs")`,
			ResourceFunc: v1.CronJobsGetter.CronJobs,
		},
	}

	for _, resource := range resources {
		resource.Service = "batch"
		resource.ServiceFunc = kubernetes.Interface.BatchV1
		resource.SkipMockFields = []string{"FieldsV1"}
	}

	return resources
}
