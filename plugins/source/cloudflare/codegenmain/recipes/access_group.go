package recipes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func AccessGroupResources() []Resource {
	return []Resource{
		{
			DefaultColumns:   []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
			Multiplex:        "client.ZoneMultiplex",
			CFStruct:         &cloudflare.AccessGroup{},
			PrimaryKey:       "id",
			Template:         "resource_manual",
			TableName:        "cloudflare_access_groups",
			TableFuncName:    "AccessGroups",
			Filename:         "access_groups.go",
			Package:          "access_groups",
			ResolverFuncName: "fetchAccessGroups",
		},
	}
}
