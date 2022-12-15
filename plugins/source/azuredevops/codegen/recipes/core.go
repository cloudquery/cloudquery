package recipes

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/core"
)

func init() {
	resources := []*Resource{
		{
			SubService: "projects",
			Struct:     &core.TeamProjectReference{},
		},
	}

	for _, resource := range resources {
		resource.Service = "core"
	}

	Resources = append(Resources, resources...)
}
