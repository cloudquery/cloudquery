package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/monitoring/v3"
)

var monitoringResources = []*Resource{
	{
		SubService:    "alert_policies",
		Struct:        &monitoring.AlertPolicy{},
		MockPostFaker: "item.Validity.Details = nil",
	},
}

func MonitoringResources() []*Resource {
	var resources []*Resource
	resources = append(resources, monitoringResources...)

	for _, resource := range resources {
		resource.Service = "monitoring"
		resource.MockImports = []string{"google.golang.org/api/monitoring/v3"}
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf(
			`c.Services.Monitoring.Projects.%s.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()`,
			strcase.ToCamel(resource.SubService))
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
