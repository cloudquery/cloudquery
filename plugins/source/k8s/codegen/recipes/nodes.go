package recipes

import (
	resource "k8s.io/api/node/v1"
	"k8s.io/client-go/kubernetes"
	resourceType "k8s.io/client-go/kubernetes/typed/node/v1"
)

func Nodes() []*Resource {
	resources := []*Resource{
		{
			SubService:     "runtime_classes",
			Struct:         &resource.RuntimeClass{},
			ResourceFunc:   resourceType.RuntimeClassesGetter.RuntimeClasses,
			GlobalResource: true,
		},
	}

	for _, resource := range resources {
		resource.Service = "nodes"
		resource.ServiceFunc = kubernetes.Interface.NodeV1
	}

	return resources
}
