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
			Relations:        []string{"wafGroups()", "wafRules()"},
			ResolverFuncName: "services.FetchWAFPackages",
		},
		{
			CFStruct: &cloudflare.WAFGroup{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "waf_package_cq_id",
					Type:     schema.TypeUUID,
					Resolver: "schema.ParentIDResolver",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_waf_groups",
			TableFuncName:    "wafGroups",
			Filename:         "waf_packages_waf_groups.go",
			ResolverFuncName: "services.FetchWAFGroups",
		},
		{
			CFStruct: &cloudflare.WAFRule{},
			DefaultColumns: []codegen.ColumnDefinition{
				{
					Name:     "waf_package_cq_id",
					Type:     schema.TypeUUID,
					Resolver: "schema.ParentIDResolver",
				},
			},
			Template:         "resource_manual",
			TableName:        "cloudflare_waf_rules",
			TableFuncName:    "wafRules",
			Filename:         "waf_packages_waf_rules.go",
			ResolverFuncName: "services.FetchWAFRules",
		},
	}
}
