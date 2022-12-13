package recipes

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/core"
)

func Core() []*Resource {
	resources := []*Resource{
		{
			SubService: "projects",
			Struct:     core.TeamProjectReference{},
		},
	}

	for _, resource := range resources {
		resource.Service = "core"
	}

	return resources
}
