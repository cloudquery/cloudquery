package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func ImageResources() []*Resource {
	return []*Resource{
		{
			ExtraColumns:     []codegen.ColumnDefinition{AccountIDColumn},
			Multiplex:        "client.AccountMultiplex",
			DataStruct:       &cloudflare.Image{},
			PKColumns:        []string{"id"},
			Template:         "resource_manual",
			TableName:        "cloudflare_images",
			TableFuncName:    "Images",
			Filename:         "images.go",
			Service:          "images",
			ResolverFuncName: "fetchImages",
		},
	}
}
