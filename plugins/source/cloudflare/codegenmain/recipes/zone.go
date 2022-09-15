package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func ZoneResources() []Resource {
	return []Resource{
		{
			DefaultColumns:   []codegen.ColumnDefinition{AccountIDColumn},
			Multiplex:        "client.AccountMultiplex",
			CFStruct:         &cloudflare.Zone{},
			PrimaryKey:       "id",
			Template:         "resource_manual",
			TableName:        "cloudflare_zones",
			TableFuncName:    "Zones",
			Filename:         "zones.go",
			Package:          "zones",
			ResolverFuncName: "fetchZones",
		},
	}
}
