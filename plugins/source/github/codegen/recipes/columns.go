package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	orgColumns = []codegen.ColumnDefinition{
		{
			Name:        "org",
			Description: "The Github Organization of the resource.",
			Type:        schema.TypeString,
			Resolver:    `client.ResolveOrg`,
			Options:     schema.ColumnCreationOptions{PrimaryKey: true},
		},
	}
	idColumn = pkColumn("id", "ID")
	skipID   = []string{"ID"}
)

func timestampField(name, path string) codegen.ColumnDefinition {
	return codegen.ColumnDefinition{
		Name:     name,
		Type:     schema.TypeTimestamp,
		Resolver: `schema.PathResolver("` + path + `.Time")`,
	}
}

func pkColumn(name, path string) codegen.ColumnDefinition {
	return codegen.ColumnDefinition{
		Name:     name,
		Type:     schema.TypeInt,
		Resolver: `schema.PathResolver("` + path + `")`,
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	}
}
