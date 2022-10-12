package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func ImageResources() []Resource {
	return []Resource{
		{
			DefaultColumns:   []codegen.ColumnDefinition{AccountIDColumn},
			Multiplex:        "client.AccountMultiplex",
			CFStruct:         &cloudflare.Image{},
			PrimaryKey:       "id",
			Template:         "resource_manual",
			TableName:        "cloudflare_images",
			TableFuncName:    "Images",
			Filename:         "images.go",
			Package:          "images",
			ResolverFuncName: "fetchImages",
		},
	}
}
