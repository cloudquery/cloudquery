package resources

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var tailnetCol = codegen.ColumnDefinition{
	Name:     "tailnet",
	Type:     schema.TypeString,
	Resolver: "client.ResolveTailnet",
	Options:  schema.ColumnCreationOptions{PrimaryKey: true},
}
