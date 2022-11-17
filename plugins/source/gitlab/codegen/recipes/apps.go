package recipes

import (
	"github.com/xanzy/go-gitlab"
)

func Users() []*Resource {
	resources := []*Resource{
		{
			SubService: "users",
			Struct:     &gitlab.User{},
		},
	}

	for _, resource := range resources {
		resource.Service = "users"
		// resource.ServiceFunc = kubernetes.Interface.AppsV1
	}

	return resources
}
