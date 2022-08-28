package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/cloudfunctions/v1"
)

var cloudFunctionsResources = []Resource{
	{
		SubService: "functions",
		Struct:     &cloudfunctions.CloudFunction{},
	},
}

func CloudFunctionsResources() []Resource {
	var resources []Resource
	resources = append(resources, cloudFunctionsResources...)

	for i := range resources {
		resources[i].Service = "cloud_functions"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].StructName = reflect.TypeOf(resources[i].Struct).Elem().Name()
		if resources[i].Template == "" {
			resources[i].Template = "resource_list_projects_locations"
		}
		if resources[i].SkipFields == nil {
			resources[i].SkipFields = []string{"ServerResponse", "NullFields", "ForceSendFields"}
		}
		resources[i].MockImports = []string{"google.golang.org/api/cloudfunctions/v1"}
		if resources[i].MockListStruct == "" {
			resources[i].MockListStruct = strcase.ToCamel(resources[i].StructName)
		}
	}

	return resources
}
