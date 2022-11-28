package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func WorkerRouteResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns: []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
			Multiplex:    "client.ZoneMultiplex",
			DataStruct:   &cloudflare.WorkerRoute{},
			PKColumns:    []string{"id"},
			Service:      "worker_routes",
		},
	}
}
