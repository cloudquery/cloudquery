package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func Hooks() []*Resource {
	const (
		deliveredAt = "DeliveredAt"
	)

	return []*Resource{
		{
			Service:      "hooks",
			SubService:   "hooks",
			Multiplex:    orgMultiplex,
			Struct:       new(github.Hook),
			TableName:    "hooks",
			SkipFields:   skipID,
			ExtraColumns: append(orgColumns, idColumn),
			Relations:    []string{"Deliveries()"},
		},
		{
			Service:    "hooks",
			SubService: "deliveries",
			Multiplex:  "", // we skip multiplexing here as it's a relation
			Struct:     new(github.HookDelivery),
			TableName:  "hook_deliveries",
			SkipFields: append(skipID, deliveredAt, "Request", "Response"),
			ExtraColumns: append(orgColumns, idColumn,
				codegen.ColumnDefinition{
					Name:        "hook_id",
					Type:        schema.TypeInt,
					Resolver:    `client.ResolveParentColumn("ID")`,
					Description: "Hook ID",
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},
				codegen.ColumnDefinition{
					Name:     "request",
					Type:     schema.TypeString,
					Resolver: `resolveRequest`,
				},
				codegen.ColumnDefinition{
					Name:     "response",
					Type:     schema.TypeString,
					Resolver: `resolveResponse`,
				},
				timestampField("delivered_at", deliveredAt)),
		},
	}
}
