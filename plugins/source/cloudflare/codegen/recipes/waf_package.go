package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WAFPackageResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns: []codegen.ColumnDefinition{AccountIDColumn}, // ZoneIDColumn is already in the response
			Multiplex:    "client.ZoneMultiplex",
			DataStruct:   &cloudflare.WAFPackage{},
			PKColumns:    []string{"id"},
			Service:      "waf_packages",
			Relations:    []string{"WAFGroups()", "WAFRules()"},
		},
		{
			DataStruct: &cloudflare.WAFGroup{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "waf_package_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Service:   "waf_packages",
			TableName: "waf_groups",
		},
		{
			DataStruct: &cloudflare.WAFRule{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "waf_package_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Service:   "waf_packages",
			TableName: "waf_rules",
		},
	}
}
