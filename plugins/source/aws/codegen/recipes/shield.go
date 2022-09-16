package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ShieldResources() []*Resource {
	resources := []*Resource{

		{
			SubService: "attacks",
			Struct:     &types.AttackDetail{},
			SkipFields: []string{"AttackId"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:        "id",
						Description: "The unique identifier (ID) of the attack",
						Type:        schema.TypeString,
						Resolver:    `schema.PathResolver("AttackId")`,
						Options:     schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},

		{
			SubService: "protection_groups",
			Struct:     &types.ProtectionGroup{},
			SkipFields: []string{"ProtectionGroupArn"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ProtectionGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveShieldProtectionGroupTags`,
					},
				}...),
		},

		{
			SubService: "protections",
			Struct:     &types.Protection{},
			SkipFields: []string{"ProtectionArn"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ProtectionArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveShieldProtectionTags`,
					},
				}...),
		},

		{
			SubService: "subscriptions",
			Struct:     &types.Subscription{},
			SkipFields: []string{"SubscriptionArn"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("SubscriptionArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "shield"
		r.Multiplex = "client.AccountMultiplex"
	}
	return resources
}
