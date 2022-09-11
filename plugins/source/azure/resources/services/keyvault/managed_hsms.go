// Auto generated code - DO NOT EDIT.

package keyvault

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func ManagedHsms() *schema.Table {
	return &schema.Table{
		Name:      "azure_keyvault_managed_hsms",
		Resolver:  fetchKeyVaultManagedHsms,
		Multiplex: client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "initial_admin_object_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Properties.InitialAdminObjectIds"),
			},
			{
				Name:     "hsm_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.HsmURI"),
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
				Name:     "enable_purge_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Properties.EnablePurgeProtection"),
			},
			{
				Name:     "create_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.CreateMode"),
			},
			{
				Name:     "status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.StatusMessage"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.ProvisioningState"),
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
				Name:     "sku",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sku"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

func fetchKeyVaultManagedHsms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().KeyVault.ManagedHsms

	maxResults := int32(100)
	response, err := svc.ListBySubscription(ctx, &maxResults)

	if err != nil {
		return errors.WithStack(err)
	}

	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
