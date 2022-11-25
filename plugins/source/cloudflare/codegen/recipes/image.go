package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func ImageResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns: []codegen.ColumnDefinition{AccountIDColumn},
			Multiplex:    "client.AccountMultiplex",
			DataStruct:   &cloudflare.Image{},
			PKColumns:    []string{"id"},
			Service:      "images",
		},
	}
}
