package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/livedns"
)

func LiveDNSResources() []Resource {
	return []Resource{
		{
			DataStruct:   &livedns.Domain{},
			PKColumns:    []string{"fqdn"},
			ExtraColumns: []codegen.ColumnDefinition{SharingIDColumn},
			Relations:    []string{"DomainSnapshots()"},
			TableName:    "gandi_livedns_domains",

			Template:         "resource_manual",
			Package:          "livedns",
			TableFuncName:    "Domains",
			Filename:         "domains.go",
			ResolverFuncName: "fetchDomains",
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
			TableName: "gandi_livedns_snapshots",

			Template:         "resource_manual",
			Package:          "livedns",
			TableFuncName:    "DomainSnapshots",
			Filename:         "domain_snapshots.go",
			ResolverFuncName: "fetchDomainSnapshots",
		},
	}
}
