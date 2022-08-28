package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/container/v1"
)

var kubernetesResources = []*Resource{
	{
		SubService: "clusters",
		Struct:     &container.Cluster{},
	},
}

func KubernetesResources() []*Resource {
	var resources []*Resource
	resources = append(resources, kubernetesResources...)

	for _, resource := range resources {
		resource.Service = "container"
		resource.ListFunction = fmt.Sprintf(`c.Services.Container.Projects.Locations.%s.List("projects/" + c.ProjectId + "/locations/-").Do()`, strcase.ToCamel(resource.SubService))
		resource.Template = "resource_get"
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
