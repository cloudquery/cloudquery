package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/iam/v1"
)

var iamResources = []Resource{
	{
		SubService: "roles",
		Struct:     &iam.Role{},
	},
	{
		SubService: "service_accounts",
		Struct:     &iam.ServiceAccount{},
	},
}

func IamResources() []Resource {
	var resources []Resource
	resources = append(resources, iamResources...)

	for i := range resources {
		resources[i].Service = "iam"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].StructName = reflect.TypeOf(resources[i].Struct).Elem().Name()
		if resources[i].Template == "" {
			resources[i].Template = "resource_list_projects"
		}
		if resources[i].SkipFields == nil {
			resources[i].SkipFields = []string{"ServerResponse", "NullFields", "ForceSendFields"}
		}
		resources[i].MockImports = []string{"google.golang.org/api/dns/v1"}
		if resources[i].MockListStruct == "" {
			resources[i].MockListStruct = strcase.ToCamel(resources[i].StructName)
		}
	}

	return resources
}
