package codegen

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/dns/v1"
)

var dnsResources = []Resource{
	{
		SubService: "policies",
		Struct:     &dns.Policy{},
	},
	{
		SubService: "managed_zones",
		Struct:     &dns.ManagedZone{},
	},
}

func DnsResources() []Resource {
	var resources []Resource
	resources = append(resources, dnsResources...)

	for i := range resources {
		resources[i].Service = "dns"
		resources[i].DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resources[i].StructName = reflect.TypeOf(resources[i].Struct).Elem().Name()
		if resources[i].Template == "" {
			resources[i].Template = "resource_list"
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
