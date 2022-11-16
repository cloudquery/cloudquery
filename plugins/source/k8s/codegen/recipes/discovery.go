package recipes

import (
	resource "k8s.io/api/discovery/v1"
	"k8s.io/client-go/kubernetes"
	resourceType "k8s.io/client-go/kubernetes/typed/discovery/v1"
)

func Discovery() []*Resource {
	resources := []*Resource{
		{
			SubService:   "endpoint_slices",
			Struct:       &resource.EndpointSlice{},
			ResourceFunc: resourceType.EndpointSlicesGetter.EndpointSlices,
		},
	}

	for _, resource := range resources {
		resource.Service = "discovery"
		resource.ServiceFunc = kubernetes.Interface.DiscoveryV1
	}

	return resources
}
