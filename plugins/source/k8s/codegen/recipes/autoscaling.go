package recipes

import (
	resource "k8s.io/api/autoscaling/v1"
	"k8s.io/client-go/kubernetes"
	resourceType "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
)

func Autoscaling() []*Resource {
	resources := []*Resource{
		{
			SubService:   "hpas",
			Struct:       &resource.HorizontalPodAutoscaler{},
			ResourceFunc: resourceType.HorizontalPodAutoscalersGetter.HorizontalPodAutoscalers,
		},
	}

	for _, resource := range resources {
		resource.Service = "autoscaling"
		resource.ServiceFunc = kubernetes.Interface.AutoscalingV1
	}

	return resources
}
