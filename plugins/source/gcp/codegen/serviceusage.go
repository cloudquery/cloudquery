package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/serviceusage/v1"
)

var serviceusageResources = []*Resource{
	{
		SubService: "services",
		Struct:     &serviceusage.GoogleApiServiceusageV1Service{},
		SkipMock:   true,
	},
}

func ServiceusageResources() []*Resource {
	var resources []*Resource
	resources = append(resources, serviceusageResources...)

	for _, resource := range resources {
		resource.Service = "serviceusage"
		resource.MockImports = []string{"google.golang.org/api/serviceusage/v1"}
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf(
			`c.Services.Serviceusage.%s.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()`,
			strcase.ToCamel(resource.SubService),
		)
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
