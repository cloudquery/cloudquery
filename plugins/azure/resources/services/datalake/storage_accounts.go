package datalake

import (
	"context"
	"fmt"
	"net"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func StorageAccounts() *schema.Table {
	return &schema.Table{
		Name:         "azure_datalake_storage_accounts",
		Description:  "Data Lake Store account",
		Resolver:     fetchDatalakeStorageAccounts,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		// TODO: This table is not in good shape alot of the fields doesn't exist or reference fields that doesn't exist
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "identity_type",
				Description: "The type of encryption being used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_principal_id",
				Description: "The principal identifier associated with the encryption",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "The tenant identifier associated with the encryption",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "default_group",
				Description: "The default owner group for all new folders and files created in the Data Lake Store account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.DefaultGroup"),
			},
			{
				Name:        "encryption_config_type",
				Description: "The type of encryption configuration being used",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.EncryptionConfig.Type"),
			},
			{
				Name:        "encryption_config_key_vault_meta_info_key_vault_resource_id",
				Description: "The resource identifier for the user managed Key Vault being used to encrypt",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.EncryptionConfig.KeyVaultMetaInfo.KeyVaultResourceID"),
			},
			{
				Name:        "encryption_config_key_vault_meta_info_encryption_key_name",
				Description: "The name of the user managed encryption key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.EncryptionConfig.KeyVaultMetaInfo.EncryptionKeyName"),
			},
			{
				Name:        "encryption_config_key_vault_meta_info_encryption_key_version",
				Description: "The version of the user managed encryption key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.EncryptionConfig.KeyVaultMetaInfo.EncryptionKeyVersion"),
			},
			{
				Name:        "encryption_state",
				Description: "The current state of encryption for this Data Lake Store account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.EncryptionState"),
			},
			{
				Name:        "encryption_provisioning_state",
				Description: "The current state of encryption provisioning for this Data Lake Store account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.EncryptionProvisioningState"),
			},
			{
				Name:        "firewall_state",
				Description: "The current state of the IP address firewall for this Data Lake Store account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.FirewallState"),
			},
			{
				Name:        "firewall_allow_azure_ips",
				Description: "The current state of allowing or disallowing IPs originating within Azure through the firewall",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.FirewallAllowAzureIps"),
			},
			{
				Name:        "trusted_id_provider_state",
				Description: "The current state of the trusted identity provider feature for this Data Lake Store account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.TrustedIDProviderState"),
			},
			{
				Name:        "new_tier",
				Description: "The commitment tier to use for next month",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.NewTier"),
			},
			{
				Name:        "current_tier",
				Description: "The commitment tier in use for the current month",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.CurrentTier"),
			},
			{
				Name:        "account_id",
				Description: "The unique identifier associated with this Data Lake Store account",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.AccountID"),
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning status of the Data Lake Store account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.ProvisioningState"),
			},
			{
				Name:        "state",
				Description: "The state of the Data Lake Store account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.State"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DataLakeStoreAccountProperties.CreationTime.Time"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DataLakeStoreAccountProperties.LastModifiedTime.Time"),
			},
			{
				Name:        "endpoint",
				Description: "The full CName endpoint for this account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataLakeStoreAccountProperties.Endpoint"),
			},
			{
				Name:        "id",
				Description: "The resource identifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "The resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_datalake_storage_account_firewall_rules",
				Description: "FirewallRule data Lake Store firewall rule information",
				Resolver:    fetchDatalakeStorageAccountFirewallRules,
				Columns: []schema.Column{
					{
						Name:        "storage_account_cq_id",
						Description: "Unique CloudQuery ID of azure_datalake_storage_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "start_ip_address",
						Description: "The start IP address for the firewall rule",
						Type:        schema.TypeInet,
						Resolver:    resolveStorageAccountFirewallRulesStartIpAddress,
					},
					{
						Name:        "end_ip_address",
						Description: "The end IP address for the firewall rule",
						Type:        schema.TypeInet,
						Resolver:    resolveStorageAccountFirewallRulesEndIpAddress,
					},
					{
						Name:        "id",
						Description: "The resource identifier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
						// This looks like a deprecated field, always returns nil
						// we might want to delete it in the future
						IgnoreInTests: true,
					},
					{
						Name:        "name",
						Description: "The resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The resource type",
						Type:        schema.TypeString,
						// This looks like a deprecated field, always returns nil
						// we might want to delete it in the future
						IgnoreInTests: true,
					},
				},
			},
			{
				Name:        "azure_datalake_storage_account_virtual_network_rules",
				Description: "VirtualNetworkRule data Lake Store virtual network rule information",
				Resolver:    fetchDatalakeStorageAccountVirtualNetworkRules,
				Columns: []schema.Column{
					{
						Name:        "storage_account_cq_id",
						Description: "Unique CloudQuery ID of azure_datalake_storage_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "subnet_id",
						Description: "The resource identifier for the subnet",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkRuleProperties.SubnetID"),
					},
					{
						Name:        "id",
						Description: "The resource identifier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The resource type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_datalake_storage_account_trusted_id_providers",
				Description: "TrustedIDProvider data Lake Store trusted identity provider information",
				Resolver:    fetchDatalakeStorageAccountTrustedIdProviders,
				Columns: []schema.Column{
					{
						Name:        "storage_account_cq_id",
						Description: "Unique CloudQuery ID of azure_datalake_storage_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id_provider",
						Description: "The URL of this trusted identity provider",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TrustedIDProviderProperties.IDProvider"),
					},
					{
						Name:        "id",
						Description: "The resource identifier",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
						// This looks like a deprecated field, always returns nil
						// we might want to delete it in the future
						IgnoreInTests: true,
					},
					{
						Name:        "name",
						Description: "The resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The resource type",
						Type:        schema.TypeString,
						// This looks like a deprecated field, always returns nil
						// we might want to delete it in the future
						IgnoreInTests: true,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchDatalakeStorageAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().DataLake.DataLakeStorageAccounts
	result, err := svc.List(ctx, "", nil, nil, "", "", nil)
	if err != nil {
		return diag.WrapError(err)
	}
	for result.NotDone() {
		accounts := result.Values()
		for _, a := range accounts {
			resourceDetails, err := client.ParseResourceID(*a.ID)
			if err != nil {
				return diag.WrapError(err)
			}
			result, err := svc.Get(ctx, resourceDetails.ResourceGroup, *a.Name)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- result
		}

		if err := result.NextWithContext(ctx); err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
func fetchDatalakeStorageAccountFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(account.DataLakeStoreAccount)
	if p.FirewallRules != nil {
		res <- *p.FirewallRules
	}

	return nil
}
func resolveStorageAccountFirewallRulesStartIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(account.FirewallRule)
	i := net.ParseIP(*p.StartIPAddress)
	if i == nil {
		return diag.WrapError(fmt.Errorf("wrong format of IP: %s", *p.StartIPAddress))
	}
	return diag.WrapError(resource.Set(c.Name, i))
}
func resolveStorageAccountFirewallRulesEndIpAddress(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(account.FirewallRule)
	i := net.ParseIP(*p.EndIPAddress)
	if i == nil {
		return diag.WrapError(fmt.Errorf("wrong format of IP: %s", *p.EndIPAddress))
	}
	return diag.WrapError(resource.Set(c.Name, i))
}
func fetchDatalakeStorageAccountVirtualNetworkRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(account.DataLakeStoreAccount)
	if p.VirtualNetworkRules != nil {
		res <- *p.VirtualNetworkRules
	}
	return nil
}
func fetchDatalakeStorageAccountTrustedIdProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(account.DataLakeStoreAccount)
	if p.TrustedIDProviders != nil {
		res <- *p.TrustedIDProviders
	}
	return nil
}
