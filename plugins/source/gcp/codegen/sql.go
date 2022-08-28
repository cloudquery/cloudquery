package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

var sqlResources = []Resource{
	{
		SubService: "instances",
		Struct:     &sqladmin.DatabaseInstance{},
	},
}

func SqlResources() []Resource {
	var resources []Resource
	resources = append(resources, sqlResources...)

	for i := range resources {
		resources[i].Service = "sql"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].StructName = reflect.TypeOf(resources[i].Struct).Elem().Name()
		if resources[i].Template == "" {
			resources[i].Template = "resource_list"
		}
		if resources[i].SkipFields == nil {
			resources[i].SkipFields = []string{"ServerResponse", "NullFields", "ForceSendFields"}
		}
		resources[i].MockImports = []string{"google.golang.org/api/sqladmin/v1beta4"}
		if resources[i].MockListStruct == "" {
			resources[i].MockListStruct = strcase.ToCamel(resources[i].StructName)
		}
	}

	return resources
}
