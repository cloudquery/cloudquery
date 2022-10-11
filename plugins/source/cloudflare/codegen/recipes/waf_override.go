package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func WAFOverrideResources() []Resource {
	return []Resource{
		{
			DefaultColumns:   []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
			Multiplex:        "client.ZoneMultiplex",
			CFStruct:         &cloudflare.WAFOverride{},
			PrimaryKey:       "id",
			Template:         "resource_manual",
			RenameColumns:    map[string]string{"ur_ls": "urls"},
			TableName:        "cloudflare_waf_overrides",
			TableFuncName:    "WAFOverrides",
			Filename:         "waf_overrides.go",
			Package:          "waf_overrides",
			ResolverFuncName: "fetchWAFOverrides",
		},
	}
}
