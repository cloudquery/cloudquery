package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WAFPackageResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns:     []codegen.ColumnDefinition{AccountIDColumn}, // ZoneIDColumn is already in the response
			Multiplex:        "client.ZoneMultiplex",
			DataStruct:       &cloudflare.WAFPackage{},
			PKColumns:        []string{"id"},
			TableName:        "cloudflare_waf_packages",
			TableFuncName:    "WAFPackages",
			Filename:         "waf_packages.go",
			Service:          "waf_packages",
			Relations:        []string{"wafGroups()", "wafRules()"},
			ResolverFuncName: "fetchWAFPackages",
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
			TableName:        "cloudflare_waf_groups",
			TableFuncName:    "wafGroups",
			Filename:         "waf_groups.go",
			Service:          "waf_packages",
			ResolverFuncName: "fetchWAFGroups",
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
			TableName:        "cloudflare_waf_rules",
			TableFuncName:    "wafRules",
			Filename:         "waf_rules.go",
			Service:          "waf_packages",
			ResolverFuncName: "fetchWAFRules",
		},
	}
}
