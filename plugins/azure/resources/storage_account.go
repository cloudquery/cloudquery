package resources

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func StorageAccounts() *schema.Table {
	return &schema.Table{
		Name:         "azure_storage_accounts",
		Resolver:     fetchStorageAccounts,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "sku_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Sku.Name"),
			},
			{
				Name:     "sku_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Sku.Tier"),
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name:     "identity_principal_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:     "identity_tenant_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:     "identity_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Identity.Type"),
			},
			{
				Name:     "provisioning_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.ProvisioningState"),
			},
			{
				Name:     "primary_endpoints_blob",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.Blob"),
			},
			{
				Name:     "primary_endpoints_queue",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.Queue"),
			},
			{
				Name:     "primary_endpoints_table",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.Table"),
			},
			{
				Name:     "primary_endpoints_file",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.File"),
			},
			{
				Name:     "primary_endpoints_web",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.Web"),
			},
			{
				Name:     "primary_endpoints_dfs",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.Dfs"),
			},
			{
				Name:     "primary_endpoints_microsoft_endpoints_blob",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Blob"),
			},
			{
				Name:     "primary_endpoints_microsoft_endpoints_queue",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Queue"),
			},
			{
				Name:     "primary_endpoints_microsoft_endpoints_table",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Table"),
			},
			{
				Name:     "primary_endpoints_microsoft_endpoints_file",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.File"),
			},
			{
				Name:     "primary_endpoints_microsoft_endpoints_web",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Web"),
			},
			{
				Name:     "primary_endpoints_microsoft_endpoints_dfs",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Dfs"),
			},
			{
				Name:     "primary_endpoints_internet_endpoints_blob",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.InternetEndpoints.Blob"),
			},
			{
				Name:     "primary_endpoints_internet_endpoints_file",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.InternetEndpoints.File"),
			},
			{
				Name:     "primary_endpoints_internet_endpoints_web",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.InternetEndpoints.Web"),
			},
			{
				Name:     "primary_endpoints_internet_endpoints_dfs",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryEndpoints.InternetEndpoints.Dfs"),
			},
			{
				Name:     "primary_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.PrimaryLocation"),
			},
			{
				Name:     "status_of_primary",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.StatusOfPrimary"),
			},
			{
				Name:     "last_geo_failover_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccountProperties.LastGeoFailoverTime.Time"),
			},
			{
				Name:     "secondary_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryLocation"),
			},
			{
				Name:     "status_of_secondary",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.StatusOfSecondary"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccountProperties.CreationTime.Time"),
			},
			{
				Name:     "custom_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.CustomDomain.Name"),
			},
			{
				Name:     "custom_domain_use_sub_domain_name",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.CustomDomain.UseSubDomainName"),
			},
			{
				Name:     "secondary_endpoints_blob",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.Blob"),
			},
			{
				Name:     "secondary_endpoints_queue",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.Queue"),
			},
			{
				Name:     "secondary_endpoints_table",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.Table"),
			},
			{
				Name:     "secondary_endpoints_file",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.File"),
			},
			{
				Name:     "secondary_endpoints_web",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.Web"),
			},
			{
				Name:     "secondary_endpoints_dfs",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.Dfs"),
			},
			{
				Name:     "secondary_endpoints_microsoft_endpoints_blob",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Blob"),
			},
			{
				Name:     "secondary_endpoints_microsoft_endpoints_queue",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Queue"),
			},
			{
				Name:     "secondary_endpoints_microsoft_endpoints_table",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Table"),
			},
			{
				Name:     "secondary_endpoints_microsoft_endpoints_file",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.File"),
			},
			{
				Name:     "secondary_endpoints_microsoft_endpoints_web",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Web"),
			},
			{
				Name:     "secondary_endpoints_microsoft_endpoints_dfs",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Dfs"),
			},
			{
				Name:     "secondary_endpoints_internet_endpoints_blob",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.InternetEndpoints.Blob"),
			},
			{
				Name:     "secondary_endpoints_internet_endpoints_file",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.InternetEndpoints.File"),
			},
			{
				Name:     "secondary_endpoints_internet_endpoints_web",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.InternetEndpoints.Web"),
			},
			{
				Name:     "secondary_endpoints_internet_endpoints_dfs",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.SecondaryEndpoints.InternetEndpoints.Dfs"),
			},
			{
				Name:     "encryption_services_blob_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.Blob.Enabled"),
			},
			{
				Name:     "encryption_services_blob_last_enabled_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.Blob.LastEnabledTime.Time"),
			},
			{
				Name:     "encryption_services_blob_key_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.Blob.KeyType"),
			},
			{
				Name:     "encryption_services_file_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.File.Enabled"),
			},
			{
				Name:     "encryption_services_file_last_enabled_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.File.LastEnabledTime.Time"),
			},
			{
				Name:     "encryption_services_file_key_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.File.KeyType"),
			},
			{
				Name:     "encryption_services_table_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.Table.Enabled"),
			},
			{
				Name:     "encryption_services_table_last_enabled_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.Table.LastEnabledTime.Time"),
			},
			{
				Name:     "encryption_services_table_key_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.Table.KeyType"),
			},
			{
				Name:     "encryption_services_queue_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.Queue.Enabled"),
			},
			{
				Name:     "encryption_services_queue_last_enabled_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.Queue.LastEnabledTime.Time"),
			},
			{
				Name:     "encryption_services_queue_key_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.Encryption.Services.Queue.KeyType"),
			},
			{
				Name:     "encryption_key_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.Encryption.KeySource"),
			},
			{
				Name:     "encryption_require_infrastructure_encryption",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.Encryption.RequireInfrastructureEncryption"),
			},
			{
				Name:     "encryption_key_vault_key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.KeyName"),
			},
			{
				Name:     "encryption_key_vault_key_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.KeyVersion"),
			},
			{
				Name:     "encryption_key_vault_key_vault_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.KeyVaultURI"),
			},
			{
				Name:     "encryption_key_vault_current_versioned_key_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.CurrentVersionedKeyIdentifier"),
			},
			{
				Name:     "encryption_key_vault_last_key_rotation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.LastKeyRotationTimestamp.Time"),
			},
			{
				Name:     "access_tier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.AccessTier"),
			},
			{
				Name:     "files_identity_based_authentication_directory_service_options",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.DirectoryServiceOptions"),
			},
			{
				Name:     "files_identity_based_authentication_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.DomainName"),
			},
			{
				Name:     "files_identity_based_authentication_net_bios_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.NetBiosDomainName"),
			},
			{
				Name:     "files_identity_based_authentication_forest_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.ForestName"),
			},
			{
				Name:     "files_identity_based_authentication_domain_guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.DomainGUID"),
			},
			{
				Name:     "files_identity_based_authentication_domain_sid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.DomainSid"),
			},
			{
				Name:     "files_identity_based_authentication_storage_sid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.AzureStorageSid"),
			},
			{
				Name:     "enable_https_traffic_only",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.EnableHTTPSTrafficOnly"),
			},
			{
				Name:     "network_rule_set_bypass",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.NetworkRuleSet.Bypass"),
			},
			{
				Name:     "network_rule_set_default_action",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.NetworkRuleSet.DefaultAction"),
			},
			{
				Name:     "is_hns_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.IsHnsEnabled"),
			},
			{
				Name:     "geo_replication_stats_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.GeoReplicationStats.Status"),
			},
			{
				Name:     "geo_replication_stats_last_sync_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccountProperties.GeoReplicationStats.LastSyncTime.Time"),
			},
			{
				Name:     "geo_replication_stats_can_failover",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.GeoReplicationStats.CanFailover"),
			},
			{
				Name:     "failover_in_progress",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.FailoverInProgress"),
			},
			{
				Name:     "large_file_shares_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.LargeFileSharesState"),
			},
			{
				Name:     "routing_preference_routing_choice",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.RoutingPreference.RoutingChoice"),
			},
			{
				Name:     "routing_preference_publish_microsoft_endpoints",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.RoutingPreference.PublishMicrosoftEndpoints"),
			},
			{
				Name:     "routing_preference_publish_internet_endpoints",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.RoutingPreference.PublishInternetEndpoints"),
			},
			{
				Name:     "blob_restore_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.BlobRestoreStatus.Status"),
			},
			{
				Name:     "blob_restore_status_failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.BlobRestoreStatus.FailureReason"),
			},
			{
				Name:     "blob_restore_status_restore_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.BlobRestoreStatus.RestoreID"),
			},
			{
				Name:     "blob_restore_status_parameters_time_to_restore_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AccountProperties.BlobRestoreStatus.Parameters.TimeToRestore.Time"),
			},
			{
				Name:     "blob_restore_status_parameters_blob_ranges",
				Type:     schema.TypeJSON,
				Resolver: resolveRestoreStatusParametersBlobRanges,
			},
			{
				Name:     "allow_blob_public_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccountProperties.AllowBlobPublicAccess"),
			},
			{
				Name:     "minimum_tls_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountProperties.MinimumTLSVersion"),
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "location",
				Type: schema.TypeString,
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
		},
		Relations: []*schema.Table{
			{
				Name:     "azure_storage_account_network_rule_set_virtual_network_rules",
				Resolver: fetchStorageAccountNetworkRuleSetVirtualNetworkRules,
				Columns: []schema.Column{
					{
						Name:     "account_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "virtual_network_resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("VirtualNetworkResourceID"),
					},
					{
						Name: "action",
						Type: schema.TypeString,
					},
					{
						Name: "state",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_storage_account_network_rule_set_ip_rules",
				Resolver: fetchStorageAccountNetworkRuleSetIpRules,
				Columns: []schema.Column{
					{
						Name:     "account_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "ip_address_or_range",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("IPAddressOrRange"),
					},
					{
						Name: "action",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "azure_storage_account_private_endpoint_connections",
				Resolver: fetchStorageAccountPrivateEndpointConnections,
				Columns: []schema.Column{
					{
						Name:     "account_id",
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
				},
			},
			storageContainers(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchStorageAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Storage.Accounts
	response, err := svc.List(ctx)
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
func fetchStorageAccountNetworkRuleSetVirtualNetworkRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	account := parent.Item.(storage.Account)
	if account.NetworkRuleSet == nil || account.NetworkRuleSet.VirtualNetworkRules == nil {
		return nil
	}
	res <- *account.NetworkRuleSet.VirtualNetworkRules
	return nil
}

func fetchStorageAccountNetworkRuleSetIpRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	account := parent.Item.(storage.Account)
	if account.NetworkRuleSet == nil || account.NetworkRuleSet.IPRules == nil {
		return nil
	}
	res <- *account.NetworkRuleSet.IPRules
	return nil
}
func fetchStorageAccountPrivateEndpointConnections(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	account := parent.Item.(storage.Account)
	if account.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *account.PrivateEndpointConnections
	return nil
}
func resolveRestoreStatusParametersBlobRanges(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	account := resource.Item.(storage.Account)
	if account.BlobRestoreStatus == nil || account.BlobRestoreStatus.Parameters == nil {
		return nil
	}
	data, err := json.Marshal(account.BlobRestoreStatus.Parameters.BlobRanges)
	if err != nil {
		return err
	}
	resource.Set("blob_restore_status_parameters_blob_ranges", data)
	return nil
}
