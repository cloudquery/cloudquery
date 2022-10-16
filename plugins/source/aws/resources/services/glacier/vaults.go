// Code generated by codegen; DO NOT EDIT.

package glacier

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Vaults() *schema.Table {
	return &schema.Table{
		Name:      "aws_glacier_vaults",
		Resolver:  fetchGlacierVaults,
		Multiplex: client.ServiceAccountRegionMultiplexer("glacier"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlacierVaultTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VaultARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "last_inventory_date",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastInventoryDate"),
			},
			{
				Name:     "number_of_archives",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumberOfArchives"),
			},
			{
				Name:     "size_in_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SizeInBytes"),
			},
			{
				Name:     "vault_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VaultName"),
			},
		},

		Relations: []*schema.Table{
			VaultAccessPolicies(),
			VaultLockPolicies(),
			VaultNotifications(),
		},
	}
}
