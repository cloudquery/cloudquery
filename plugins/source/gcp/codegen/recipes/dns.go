package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/dns/v1"
)



func init() {
	resources := []*Resource{
		{
			SubService:   "policies",
			Struct:       &dns.Policy{},
			NewFunction:  dns.NewService,
			ListFunction: (&dns.PoliciesService{}).List,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "id",
					Type:     schema.TypeInt,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("Id")`,
				},
			},
		},
		{
			SubService:   "managed_zones",
			Struct:       &dns.ManagedZone{},
			NewFunction:  dns.NewManagedZoneOperationsService,
			ListFunction: (&dns.ManagedZonesService{}).List,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "id",
					Type:     schema.TypeInt,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("Id")`,
				},
			},
		},
	}

	for _, resource := range resources {
		resource.Service = "dns"
		resource.SkipFetch = true
		resource.MockImports = []string{"google.golang.org/api/dns/v1"}
		resource.Template = "newapi_list"
		resource.MockTemplate = "resource_list_mock"
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	Resources = append(Resources, resources...)
}
