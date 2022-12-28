package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func Hooks() []*Resource {
	return []*Resource{
		{
			TableName:    "hooks",
			Service:      "hooks",
			SubService:   "hooks",
			Struct:       new(github.Hook),
			PKColumns:    []string{"id"},
			ExtraColumns: codegen.ColumnDefinitions{orgColumn},
			Multiplex:    orgMultiplex,
			Relations:    []string{"Deliveries()"},
		},
		{
			TableName:  "hook_deliveries",
			Service:    "hooks",
			SubService: "deliveries",
			Struct:     new(github.HookDelivery),
			PKColumns:  []string{"id"},
			SkipFields: []string{"Request", "Response"},
			ExtraColumns: codegen.ColumnDefinitions{
				orgColumn,
				{
					Name:        "hook_id",
					Type:        schema.TypeInt,
					Resolver:    `client.ResolveParentColumn("ID")`,
					Description: "Hook ID",
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "request",
					Type:     schema.TypeString,
					Resolver: `resolveRequest`,
				},
				{
					Name:     "response",
					Type:     schema.TypeString,
					Resolver: `resolveResponse`,
				},
			},
			Multiplex: "", // we skip multiplexing here as it's a relation
		},
	}
}
