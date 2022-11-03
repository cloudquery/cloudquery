package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func KMSResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "aliases",
			Struct:      &types.AliasListEntry{},
			Description: "https://docs.aws.amazon.com/kms/latest/APIReference/API_AliasListEntry.html",
			SkipFields:  []string{"AliasArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
						Resolver: `schema.PathResolver("AliasArn")`,
					},
				}...),
		}, {
			SubService:          "keys",
			Struct:              &types.KeyMetadata{},
			Description:         "https://docs.aws.amazon.com/kms/latest/APIReference/API_KeyMetadata.html",
			PreResourceResolver: "getKey",
			SkipFields:          []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rotation_enabled",
						Type:     schema.TypeBool,
						Resolver: `resolveKeysRotationEnabled`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveKeysTags`,
					},
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:          "replica_keys",
						Type:          schema.TypeJSON,
						Resolver:      `resolveKeysReplicaKeys`,
						IgnoreInTests: true,
					},
				}...),
			Relations: []string{"KeyGrants()"},
		}, {
			SubService:  "key_grants",
			Struct:      &types.GrantListEntry{},
			Description: "https://docs.aws.amazon.com/kms/latest/APIReference/API_GrantListEntry.html",
			SkipFields:  []string{"GrantId"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "key_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "grant_id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("GrantId")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "kms"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("kms")`
	}
	return resources
}
