package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/cloudbilling/v1"
)

var cloudbillingResources = []Resource{
	{
		SubService: "accounts",
		Struct:     &cloudbilling.BillingAccount{},
	},
}

func CloudBillingResources() []Resource {
	var resources []Resource
	resources = append(resources, cloudbillingResources...)

	for i := range resources {
		resources[i].Service = "cloudbilling"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].StructName = reflect.TypeOf(resources[i].Struct).Elem().Name()
		if resources[i].Template == "" {
			resources[i].Template = "resource_list_global"
		}
		if resources[i].SkipFields == nil {
			resources[i].SkipFields = []string{"ServerResponse", "NullFields", "ForceSendFields"}
		}
		resources[i].MockImports = []string{"google.golang.org/api/cloudbilling/v1"}
		if resources[i].MockListStruct == "" {
			resources[i].MockListStruct = strcase.ToCamel(resources[i].StructName)
		}
	}

	return resources
}
