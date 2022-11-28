package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func AccessGroupResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns: []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
			Multiplex:    "client.ZoneMultiplex",
			DataStruct:   &cloudflare.AccessGroup{},
			PKColumns:    []string{"id"},
			Service:      "access_groups",
		},
	}
}
