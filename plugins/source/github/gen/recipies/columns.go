package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var defaultOrgColumns = []codegen.ColumnDefinition{
	{
		Name:        "org",
		Description: "The Github Organization of the resource.",
		Type:        schema.TypeString,
		Resolver:    `client.ResolveOrg`,
		Options:     schema.ColumnCreationOptions{PrimaryKey: true},
	},
}

func timestampField(name, path string) codegen.ColumnDefinition {
	return codegen.ColumnDefinition{
		Name:     name,
		Type:     schema.TypeTimestamp,
		Resolver: `schema.PathResolver("` + path + `.Time")`,
	}
}
