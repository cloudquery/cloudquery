package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/logging/v2"
)

var loggingResources = []Resource{
	{
		SubService: "metrics",
		Struct:     &logging.LogMetric{},
	},
	{
		SubService: "sinks",
		Struct:     &logging.LogSink{},
	},
}

func LoggingResources() []Resource {
	var resources []Resource
	resources = append(resources, loggingResources...)

	for i := range resources {
		resources[i].Service = "logging"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].StructName = reflect.TypeOf(resources[i].Struct).Elem().Name()
		if resources[i].Template == "" {
			resources[i].Template = "resource_list_projects"
		}
		if resources[i].SkipFields == nil {
			resources[i].SkipFields = []string{"ServerResponse", "NullFields", "ForceSendFields"}
		}
		resources[i].MockImports = []string{"google.golang.org/api/logging/v2"}
		if resources[i].MockListStruct == "" {
			resources[i].MockListStruct = strcase.ToCamel(resources[i].StructName)
		}
	}

	return resources
}
