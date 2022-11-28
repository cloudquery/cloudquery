package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainResources() []*Resource {
	return []*Resource{
		{
			DataStruct:          &domain.Details{},
			PKColumns:           []string{"id"},
			Relations:           []string{"DomainLiveDNS()", "DomainWebRedirections()", "DomainGlueRecords()", "DomainDNSSecKeys()"},
			PreResourceResolver: "getDomain",
			SkipSubserviceName:  true,
		},
		{
			DataStruct: &domain.LiveDNS{},
			PKColumns:  []string{"fqdn"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "fqdn",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("fqdn")`,
				},
			},
		},
		{
			DataStruct: &domain.DNSSECKey{},
			PKColumns:  []string{"fqdn", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "fqdn",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("fqdn")`,
				},
			},
		},
		{
			DataStruct: &domain.GlueRecord{},
			PKColumns:  []string{"fqdn", "name"},
		},
		{
			DataStruct: &domain.WebRedirection{},
			PKColumns:  []string{"fqdn", "host", "type"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "fqdn",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("fqdn")`,
				},
			},
		},
	}
}
