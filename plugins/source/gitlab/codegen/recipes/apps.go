package recipes

import (
	"github.com/xanzy/go-gitlab"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func Users() []*Resource {
	resources := []*Resource{
		{
			SubService:   "users",
			Struct:       &gitlab.User{},
			ResourceFunc: v1.DaemonSetsGetter.DaemonSets,
		},
	}

	for _, resource := range resources {
		resource.Service = "users"
		// resource.ServiceFunc = kubernetes.Interface.AppsV1
	}

	return resources
}
