package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GlacierResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "vaults",
			Struct:     &types.DescribeVaultOutput{},
			SkipFields: []string{"VaultARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveGlacierVaultTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("VaultARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"VaultAccessPolicies()",
				"VaultLockPolicies()",
				"VaultNotifications()",
			},
		},
		{
			SubService: "vault_access_policies",
			Struct:     &types.VaultAccessPolicy{},
			SkipFields: []string{"Policy"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "vault_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "policy",
						Type:     schema.TypeJSON,
						Resolver: `client.MarshaledJsonResolver("Policy")`,
					},
				}...),
		},
		{
			SubService: "vault_lock_policies",
			Struct:     &types.VaultLockPolicy{},
			SkipFields: []string{"Policy"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "vault_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "policy",
						Type:     schema.TypeJSON,
						Resolver: `client.MarshaledJsonResolver("Policy")`,
					},
				}...),
		},
		{
			SubService: "vault_notifications",
			Struct:     &types.VaultNotificationConfig{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "vault_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:   "data_retrieval_policies",
			Struct:       &types.DataRetrievalPolicy{},
			ExtraColumns: defaultRegionalColumnsPK,
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "glacier"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("glacier")`
	}
	return resources
}
