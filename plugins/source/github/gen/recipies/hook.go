package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
)

func Hook() []*Resource {
	return []*Resource{
		{
			Service:      "hooks",
			SubService:   "hooks",
			Struct:       new(github.Hook),
			TableName:    "hooks",
			ExtraColumns: orgColumns,
			Relations:    []string{"Deliveries()"},
		},
		{
			Service:    "hooks",
			SubService: "deliveries",
			Struct:     new(github.HookDelivery),
			TableName:  "hook_deliveries",
			SkipFields: []string{"DeliveredAt"},
			ExtraColumns: append(orgColumns,
				codegen.ColumnDefinition{
					Name:        "hook_id",
					Type:        schema.TypeString,
					Resolver:    `client.ResolveParentColumn("ID")`,
					Description: "Hook ID",
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},
				timestampField("delivered_at", "DeliveredAt")),
		},
	}
}
