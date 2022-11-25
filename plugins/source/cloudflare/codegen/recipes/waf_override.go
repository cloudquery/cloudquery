package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func WAFOverrideResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns: []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
			Multiplex:    "client.ZoneMultiplex",
			DataStruct:   &cloudflare.WAFOverride{},
			PKColumns:    []string{"id"},
			//RenameColumns:    map[string]string{"ur_ls": "urls"}, // TODO
			TableName:        "cloudflare_waf_overrides",
			TableFuncName:    "WAFOverrides",
			Filename:         "waf_overrides.go",
			Service:          "waf_overrides",
			ResolverFuncName: "fetchWAFOverrides",
		},
	}
}
