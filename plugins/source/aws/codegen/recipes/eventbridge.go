package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EventbridgeResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "event_buses",
			Struct:     &types.EventBus{},
			SkipFields: []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveEventbridgeEventBusTags`,
					},
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"EventBusRules()",
			},
		},
		{
			SubService: "event_bus_rules",
			Struct:     &types.Rule{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "event_bus_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveEventbridgeEventBusRuleTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "eventbridge"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("events")`
	}
	return resources
}
