package recipes

import (
	batchv1 "k8s.io/api/batch/v1"
)

func BatchResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "jobs",
			Struct:     &batchv1.Job{},
		},
		{
			SubService: "cron_jobs",
			Struct:     &batchv1.CronJob{},
			Multiplex:  `client.APIFilterContextMultiplex("/apis/batch/v1/cronjobs")`,
		},
	}

	for _, resource := range resources {
		resource.Service = "batch"
	}

	return resources
}
