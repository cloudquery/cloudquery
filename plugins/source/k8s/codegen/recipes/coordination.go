package recipes

import (
	resource "k8s.io/api/coordination/v1"
	"k8s.io/client-go/kubernetes"
	resourceType "k8s.io/client-go/kubernetes/typed/coordination/v1"
)

func Coordination() []*Resource {
	resources := []*Resource{
		{
			SubService:   "leases",
			Struct:       &resource.Lease{},
			ResourceFunc: resourceType.LeasesGetter.Leases,
		},
	}

	for _, resource := range resources {
		resource.Service = "coordination"
		resource.ServiceFunc = kubernetes.Interface.CoordinationV1
	}

	return resources
}
