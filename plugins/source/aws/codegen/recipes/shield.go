package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ShieldResources() []*Resource {
	resources := []*Resource{

		{
			SubService:          "attacks",
			Struct:              &types.AttackDetail{},
			Description:         "https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_AttackDetail.html",
			SkipFields:          []string{"AttackId"},
			PreResourceResolver: "getAttack",
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
			SubService:  "protection_groups",
			Struct:      &types.ProtectionGroup{},
			Description: "https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_ProtectionGroup.html",
			SkipFields:  []string{"ProtectionGroupArn"},
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
			SubService:  "protections",
			Struct:      &types.Protection{},
			Description: "https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Protection.html",
			SkipFields:  []string{"ProtectionArn"},
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
			SubService:  "subscriptions",
			Struct:      &types.Subscription{},
			Description: "https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Subscription.html",
			SkipFields:  []string{"SubscriptionArn"},
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
