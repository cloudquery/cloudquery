package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func KeyVaultVaults() *schema.Table {
	return &schema.Table{
		Name:         "azure_keyvault_vaults",
		Resolver:     fetchKeyvaultVaults,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name:     "tenant_id",
				Type:     schema.TypeUUID,
				Resolver: schema.PathResolver("Properties.TenantID"),
			},
			{
				Name:     "sku_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Sku.Family"),
			},
			{
				Name:     "sku_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.Sku.Name"),
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
				Name:     "network_acls_bypass",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.NetworkAcls.Bypass"),
			},
			{
				Name:     "network_acls_default_action",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Properties.NetworkAcls.DefaultAction"),
			},
			{
				Name:     "network_acls_ip_rules",
				Type:     schema.TypeStringArray,
				Resolver: resolveKeyvaultVaultNetworkAclsIPRules,
			},
			{
				Name:     "network_acls_virtual_network_rules",
				Type:     schema.TypeStringArray,
				Resolver: resolveKeyvaultVaultNetworkAclsVirtualNetworkRules,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_keyvault_vault_access_policies",
				Resolver: fetchKeyvaultVaultAccessPolicies,
				Columns: []schema.Column{
					{
						Name:     "vault_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "tenant_id",
						Type:     schema.TypeUUID,
						Resolver: schema.PathResolver("TenantID"),
					},
					{
						Name:     "object_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ObjectID"),
					},
					{
						Name:     "application_id",
						Type:     schema.TypeUUID,
						Resolver: schema.PathResolver("ApplicationID"),
					},
					{
						Name:     "permissions_keys",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Permissions.Keys"),
					},
					{
						Name:     "permissions_secrets",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Permissions.Secrets"),
					},
					{
						Name:     "permissions_certificates",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Permissions.Certificates"),
					},
					{
						Name:     "permissions_storage",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("Permissions.Storage"),
					},
				},
			},
			{
				Name:     "azure_keyvault_vault_private_endpoint_connections",
				Resolver: fetchKeyvaultVaultPrivateEndpointConnections,
				Columns: []schema.Column{
					{
						Name:     "vault_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "private_endpoint_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PrivateEndpointConnectionProperties.PrivateEndpoint.ID"),
					},
					{
						Name:     "private_link_service_connection_state_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:     "private_link_service_connection_state_description",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:     "private_link_service_connection_state_action_required",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.ActionRequired"),
					},
					{
						Name:     "provisioning_state",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("PrivateEndpointConnectionProperties.ProvisioningState"),
					},
				},
			},
			{
				Name:     "azure_keyvault_vault_keys",
				Resolver: fetchKeyvaultVaultKeys,
				Columns: []schema.Column{
					{
						Name:     "vault_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "attributes_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("KeyProperties.Attributes.Enabled"),
					},
					{
						Name:     "attributes_not_before",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("KeyProperties.Attributes.NotBefore"),
					},
					{
						Name:     "attributes_expires",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("KeyProperties.Attributes.Expires"),
					},
					{
						Name:     "attributes_created",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("KeyProperties.Attributes.Created"),
					},
					{
						Name:     "attributes_updated",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("KeyProperties.Attributes.Updated"),
					},
					{
						Name:     "attributes_recovery_level",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KeyProperties.Attributes.RecoveryLevel"),
					},
					{
						Name:     "kty",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KeyProperties.Kty"),
					},
					{
						Name:     "key_ops",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("KeyProperties.KeyOps"),
					},
					{
						Name:     "key_size",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("KeyProperties.KeySize"),
					},
					{
						Name:     "curve_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KeyProperties.CurveName"),
					},
					{
						Name:     "key_uri",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KeyProperties.KeyURI"),
					},
					{
						Name:     "key_uri_with_version",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("KeyProperties.KeyURIWithVersion"),
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ID"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
					{
						Name: "location",
						Type: schema.TypeString,
					},
					{
						Name: "tags",
						Type: schema.TypeJSON,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchKeyvaultVaults(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
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
func resolveKeyvaultVaultNetworkAclsIPRules(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	vault := resource.Item.(keyvault.Vault)
	if vault.Properties.NetworkAcls == nil || vault.Properties.NetworkAcls.IPRules == nil {
		return nil
	}
	ips := make([]*string, len(*vault.Properties.NetworkAcls.IPRules))
	for i, ip := range *vault.Properties.NetworkAcls.IPRules {
		ips[i] = ip.Value
	}
	return resource.Set("network_acls_ip_rules", ips)
}

func resolveKeyvaultVaultNetworkAclsVirtualNetworkRules(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	vault := resource.Item.(keyvault.Vault)
	if vault.Properties.NetworkAcls == nil || vault.Properties.NetworkAcls.IPRules == nil {
		return nil
	}
	ipRules := make([]*string, len(*vault.Properties.NetworkAcls.VirtualNetworkRules))
	for i, rule := range *vault.Properties.NetworkAcls.VirtualNetworkRules {
		ipRules[i] = rule.ID
	}
	return resource.Set("network_acls_virtual_network_rules", ipRules)
}
func fetchKeyvaultVaultAccessPolicies(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vault := parent.Item.(keyvault.Vault)
	if vault.Properties.AccessPolicies == nil {
		return nil
	}
	res <- *vault.Properties.AccessPolicies
	return nil

}
func fetchKeyvaultVaultPrivateEndpointConnections(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vault := parent.Item.(keyvault.Vault)
	if vault.Properties.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *vault.Properties.PrivateEndpointConnections
	return nil
}

func fetchKeyvaultVaultKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vault := parent.Item.(keyvault.Vault)
	svc := meta.(*client.Client).Services().KeyVault.Keys

	resourceDetails, err := client.ParseResourceID(*vault.ID)
	if err != nil {
		return err
	}
	response, err := svc.List(ctx, resourceDetails.ResourceGroup, *vault.Name)
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
