package recipes

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/api/dns/v1"
)

func init() {
	resources := []*Resource{
		{
			SubService:  "policies",
			Struct:      &dns.Policy{},
			NewFunction: dns.NewService,
			PrimaryKeys: []string{"id"},
			Description: "https://cloud.google.com/dns/docs/reference/v1/policies#resource",
		},
		{
			SubService:  "managed_zones",
			Struct:      &dns.ManagedZone{},
			NewFunction: dns.NewManagedZoneOperationsService,
			PrimaryKeys: []string{"id"},
			Description: "https://cloud.google.com/dns/docs/reference/v1/managedZones#resource",
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
