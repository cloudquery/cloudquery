package keyvault

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func KeyvaultManagedHSM() *schema.Table {
	return &schema.Table{
		Name:          "azure_keyvault_managed_hsm",
		Description:   "Managed HSM resource information with extended details.",
		Resolver:      fetchKeyvaultManagedHSM,
		Multiplex:     client.SubscriptionMultiplex,
		DeleteFilter:  client.DeleteSubscriptionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "tenant_id",
				Description: "The Azure Active Directory tenant ID that should be used for authenticating requests to the managed HSM pool.",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Properties.TenantID"),
			},
			{
				Name:        "initial_admin_object_ids",
				Description: "Array of initial administrators object ids for this managed hsm pool.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Properties.InitialAdminObjectIds"),
			},
			{
				Name:        "hsm_uri",
				Description: "The URI of the managed hsm pool for performing operations on keys.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.HsmURI"),
			},
			{
				Name:        "enable_soft_delete",
				Description: "Property to specify whether the 'soft delete' functionality is enabled for this managed HSM pool",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnableSoftDelete"),
			},
			{
				Name:        "soft_delete_retention_in_days",
				Description: "Soft delete data retention days.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Properties.SoftDeleteRetentionInDays"),
			},
			{
				Name:        "enable_purge_protection",
				Description: "Property specifying whether protection against purge is enabled for this managed HSM pool.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnablePurgeProtection"),
			},
			{
				Name:        "create_mode",
				Description: "The create mode to indicate whether the resource is being created or is being recovered from a deleted resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.CreateMode"),
			},
			{
				Name:        "status_message",
				Description: "Resource Status Message.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.StatusMessage"),
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:        "id",
				Description: "The Azure Resource Manager resource ID for the managed HSM Pool.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the managed HSM Pool.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The resource type of the managed HSM Pool.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "The supported Azure location where the managed HSM Pool should be created.",
				Type:        schema.TypeString,
			},
			{
				Name:        "sku_family",
				Description: "SKU Family of the managed HSM Pool.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Family"),
			},
			{
				Name:        "sku_name",
				Description: "SKU of the managed HSM Pool.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchKeyvaultManagedHSM(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().KeyVault.ManagedHSM
	maxResults := int32(100)
	response, err := svc.ListBySubscription(ctx, &maxResults)
	if err != nil {
		return diag.WrapError(err)
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
