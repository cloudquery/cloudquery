package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"
)

func StorageAccounts() *schema.Table {
	return &schema.Table{
		Name:         "azure_storage_accounts",
		Description:  "Azure storage account",
		Resolver:     fetchStorageAccounts,
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
				Name:        "sku_name",
				Description: "Possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS', 'PremiumZRS', 'StandardGZRS', 'StandardRAGZRS'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Possible values include: 'Standard', 'Premium'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "kind",
				Description: "Gets the Kind Possible values include: 'Storage', 'StorageV2', 'BlobStorage', 'FileStorage', 'BlockBlobStorage'",
				Type:        schema.TypeString,
			},
			{
				Name:          "identity_principal_id",
				Description:   "The principal ID of resource identity",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.PrincipalID"),
				IgnoreInTests: true,
			},
			{
				Name:          "identity_tenant_id",
				Description:   "The tenant ID of resource",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Identity.TenantID"),
				IgnoreInTests: true,
			},
			{
				Name:        "identity_type",
				Description: "The identity type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "provisioning_state",
				Description: "Gets the status of the storage account at the time the operation was called Possible values include: 'Creating', 'ResolvingDNS', 'Succeeded'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.ProvisioningState"),
			},
			{
				Name:        "primary_endpoints_blob",
				Description: "Gets the blob endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.PrimaryEndpoints.Blob"),
			},
			{
				Name:        "primary_endpoints_queue",
				Description: "Gets the queue endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.PrimaryEndpoints.Queue"),
			},
			{
				Name:        "primary_endpoints_table",
				Description: "Gets the table endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.PrimaryEndpoints.Table"),
			},
			{
				Name:        "primary_endpoints_file",
				Description: "Gets the file endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.PrimaryEndpoints.File"),
			},
			{
				Name:        "primary_endpoints_web",
				Description: "Gets the web endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.PrimaryEndpoints.Web"),
			},
			{
				Name:        "primary_endpoints_dfs",
				Description: "Gets the dfs endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.PrimaryEndpoints.Dfs"),
			},
			{
				Name:          "primary_endpoints_microsoft_endpoints_blob",
				Description:   "Gets the blob endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Blob"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_endpoints_microsoft_endpoints_queue",
				Description:   "Gets the queue endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Queue"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_endpoints_microsoft_endpoints_table",
				Description:   "Gets the table endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Table"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_endpoints_microsoft_endpoints_file",
				Description:   "Gets the file endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.File"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_endpoints_microsoft_endpoints_web",
				Description:   "Gets the web endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Web"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_endpoints_microsoft_endpoints_dfs",
				Description:   "Gets the dfs endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.MicrosoftEndpoints.Dfs"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_endpoints_internet_endpoints_blob",
				Description:   "Gets the blob endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.InternetEndpoints.Blob"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_endpoints_internet_endpoints_file",
				Description:   "Gets the file endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.InternetEndpoints.File"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_endpoints_internet_endpoints_web",
				Description:   "Gets the web endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.InternetEndpoints.Web"),
				IgnoreInTests: true,
			},
			{
				Name:          "primary_endpoints_internet_endpoints_dfs",
				Description:   "Gets the dfs endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.PrimaryEndpoints.InternetEndpoints.Dfs"),
				IgnoreInTests: true,
			},
			{
				Name:        "primary_location",
				Description: "Gets the location of the primary data center for the storage account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.PrimaryLocation"),
			},
			{
				Name:        "status_of_primary",
				Description: "Gets the status indicating whether the primary location of the storage account is available or unavailable Possible values include: 'Available', 'Unavailable'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.StatusOfPrimary"),
			},
			{
				Name:        "last_geo_failover_time",
				Description: "Gets the timestamp of the most recent instance of a failover to the secondary location Only the most recent timestamp is retained This element is not returned if there has never been a failover instance Only available if the accountType is Standard_GRS or Standard_RAGRS",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.LastGeoFailoverTime.Time"),
			},
			{
				Name:          "secondary_location",
				Description:   "Gets the location of the geo-replicated secondary for the storage account Only available if the accountType is Standard_GRS or Standard_RAGRS",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryLocation"),
				IgnoreInTests: true,
			},
			{
				Name:        "status_of_secondary",
				Description: "Gets the status indicating whether the secondary location of the storage account is available or unavailable Only available if the SKU name is Standard_GRS or Standard_RAGRS Possible values include: 'Available', 'Unavailable'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.StatusOfSecondary"),
			},
			{
				Name:        "creation_time",
				Description: "Gets the creation date and time of the storage account in UTC",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.CreationTime.Time"),
			},
			{
				Name:          "custom_domain_name",
				Description:   "Gets or sets the custom domain name assigned to the storage account Name is the CNAME source",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.CustomDomain.Name"),
				IgnoreInTests: true,
			},
			{
				Name:          "custom_domain_use_sub_domain_name",
				Description:   "Indicates whether indirect CName validation is enabled Default value is false This should only be set on updates",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("AccountProperties.CustomDomain.UseSubDomainName"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_blob",
				Description:   "Gets the blob endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.Blob"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_queue",
				Description:   "Gets the queue endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.Queue"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_table",
				Description:   "Gets the table endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.Table"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_file",
				Description:   "Gets the file endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.File"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_web",
				Description:   "Gets the web endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.Web"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_dfs",
				Description:   "Gets the dfs endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.Dfs"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_microsoft_endpoints_blob",
				Description:   "Gets the blob endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Blob"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_microsoft_endpoints_queue",
				Description:   "Gets the queue endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Queue"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_microsoft_endpoints_table",
				Description:   "Gets the table endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Table"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_microsoft_endpoints_file",
				Description:   "Gets the file endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.File"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_microsoft_endpoints_web",
				Description:   "Gets the web endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Web"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_microsoft_endpoints_dfs",
				Description:   "Gets the dfs endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.MicrosoftEndpoints.Dfs"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_internet_endpoints_blob",
				Description:   "Gets the blob endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.InternetEndpoints.Blob"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_internet_endpoints_file",
				Description:   "Gets the file endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.InternetEndpoints.File"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_internet_endpoints_web",
				Description:   "Gets the web endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.InternetEndpoints.Web"),
				IgnoreInTests: true,
			},
			{
				Name:          "secondary_endpoints_internet_endpoints_dfs",
				Description:   "Gets the dfs endpoint",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.SecondaryEndpoints.InternetEndpoints.Dfs"),
				IgnoreInTests: true,
			},
			{
				Name:        "encryption_services_blob_enabled",
				Description: "A boolean indicating whether or not the service encrypts the data as it is stored",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.Blob.Enabled"),
			},
			{
				Name:        "encryption_services_blob_last_enabled_time",
				Description: "Gets a rough estimate of the date/time when the encryption was last enabled by the user Only returned when encryption is enabled There might be some unencrypted blobs which were written after this time, as it is just a rough estimate",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.Blob.LastEnabledTime.Time"),
			},
			{
				Name:        "encryption_services_blob_key_type",
				Description: "Encryption key type to be used for the encryption service 'Account' key type implies that an account-scoped encryption key will be used 'Service' key type implies that a default service key is used Possible values include: 'KeyTypeService', 'KeyTypeAccount'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.Blob.KeyType"),
			},
			{
				Name:        "encryption_services_file_enabled",
				Description: "A boolean indicating whether or not the service encrypts the data as it is stored",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.File.Enabled"),
			},
			{
				Name:        "encryption_services_file_last_enabled_time",
				Description: "Gets a rough estimate of the date/time when the encryption was last enabled by the user Only returned when encryption is enabled There might be some unencrypted blobs which were written after this time, as it is just a rough estimate",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.File.LastEnabledTime.Time"),
			},
			{
				Name:        "encryption_services_file_key_type",
				Description: "Encryption key type to be used for the encryption service 'Account' key type implies that an account-scoped encryption key will be used 'Service' key type implies that a default service key is used Possible values include: 'KeyTypeService', 'KeyTypeAccount'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.File.KeyType"),
			},
			{
				Name:          "encryption_services_table_enabled",
				Description:   "A boolean indicating whether or not the service encrypts the data as it is stored",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("AccountProperties.Encryption.Services.Table.Enabled"),
				IgnoreInTests: true,
			},
			{
				Name:        "encryption_services_table_last_enabled_time",
				Description: "Gets a rough estimate of the date/time when the encryption was last enabled by the user Only returned when encryption is enabled There might be some unencrypted blobs which were written after this time, as it is just a rough estimate",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.Table.LastEnabledTime.Time"),
			},
			{
				Name:        "encryption_services_table_key_type",
				Description: "Encryption key type to be used for the encryption service 'Account' key type implies that an account-scoped encryption key will be used 'Service' key type implies that a default service key is used Possible values include: 'KeyTypeService', 'KeyTypeAccount'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.Table.KeyType"),
			},
			{
				Name:          "encryption_services_queue_enabled",
				Description:   "A boolean indicating whether or not the service encrypts the data as it is stored",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("AccountProperties.Encryption.Services.Queue.Enabled"),
				IgnoreInTests: true,
			},
			{
				Name:        "encryption_services_queue_last_enabled_time",
				Description: "Gets a rough estimate of the date/time when the encryption was last enabled by the user Only returned when encryption is enabled There might be some unencrypted blobs which were written after this time, as it is just a rough estimate",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.Queue.LastEnabledTime.Time"),
			},
			{
				Name:        "encryption_services_queue_key_type",
				Description: "Encryption key type to be used for the encryption service 'Account' key type implies that an account-scoped encryption key will be used 'Service' key type implies that a default service key is used Possible values include: 'KeyTypeService', 'KeyTypeAccount'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.Services.Queue.KeyType"),
			},
			{
				Name:        "encryption_key_source",
				Description: "The encryption keySource (provider) Possible values (case-insensitive):  MicrosoftStorage, MicrosoftKeyvault Possible values include: 'KeySourceMicrosoftStorage', 'KeySourceMicrosoftKeyvault'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.KeySource"),
			},
			{
				Name:          "encryption_require_infrastructure_encryption",
				Description:   "A boolean indicating whether or not the service applies a secondary layer of encryption with platform managed keys for data at rest",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("AccountProperties.Encryption.RequireInfrastructureEncryption"),
				IgnoreInTests: true,
			},
			{
				Name:          "encryption_key_vault_properties_key_name",
				Description:   "The name of KeyVault key",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.KeyName"),
				IgnoreInTests: true,
			},
			{
				Name:          "encryption_key_vault_properties_key_version",
				Description:   "The version of KeyVault key",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.KeyVersion"),
				IgnoreInTests: true,
			},
			{
				Name:          "encryption_key_vault_properties_key_vault_uri",
				Description:   "The Uri of KeyVault",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.KeyVaultURI"),
				IgnoreInTests: true,
			},
			{
				Name:          "encryption_key_current_versioned_key_identifier",
				Description:   "The object identifier of the current versioned Key Vault Key in use",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.CurrentVersionedKeyIdentifier"),
				IgnoreInTests: true,
			},
			{
				Name:        "encryption_key_last_key_rotation_timestamp_time",
				Description: "Timestamp of last rotation of the Key Vault Key",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.Encryption.KeyVaultProperties.LastKeyRotationTimestamp.Time"),
			},
			{
				Name:        "access_tier",
				Description: "Required for storage accounts where kind = BlobStorage The access tier used for billing Possible values include: 'Hot', 'Cool'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.AccessTier"),
			},
			{
				Name:        "files_identity_auth_directory_service_options",
				Description: "Indicates the directory service used Possible values include: 'DirectoryServiceOptionsNone', 'DirectoryServiceOptionsAADDS', 'DirectoryServiceOptionsAD'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.DirectoryServiceOptions"),
			},
			{
				Name:          "files_identity_auth_ad_properties_domain_name",
				Description:   "Specifies the primary domain that the AD DNS server is authoritative for",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.DomainName"),
				IgnoreInTests: true,
			},
			{
				Name:          "files_identity_auth_ad_properties_net_bios_domain_name",
				Description:   "Specifies the NetBIOS domain name",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.NetBiosDomainName"),
				IgnoreInTests: true,
			},
			{
				Name:          "files_identity_auth_ad_properties_forest_name",
				Description:   "Specifies the Active Directory forest to get",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.ForestName"),
				IgnoreInTests: true,
			},
			{
				Name:          "files_identity_auth_ad_properties_domain_guid",
				Description:   "Specifies the domain GUID",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.DomainGUID"),
				IgnoreInTests: true,
			},
			{
				Name:          "files_identity_auth_ad_properties_net_bios_domain_sid",
				Description:   "Specifies the security identifier (SID)",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.DomainSid"),
				IgnoreInTests: true,
			},
			{
				Name:          "files_identity_auth_ad_properties_azure_storage_sid",
				Description:   "Specifies the security identifier (SID) for Azure Storage",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties.AzureStorageSid"),
				IgnoreInTests: true,
			},
			{
				Name:        "enable_https_traffic_only",
				Description: "Allows https traffic only to storage service if sets to true",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccountProperties.EnableHTTPSTrafficOnly"),
			},
			{
				Name:        "network_rule_set_bypass",
				Description: "Specifies whether traffic is bypassed for Logging/Metrics/AzureServices Possible values are any combination of Logging|Metrics|AzureServices (For example, \"Logging, Metrics\"), or None to bypass none of those traffics Possible values include: 'None', 'Logging', 'Metrics', 'AzureServices'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.NetworkRuleSet.Bypass"),
			},
			{
				Name:        "network_rule_set_default_action",
				Description: "Specifies the default action of allow or deny when no other rules match Possible values include: 'DefaultActionAllow', 'DefaultActionDeny'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.NetworkRuleSet.DefaultAction"),
			},
			{
				Name:          "is_hns_enabled",
				Description:   "Account HierarchicalNamespace enabled if sets to true",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("AccountProperties.IsHnsEnabled"),
				IgnoreInTests: true,
			},
			{
				Name:        "geo_replication_stats_status",
				Description: "The status of the secondary location Possible values are: - Live: Indicates that the secondary location is active and operational - Bootstrap: Indicates initial synchronization from the primary location to the secondary location is in progressThis typically occurs when replication is first enabled - Unavailable: Indicates that the secondary location is temporarily unavailable Possible values include: 'GeoReplicationStatusLive', 'GeoReplicationStatusBootstrap', 'GeoReplicationStatusUnavailable'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.GeoReplicationStats.Status"),
			},
			{
				Name:        "geo_replication_stats_last_sync_time",
				Description: "All primary writes preceding this UTC date/time value are guaranteed to be available for read operations Primary writes following this point in time may or may not be available for reads Element may be default value if value of LastSyncTime is not available, this can happen if secondary is offline or we are in bootstrap",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.GeoReplicationStats.LastSyncTime.Time"),
			},
			{
				Name:          "geo_replication_stats_can_failover",
				Description:   "A boolean flag which indicates whether or not account failover is supported for the account",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("AccountProperties.GeoReplicationStats.CanFailover"),
				IgnoreInTests: true,
			},
			{
				Name:          "failover_in_progress",
				Description:   "If the failover is in progress, the value will be true, otherwise, it will be null",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("AccountProperties.FailoverInProgress"),
				IgnoreInTests: true,
			},
			{
				Name:        "large_file_shares_state",
				Description: "Allow large file shares if sets to Enabled It cannot be disabled once it is enabled Possible values include: 'LargeFileSharesStateDisabled', 'LargeFileSharesStateEnabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.LargeFileSharesState"),
			},
			{
				Name:        "routing_preference_routing_choice",
				Description: "Routing Choice defines the kind of network routing opted by the user Possible values include: 'MicrosoftRouting', 'InternetRouting'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.RoutingPreference.RoutingChoice"),
			},
			{
				Name:          "routing_preference_publish_microsoft_endpoints",
				Description:   "A boolean flag which indicates whether microsoft routing storage endpoints are to be published",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("AccountProperties.RoutingPreference.PublishMicrosoftEndpoints"),
				IgnoreInTests: true,
			},
			{
				Name:          "routing_preference_publish_internet_endpoints",
				Description:   "A boolean flag which indicates whether internet routing storage endpoints are to be published",
				Type:          schema.TypeBool,
				Resolver:      schema.PathResolver("AccountProperties.RoutingPreference.PublishInternetEndpoints"),
				IgnoreInTests: true,
			},
			{
				Name:        "blob_restore_status",
				Description: "The status of blob restore progress Possible values are: - InProgress: Indicates that blob restore is ongoing - Complete: Indicates that blob restore has been completed successfully - Failed: Indicates that blob restore is failed Possible values include: 'InProgress', 'Complete', 'Failed'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.BlobRestoreStatus.Status"),
			},
			{
				Name:          "blob_restore_status_failure_reason",
				Description:   "Failure reason when blob restore is failed",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.BlobRestoreStatus.FailureReason"),
				IgnoreInTests: true,
			},
			{
				Name:          "blob_restore_status_restore_id",
				Description:   "Id for tracking blob restore request",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("AccountProperties.BlobRestoreStatus.RestoreID"),
				IgnoreInTests: true,
			},
			{
				Name:        "blob_restore_status_parameters_time_to_restore_time",
				Description: "Restore blob to the specified time",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("AccountProperties.BlobRestoreStatus.Parameters.TimeToRestore.Time"),
			},
			{
				Name:          "blob_restore_status_parameters_blob_ranges",
				Description:   "Blob ranges to restore",
				Type:          schema.TypeJSON,
				Resolver:      resolveStorageAccountBlobRestoreStatusParametersBlobRanges,
				IgnoreInTests: true,
			},
			{
				Name:        "allow_blob_public_access",
				Description: "Allow or disallow public access to all blobs or containers in the storage account The default interpretation is true for this property",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("AccountProperties.AllowBlobPublicAccess"),
			},
			{
				Name:        "minimum_tls_version",
				Description: "Set the minimum TLS version to be permitted on requests to storage The default interpretation is TLS 10 for this property Possible values include: 'TLS10', 'TLS11', 'TLS12'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AccountProperties.MinimumTLSVersion"),
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "location",
				Description: "The geo-location where the resource lives",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "The name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the resource Eg \"MicrosoftCompute/virtualMachines\" or \"MicrosoftStorage/storageAccounts\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "blob_logging_settings",
				Type:        schema.TypeJSON,
				Description: "BLOB service loggging settings (only for storage account types that support blobs)",
				Resolver:    fetchStorageAccountBlobLoggingSettings,
			},
			{
				Name:        "queue_logging_settings",
				Type:        schema.TypeJSON,
				Description: "Queue service loggging settings (only for storage account types that support queues)",
				Resolver:    fetchStorageAccountQueueLoggingSettings,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_storage_account_network_rule_set_virtual_network_rules",
				Description:   "VirtualNetworkRule virtual Network rule. ",
				Resolver:      fetchStorageAccountNetworkRuleSetVirtualNetworkRules,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique ID of azure_storage_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "virtual_network_resource_id",
						Description: "Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/MicrosoftNetwork/virtualNetworks/{vnetName}/subnets/{subnetName}",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualNetworkResourceID"),
					},
					{
						Name:        "action",
						Description: "The action of virtual network rule.",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "Gets the state of virtual network rule Possible values include: 'StateProvisioning', 'StateDeprovisioning', 'StateSucceeded', 'StateFailed', 'StateNetworkSourceDeleted'",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "azure_storage_account_network_rule_set_ip_rules",
				Description:   "IPRule IP rule with specific IP or IP range in CIDR format. ",
				Resolver:      fetchStorageAccountNetworkRuleSetIpRules,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique ID of azure_storage_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "ip_address_or_range",
						Description: "Specifies the IP or IP range in CIDR format Only IPV4 address is allowed",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IPAddressOrRange"),
					},
					{
						Name:        "action",
						Description: "The action of IP ACL rule Possible values include: 'Allow'",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "azure_storage_account_private_endpoint_connections",
				Description:   "Azure storage account private endpoint connection",
				Resolver:      fetchStorageAccountPrivateEndpointConnections,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "account_cq_id",
						Description: "Unique ID of azure_storage_accounts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "private_endpoint_id",
						Description: "The ARM identifier for Private Endpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateEndpoint.ID"),
					},
					{
						Name:        "private_link_service_connection_state_status",
						Description: "Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service Possible values include: 'Pending', 'Approved', 'Rejected'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "private_link_service_connection_state_description",
						Description: "The reason for approval/rejection of the connection",
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
						Description: "The provisioning state of the private endpoint connection resource Possible values include: 'PrivateEndpointConnectionProvisioningStateSucceeded', 'PrivateEndpointConnectionProvisioningStateCreating', 'PrivateEndpointConnectionProvisioningStateDeleting', 'PrivateEndpointConnectionProvisioningStateFailed'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PrivateEndpointConnectionProperties.ProvisioningState"),
					},
					{
						Name:        "id",
						Description: "Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "The name of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of the resource Eg \"MicrosoftCompute/virtualMachines\" or \"MicrosoftStorage/storageAccounts\"",
						Type:        schema.TypeString,
					},
				},
			},
			StorageContainers(),
			StorageBlobServices(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchStorageAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Storage.Accounts
	response, err := svc.List(ctx)
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
func resolveStorageAccountBlobRestoreStatusParametersBlobRanges(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	account := resource.Item.(storage.Account)
	if account.BlobRestoreStatus == nil || account.BlobRestoreStatus.Parameters == nil {
		return nil
	}
	data, err := json.Marshal(account.BlobRestoreStatus.Parameters.BlobRanges)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set("blob_restore_status_parameters_blob_ranges", data)
}
func fetchStorageAccountNetworkRuleSetVirtualNetworkRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account := parent.Item.(storage.Account)
	if account.NetworkRuleSet == nil || account.NetworkRuleSet.VirtualNetworkRules == nil {
		return nil
	}
	res <- *account.NetworkRuleSet.VirtualNetworkRules
	return nil
}
func fetchStorageAccountNetworkRuleSetIpRules(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account := parent.Item.(storage.Account)
	if account.NetworkRuleSet == nil || account.NetworkRuleSet.IPRules == nil {
		return nil
	}
	res <- *account.NetworkRuleSet.IPRules
	return nil
}
func fetchStorageAccountPrivateEndpointConnections(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	account := parent.Item.(storage.Account)
	if account.PrivateEndpointConnections == nil {
		return nil
	}
	res <- *account.PrivateEndpointConnections
	return nil
}

func fetchStorageAccountBlobLoggingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	acc, ok := resource.Item.(storage.Account)
	if !ok {
		return fmt.Errorf("not a storage.Account: %T", resource.Item)
	}

	if !isBlobSupported(&acc) {
		return nil
	}

	// fetch storageClient account keys for Shared Key authentication
	storageClient := meta.(*client.Client).Services().Storage
	details, err := client.ParseResourceID(*acc.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	keysResult, err := storageClient.Accounts.ListKeys(ctx, details.ResourceGroup, *acc.Name, "")
	if err != nil {
		if client.IgnoreAccessDenied(err) {
			meta.Logger().Warn("received access denied on Accounts.ListKeys", "resource_group", details.ResourceGroup, "account", *acc.Name, "err", err)
			return nil
		}
		return diag.WrapError(err)
	}
	if keysResult.Keys == nil || len(*keysResult.Keys) == 0 {
		return nil
	}

	// use account key to create a new authorizer and then fetch service properties
	auth, err := autorest.NewSharedKeyAuthorizer(*acc.Name, *(*keysResult.Keys)[0].Value, autorest.SharedKeyLite)
	if err != nil {
		return diag.WrapError(err)
	}
	blobProps := storageClient.NewBlobServiceProperties(auth)
	result, err := blobProps.GetServiceProperties(ctx, *acc.Name)
	if err != nil {
		// For premium 'page blob' storage accounts, we sometimes get "authorization error", not sure why.
		// In any case, we can probably ignore this since it only happens for premium 'page blob' storage accounts.
		if client.IgnoreAccessDenied(err) {
			meta.Logger().Warn("received access denied on GetServiceProperties", "resource_group", details.ResourceGroup, "account", *acc.Name, "err", err)
			return nil
		}
		return diag.WrapError(err)
	}
	var logging *accounts.Logging
	if result.StorageServiceProperties != nil {
		logging = result.StorageServiceProperties.Logging
	}
	data, err := json.Marshal(logging)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}

func fetchStorageAccountQueueLoggingSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	acc, ok := resource.Item.(storage.Account)
	if !ok {
		return fmt.Errorf("not a storage.Account: %T", resource.Item)
	}

	if !isQueueSupported(&acc) {
		return nil
	}

	// fetch storage account keys for Shared Key authentication
	storageClient := meta.(*client.Client).Services().Storage
	details, err := client.ParseResourceID(*acc.ID)
	if err != nil {
		return diag.WrapError(err)
	}
	keysResult, err := storageClient.Accounts.ListKeys(ctx, details.ResourceGroup, *acc.Name, "")
	if err != nil {
		if client.IgnoreAccessDenied(err) {
			meta.Logger().Warn("received access denied on Accounts.ListKeys", "resource_group", details.ResourceGroup, "account", *acc.Name, "err", err)
			return nil
		}
	}
	if keysResult.Keys == nil || len(*keysResult.Keys) == 0 {
		return nil
	}

	// use account key to create a new authorizer and then fetch service properties
	auth, err := autorest.NewSharedKeyAuthorizer(*acc.Name, *(*keysResult.Keys)[0].Value, autorest.SharedKeyLite)
	if err != nil {
		return diag.WrapError(err)
	}
	blobProps := storageClient.NewQueueServiceProperties(auth)
	result, err := blobProps.GetServiceProperties(ctx, *acc.Name)
	if err != nil {
		return diag.WrapError(err)
	}
	data, err := json.Marshal(result.Logging)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, data)
}

// isQueueSupported checks whether queues are supported for a storage account.
// Premium storage accounts don't support queues.
func isQueueSupported(account *storage.Account) bool {
	return account.Sku.Tier == storage.Standard && account.Kind == storage.StorageV2
}

// isBlobSupported checks whether blobs are supported for a storage account.
func isBlobSupported(account *storage.Account) bool {
	return (account.Kind == storage.Storage) || (account.Kind == storage.StorageV2) ||
		(account.Kind == storage.BlockBlobStorage) || (account.Kind == storage.BlobStorage)
}
