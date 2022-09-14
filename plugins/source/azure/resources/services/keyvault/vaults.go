// Auto generated code - DO NOT EDIT.

package keyvault

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Vaults() *schema.Table {
	return &schema.Table{
		Name:      "azure_keyvault_vaults",
		Resolver:  fetchKeyVaultVaults,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.Sku"),
			},
			{
				Name:     "access_policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.AccessPolicies"),
			},
			{
				Name:     "vault_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.VaultURI"),
			},
			{
				Name:     "enabled_for_deployment",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnabledForDeployment"),
			},
			{
				Name:     "enabled_for_disk_encryption",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnabledForDiskEncryption"),
			},
			{
				Name:     "enabled_for_template_deployment",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnabledForTemplateDeployment"),
			},
			{
				Name:     "enable_soft_delete",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnableSoftDelete"),
			},
			{
				Name:     "soft_delete_retention_in_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Properties.SoftDeleteRetentionInDays"),
			},
			{
				Name:     "enable_rbac_authorization",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnableRbacAuthorization"),
			},
			{
				Name:     "create_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.CreateMode"),
			},
			{
				Name:     "enable_purge_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnablePurgeProtection"),
			},
			{
				Name:     "network_acls",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.NetworkAcls"),
			},
			{
				Name:     "private_endpoint_connections",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties.PrivateEndpointConnections"),
			},
		},

		Relations: []*schema.Table{
			keys(),
		},
	}
}

func fetchKeyVaultVaults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().KeyVault.Vaults

	maxResults := int32(1000)
	response, err := svc.ListBySubscription(ctx, &maxResults)

	if err != nil {
		return err
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}

	return nil
}
