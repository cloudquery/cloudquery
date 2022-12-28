package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	orgColumn = codegen.ColumnDefinition{
		Name:        "org",
		Description: "The Github Organization of the resource.",
		Type:        schema.TypeString,
		Resolver:    `client.ResolveOrg`,
		Options:     schema.ColumnCreationOptions{PrimaryKey: true},
	}
)
