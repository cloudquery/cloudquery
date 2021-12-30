package keyvault

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	keyvault71 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func KeyvaultVaults() *schema.Table {
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
				Description: "AccessPolicyEntry an identity that have access to the key vault All identities in the array must use the same tenant ID as the key vault's tenant ID",
				Resolver:    fetchKeyvaultVaultAccessPolicies,
				Columns: []schema.Column{
					{
						Name:        "vault_cq_id",
						Description: "Unique CloudQuery ID of azure_keyvault_vaults table (FK)",
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
						Description: "Unique CloudQuery ID of azure_keyvault_vaults table (FK)",
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
				Description: "KeyItem the key item containing key metadata",
				Resolver:    fetchKeyvaultVaultKeys,
				Columns: []schema.Column{
					{
						Name:        "vault_cq_id",
						Description: "Unique CloudQuery ID of azure_keyvault_vaults table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kid",
						Description: "Key identifier",
						Type:        schema.TypeString,
					},
					{
						Name:        "recoverable_days",
						Description: "softDelete data retention days Value should be >=7 and <=90 when softDelete enabled, otherwise 0",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Attributes.RecoverableDays"),
					},
					{
						Name:        "recovery_level",
						Description: "Reflects the deletion recovery level currently in effect for keys in the current vault If it contains 'Purgeable' the key can be permanently deleted by a privileged user; otherwise, only the system can purge the key, at the end of the retention interval Possible values include: 'Purgeable', 'RecoverablePurgeable', 'Recoverable', 'RecoverableProtectedSubscription', 'CustomizedRecoverablePurgeable', 'CustomizedRecoverable', 'CustomizedRecoverableProtectedSubscription'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Attributes.RecoveryLevel"),
					},
					{
						Name:        "enabled",
						Description: "Determines whether the object is enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Attributes.Enabled"),
					},
					{
						Name:        "not_before",
						Description: "Not before date in UTC",
						Type:        schema.TypeTimestamp,
						Resolver:    resolveKeyvaultVaultKeyNotBefore,
					},
					{
						Name:        "expires",
						Description: "Expiry date in UTC",
						Type:        schema.TypeTimestamp,
						Resolver:    resolveKeyvaultVaultKeyExpires,
					},
					{
						Name:        "created",
						Description: "Creation time in UTC",
						Type:        schema.TypeTimestamp,
						Resolver:    resolveKeyvaultVaultKeyCreated,
					},
					{
						Name:        "updated",
						Description: "Last updated time in UTC",
						Type:        schema.TypeTimestamp,
						Resolver:    resolveKeyvaultVaultKeyUpdated,
					},
					{
						Name:        "tags",
						Description: "Application specific metadata in the form of key-value pairs",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "managed",
						Description: "True if the key's lifetime is managed by key vault If this is a key backing a certificate, then managed will be true",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "azure_keyvault_vault_secrets",
				Description: "SecretItem the secret item containing secret metadata",
				Resolver:    fetchKeyvaultVaultSecrets,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"vault_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "vault_cq_id",
						Description: "Unique CloudQuery ID of azure_keyvault_vaults table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Secret identifier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "recoverable_days",
						Description: "softDelete data retention days Value should be >=7 and <=90 when softDelete enabled, otherwise 0",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Attributes.RecoverableDays"),
					},
					{
						Name:        "recovery_level",
						Description: "Reflects the deletion recovery level currently in effect for secrets in the current vault If it contains 'Purgeable', the secret can be permanently deleted by a privileged user; otherwise, only the system can purge the secret, at the end of the retention interval Possible values include: 'Purgeable', 'RecoverablePurgeable', 'Recoverable', 'RecoverableProtectedSubscription', 'CustomizedRecoverablePurgeable', 'CustomizedRecoverable', 'CustomizedRecoverableProtectedSubscription'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Attributes.RecoveryLevel"),
					},
					{
						Name:        "enabled",
						Description: "Determines whether the object is enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Attributes.Enabled"),
					},
					{
						Name:        "not_before",
						Description: "Not before date in UTC",
						Type:        schema.TypeTimestamp,
						Resolver:    resolveKeyvaultVaultSecretNotBefore,
					},
					{
						Name:        "expires",
						Description: "Expiry date in UTC",
						Type:        schema.TypeTimestamp,
						Resolver:    resolveKeyvaultVaultSecretExpires,
					},
					{
						Name:        "created",
						Description: "Creation time in UTC",
						Type:        schema.TypeTimestamp,
						Resolver:    resolveKeyvaultVaultSecretCreated,
					},
					{
						Name:        "updated",
						Description: "Last updated time in UTC",
						Type:        schema.TypeTimestamp,
						Resolver:    resolveKeyvaultVaultSecretUpdated,
					},
					{
						Name:        "tags",
						Description: "Application specific metadata in the form of key-value pairs",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "content_type",
						Description: "Type of the secret value such as a password",
						Type:        schema.TypeString,
					},
					{
						Name:        "managed",
						Description: "True if the secret's lifetime is managed by key vault If this is a key backing a certificate, then managed will be true",
						Type:        schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchKeyvaultVaults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
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
func resolveKeyvaultVaultNetworkAclsIPRules(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func resolveKeyvaultVaultNetworkAclsVirtualNetworkRules(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
func fetchKeyvaultVaultAccessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vault := parent.Item.(keyvault.Vault)
	if vault.Properties.AccessPolicies == nil {
		return nil
	}
	res <- *vault.Properties.AccessPolicies
	return nil
}
func fetchKeyvaultVaultPrivateEndpointConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vault := parent.Item.(keyvault.Vault)
	if vault.Properties.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *vault.Properties.PrivateEndpointConnections
	return nil
}
func fetchKeyvaultVaultKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vault := parent.Item.(keyvault.Vault)
	svc := meta.(*client.Client).Services().KeyVault.KeyVault71
	maxResults := int32(25)
	response, err := svc.GetKeys(ctx, *vault.Properties.VaultURI, &maxResults)
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
func resolveKeyvaultVaultKeyNotBefore(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key, ok := resource.Item.(keyvault71.KeyItem)
	if !ok {
		return fmt.Errorf("not a keyvault71.KeyItem instance: %#v", resource.Item)
	}

	if key.Attributes == nil || key.Attributes.NotBefore == nil {
		return nil
	}

	return resource.Set(c.Name, time.Time(*key.Attributes.NotBefore))
}
func resolveKeyvaultVaultKeyExpires(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key, ok := resource.Item.(keyvault71.KeyItem)
	if !ok {
		return fmt.Errorf("not a keyvault71.KeyItem instance: %#v", resource.Item)
	}

	if key.Attributes == nil || key.Attributes.Expires == nil {
		return nil
	}

	return resource.Set(c.Name, time.Time(*key.Attributes.Expires))
}
func resolveKeyvaultVaultKeyCreated(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key, ok := resource.Item.(keyvault71.KeyItem)
	if !ok {
		return fmt.Errorf("not a keyvault71.KeyItem instance: %#v", resource.Item)
	}

	if key.Attributes == nil || key.Attributes.Created == nil {
		return nil
	}

	return resource.Set(c.Name, time.Time(*key.Attributes.Created))
}
func resolveKeyvaultVaultKeyUpdated(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key, ok := resource.Item.(keyvault71.KeyItem)
	if !ok {
		return fmt.Errorf("not a keyvault71.KeyItem instance: %#v", resource.Item)
	}

	if key.Attributes == nil || key.Attributes.Updated == nil {
		return nil
	}

	return resource.Set(c.Name, time.Time(*key.Attributes.Updated))
}
func fetchKeyvaultVaultSecrets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	vault, ok := parent.Item.(keyvault.Vault)
	if !ok {
		return fmt.Errorf("not a keyvault.Vault instance: %#v", parent.Item)
	}
	svc := meta.(*client.Client).Services().KeyVault.KeyVault71
	maxResults := int32(25)
	result, err := svc.GetSecrets(ctx, *vault.Properties.VaultURI, &maxResults)
	if err != nil {
		return err
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
func resolveKeyvaultVaultSecretNotBefore(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key, ok := resource.Item.(keyvault71.SecretItem)
	if !ok {
		return fmt.Errorf("not a keyvault71.SecretItem instance: %#v", resource.Item)
	}

	if key.Attributes == nil || key.Attributes.NotBefore == nil {
		return nil
	}

	return resource.Set(c.Name, time.Time(*key.Attributes.NotBefore))
}
func resolveKeyvaultVaultSecretExpires(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key, ok := resource.Item.(keyvault71.SecretItem)
	if !ok {
		return fmt.Errorf("not a keyvault71.SecretItem instance: %#v", resource.Item)
	}

	if key.Attributes == nil || key.Attributes.Expires == nil {
		return nil
	}

	return resource.Set(c.Name, time.Time(*key.Attributes.Expires))
}
func resolveKeyvaultVaultSecretCreated(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key, ok := resource.Item.(keyvault71.SecretItem)
	if !ok {
		return fmt.Errorf("not a keyvault71.SecretItem instance: %#v", resource.Item)
	}

	if key.Attributes == nil || key.Attributes.Created == nil {
		return nil
	}

	return resource.Set(c.Name, time.Time(*key.Attributes.Created))
}
func resolveKeyvaultVaultSecretUpdated(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key, ok := resource.Item.(keyvault71.SecretItem)
	if !ok {
		return fmt.Errorf("not a keyvault71.SecretItem instance: %#v", resource.Item)
	}

	if key.Attributes == nil || key.Attributes.Updated == nil {
		return nil
	}

	return resource.Set(c.Name, time.Time(*key.Attributes.Updated))
}
