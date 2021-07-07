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
		Description:  "Azure ketvault vault",
		Resolver:     fetchKeyvaultVaults,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "id",
				Description: "Fully qualified identifier of the key vault resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Name of the key vault resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type of the key vault resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Azure location of the key vault resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Tags assigned to the key vault resource",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "tenant_id",
				Description: "The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Properties.TenantID"),
			},
			{
				Name:        "sku_family",
				Description: "SKU family name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Sku.Family"),
			},
			{
				Name:        "sku_name",
				Description: "SKU name to specify whether the key vault is a standard vault or a premium vault Possible values include: 'Standard', 'Premium'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Sku.Name"),
			},
			{
				Name:        "vault_uri",
				Description: "The URI of the vault for performing operations on keys and secrets",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.VaultURI"),
			},
			{
				Name:        "enabled_for_deployment",
				Description: "Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnabledForDeployment"),
			},
			{
				Name:        "enabled_for_disk_encryption",
				Description: "Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnabledForDiskEncryption"),
			},
			{
				Name:        "enabled_for_template_deployment",
				Description: "Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnabledForTemplateDeployment"),
			},
			{
				Name:        "enable_soft_delete",
				Description: "Property to specify whether the 'soft delete' functionality is enabled for this key vault If it's not set to any value(true or false) when creating new key vault, it will be set to true by default Once set to true, it cannot be reverted to false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnableSoftDelete"),
			},
			{
				Name:        "soft_delete_retention_in_days",
				Description: "softDelete data retention days It accepts >=7 and <=90",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Properties.SoftDeleteRetentionInDays"),
			},
			{
				Name:        "enable_rbac_authorization",
				Description: "Property that controls how data actions are authorized When true, the key vault will use Role Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties will be  ignored (warning: this is a preview feature) When false, the key vault will use the access policies specified in vault properties, and any policy stored on Azure Resource Manager will be ignored If null or not specified, the vault is created with the default value of false Note that management actions are always authorized with RBAC",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnableRbacAuthorization"),
			},
			{
				Name:        "create_mode",
				Description: "The vault's create mode to indicate whether the vault need to be recovered or not Possible values include: 'CreateModeRecover', 'CreateModeDefault'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.CreateMode"),
			},
			{
				Name:        "enable_purge_protection",
				Description: "Property specifying whether protection against purge is enabled for this vault Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion The setting is effective only if soft delete is also enabled Enabling this functionality is irreversible - that is, the property does not accept false as its value",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Properties.EnablePurgeProtection"),
			},
			{
				Name:        "network_acls_bypass",
				Description: "Tells what traffic can bypass network rules This can be 'AzureServices' or 'None'  If not specified the default is 'AzureServices' Possible values include: 'AzureServices', 'None'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.NetworkAcls.Bypass"),
			},
			{
				Name:        "network_acls_default_action",
				Description: "The default action when no rule from ipRules and from virtualNetworkRules match This is only used after the bypass property has been evaluated Possible values include: 'Allow', 'Deny'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.NetworkAcls.DefaultAction"),
			},
			{
				Name:        "network_acls_ip_rules",
				Description: "The list of IP address rules",
				Type:        schema.TypeStringArray,
				Resolver:    resolveKeyvaultVaultNetworkAclsIPRules,
			},
			{
				Name:        "network_acls_virtual_network_rules",
				Description: "The list of virtual network rules",
				Type:        schema.TypeStringArray,
				Resolver:    resolveKeyvaultVaultNetworkAclsVirtualNetworkRules,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_keyvault_vault_access_policies",
				Description: "AccessPolicyEntry an identity that have access to the key vault",
				Resolver:    fetchKeyvaultVaultAccessPolicies,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"vault_cq_id", "object_id"}},
				Columns: []schema.Column{
					{
						Name:        "vault_cq_id",
						Description: "Unique ID of azure_keyvault_vaults table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "tenant_id",
						Description: "The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("TenantID"),
					},
					{
						Name:        "object_id",
						Description: "The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault The object ID must be unique for the list of access policies",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ObjectID"),
					},
					{
						Name:        "application_id",
						Description: "Application ID of the client making request on behalf of a principal",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("ApplicationID"),
					},
					{
						Name:        "permissions_keys",
						Description: "Permissions to keys",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Permissions.Keys"),
					},
					{
						Name:        "permissions_secrets",
						Description: "Permissions to secrets",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Permissions.Secrets"),
					},
					{
						Name:        "permissions_certificates",
						Description: "Permissions to certificates",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Permissions.Certificates"),
					},
					{
						Name:        "permissions_storage",
						Description: "Permissions to storage accounts",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("Permissions.Storage"),
					},
				},
			},
			{
				Name:        "azure_keyvault_vault_private_endpoint_connections",
				Description: "Azure ketvault vault endpoint connection",
				Resolver:    fetchKeyvaultVaultPrivateEndpointConnections,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"vault_cq_id", "private_endpoint_id"}},
				Columns: []schema.Column{
					{
						Name:        "vault_cq_id",
						Description: "Unique ID of azure_keyvault_vaults table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "private_endpoint_id",
						Description: "Full identifier of the private endpoint resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateEndpoint.ID"),
					},
					{
						Name:        "private_link_service_connection_state_status",
						Description: "Indicates whether the connection has been approved, rejected or removed by the key vault owner Possible values include: 'PrivateEndpointServiceConnectionStatusPending', 'PrivateEndpointServiceConnectionStatusApproved', 'PrivateEndpointServiceConnectionStatusRejected', 'PrivateEndpointServiceConnectionStatusDisconnected'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "private_link_service_connection_state_description",
						Description: "The reason for approval or rejection",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "private_link_service_connection_state_action_required",
						Description: "A message indicating if changes on the service provider require any updates on the consumer",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.ActionRequired"),
					},
					{
						Name:        "provisioning_state",
						Description: "Provisioning state of the private endpoint connection Possible values include: 'Succeeded', 'Creating', 'Updating', 'Deleting', 'Failed', 'Disconnected'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.ProvisioningState"),
					},
				},
			},
			{
				Name:        "azure_keyvault_vault_keys",
				Description: "Azure ketvault vault key",
				Resolver:    fetchKeyvaultVaultKeys,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"vault_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "vault_cq_id",
						Description: "Unique ID of azure_keyvault_vaults table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "attributes_enabled",
						Description: "Determines whether or not the object is enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("KeyProperties.Attributes.Enabled"),
					},
					{
						Name:        "attributes_not_before",
						Description: "Not before date in seconds since 1970-01-01T00:00:00Z",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("KeyProperties.Attributes.NotBefore"),
					},
					{
						Name:        "attributes_expires",
						Description: "Expiry date in seconds since 1970-01-01T00:00:00Z",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("KeyProperties.Attributes.Expires"),
					},
					{
						Name:        "attributes_created",
						Description: "Creation time in seconds since 1970-01-01T00:00:00Z",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("KeyProperties.Attributes.Created"),
					},
					{
						Name:        "attributes_updated",
						Description: "Last updated time in seconds since 1970-01-01T00:00:00Z",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("KeyProperties.Attributes.Updated"),
					},
					{
						Name:        "attributes_recovery_level",
						Description: "The deletion recovery level currently in effect for the object If it contains 'Purgeable', then the object can be permanently deleted by a privileged user; otherwise, only the system can purge the object at the end of the retention interval Possible values include: 'Purgeable', 'RecoverablePurgeable', 'Recoverable', 'RecoverableProtectedSubscription'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KeyProperties.Attributes.RecoveryLevel"),
					},
					{
						Name:        "kty",
						Description: "The type of the key For valid values, see JsonWebKeyType Possible values include: 'EC', 'ECHSM', 'RSA', 'RSAHSM'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KeyProperties.Kty"),
					},
					{
						Name:        "key_ops",
						Description: "Enumerates the values for json web key operation",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("KeyProperties.KeyOps"),
					},
					{
						Name:        "key_size",
						Description: "The key size in bits For example: 2048, 3072, or 4096 for RSA",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("KeyProperties.KeySize"),
					},
					{
						Name:        "curve_name",
						Description: "The elliptic curve name For valid values, see JsonWebKeyCurveName Possible values include: 'P256', 'P384', 'P521', 'P256K'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KeyProperties.CurveName"),
					},
					{
						Name:        "key_uri",
						Description: "The URI to retrieve the current version of the key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KeyProperties.KeyURI"),
					},
					{
						Name:        "key_uri_with_version",
						Description: "The URI to retrieve the specific version of the key",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("KeyProperties.KeyURIWithVersion"),
					},
					{
						Name:        "id",
						Description: "Fully qualified identifier of the key vault resource",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Name of the key vault resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type of the key vault resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "location",
						Description: "Azure location of the key vault resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "Tags assigned to the key vault resource",
						Type:        schema.TypeJSON,
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
