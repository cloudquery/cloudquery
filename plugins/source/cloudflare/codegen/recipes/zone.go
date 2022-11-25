package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func ZoneResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns:     []codegen.ColumnDefinition{AccountIDColumn},
			Multiplex:        "client.AccountMultiplex",
			DataStruct:       &cloudflare.Zone{},
			PKColumns:        []string{"id"},
			TableName:        "cloudflare_zones",
			TableFuncName:    "Zones",
			Filename:         "zones.go",
			Service:          "zones",
			ResolverFuncName: "fetchZones",
		},
	}
}
