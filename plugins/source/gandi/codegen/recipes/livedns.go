package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/livedns"
)

func LiveDNSResources() []*Resource {
	return []*Resource{
		{
			DataStruct:   &livedns.Domain{},
			PKColumns:    []string{"fqdn"},
			ExtraColumns: []codegen.ColumnDefinition{SharingIDColumn},
			Relations:    []string{"LiveDNSDomainSnapshots()"},
		},
		{
			DataStruct: &livedns.Snapshot{},
			PKColumns:  []string{"fqdn", "id"},
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
