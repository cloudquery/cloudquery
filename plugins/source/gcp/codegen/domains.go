package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	domains "google.golang.org/api/domains/v1beta1"
)

var domainsResources = []Resource{
	{
		SubService: "registrations",
		Struct:     &domains.Registration{},
	},
}

func DomainsResources() []Resource {
	var resources []Resource
	resources = append(resources, domainsResources...)

	for i := range resources {
		resources[i].Service = "domains"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].StructName = reflect.TypeOf(resources[i].Struct).Elem().Name()
		if resources[i].Template == "" {
			resources[i].Template = "resource_list_projects_locations"
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
