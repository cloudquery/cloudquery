package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/logging/v2"
)

var loggingResources = []*Resource{
	{
		SubService:   "metrics",
		Struct:       &logging.LogMetric{},
		ListFunction: `c.Services.Logging.Projects.Metrics.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()`,
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.PathResolver("Name")`,
			},
		},
	},
	{
		SubService: "sinks",
		Struct:     &logging.LogSink{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.PathResolver("Name")`,
			},
		},
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
