package recipes

import (
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

func Batch() []*Resource {
	resources := []*Resource{
		{
			SubService:   "jobs",
			Struct:       &batchv1.Job{},
			ResourceFunc: v1.JobsGetter.Jobs,
			FakerOverride: `
			r.Spec.Template = corev1.PodTemplateSpec{}`,
			MockImports: []string{`corev1 "k8s.io/api/core/v1"`},
		},
		{
			SubService:   "cron_jobs",
			Struct:       &batchv1.CronJob{},
			Multiplex:    `client.APIFilterContextMultiplex("/apis/batch/v1/cronjobs")`,
			ResourceFunc: v1.CronJobsGetter.CronJobs,
			FakerOverride: `
			r.Spec.JobTemplate = resource.JobTemplateSpec{}
			`,
		},
	}

	for _, resource := range resources {
		resource.Service = "batch"
		resource.ServiceFunc = kubernetes.Interface.BatchV1
	}

	return resources
}
