package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DomainResources() []*Resource {
	return []*Resource{
		{
			DataStruct: &vercel.Domain{},
			Service:    "domain",
			Multiplex:  "client.TeamMultiplex",
			PKColumns:  []string{"id"},
			Relations:  []string{"DomainRecords()"},
		},
		{
			DataStruct: &vercel.DomainRecord{},
			Service:    "domain",
			Multiplex:  "client.TeamMultiplex",
			PKColumns:  []string{"domain_name", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "domain_name",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("name")`,
				},
			},
		},
	}
}
