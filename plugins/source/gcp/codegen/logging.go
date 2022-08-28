package codegen

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/api/logging/v2"
)

var loggingResources = []*Resource{
	{
		SubService:   "metrics",
		Struct:       &logging.LogMetric{},
		ListFunction: `c.Services.Logging.Projects.Metrics.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()`,
	},
	{
		SubService: "sinks",
		Struct:     &logging.LogSink{},
	},
}

func LoggingResources() []*Resource {
	var resources []*Resource
	resources = append(resources, loggingResources...)

	for _, resource := range resources {
		resource.Service = "logging"
		resource.MockImports = []string{"google.golang.org/api/logging/v2"}
		resource.Template = "resource_list"
		if resource.ListFunction == "" {
			resource.ListFunction = `c.Services.Logging.` + strcase.ToCamel(resource.SubService) + `.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()`
		}
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
