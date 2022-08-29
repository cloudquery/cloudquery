package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/run/v1"
)

var runResources = []*Resource{
	{
		SubService: "services",
		Struct:     &run.Service{},
	},
}

func RunResources() []*Resource {
	var resources []*Resource
	resources = append(resources, runResources...)

	for _, resource := range resources {
		resource.Service = "run"
		resource.Template = "resource_list"
		resource.MockImports = []string{"google.golang.org/api/run/v1"}
		resource.ListFunction = fmt.Sprintf(`c.Services.Run.Projects.Locations.%s.List("projects/" + c.ProjectId + "/locations/-").Continue(nextPageToken)`, strcase.ToCamel(resource.SubService))
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
