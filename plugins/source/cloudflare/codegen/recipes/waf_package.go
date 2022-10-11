package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WAFPackageResources() []Resource {
	return []Resource{
		{
			DefaultColumns:   []codegen.ColumnDefinition{AccountIDColumn}, // ZoneIDColumn is already in the response
			Multiplex:        "client.ZoneMultiplex",
			CFStruct:         &cloudflare.WAFPackage{},
			PrimaryKey:       "id",
			Template:         "resource_manual",
			TableName:        "cloudflare_waf_packages",
			TableFuncName:    "WAFPackages",
			Filename:         "waf_packages.go",
			Package:          "waf_packages",
			Relations:        []string{"wafGroups()", "wafRules()"},
			ResolverFuncName: "fetchWAFPackages",
		},
		{
			CFStruct: &cloudflare.WAFGroup{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "waf_package_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_waf_groups",
			TableFuncName:    "wafGroups",
			Filename:         "waf_groups.go",
			Package:          "waf_packages",
			ResolverFuncName: "fetchWAFGroups",
		},
		{
			CFStruct: &cloudflare.WAFRule{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "waf_package_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_waf_rules",
			TableFuncName:    "wafRules",
			Filename:         "waf_rules.go",
			Package:          "waf_packages",
			ResolverFuncName: "fetchWAFRules",
		},
	}
}
