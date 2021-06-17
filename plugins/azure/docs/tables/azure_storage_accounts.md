
# Table: azure_storage_accounts
Azure storage account
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|sku_name|text|Possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS', 'PremiumZRS', 'StandardGZRS', 'StandardRAGZRS'|
|sku_tier|text|Possible values include: 'Standard', 'Premium'|
|kind|text|Gets the Kind Possible values include: 'Storage', 'StorageV2', 'BlobStorage', 'FileStorage', 'BlockBlobStorage'|
|identity_principal_id|text|The principal ID of resource identity|
|identity_tenant_id|text|The tenant ID of resource|
|identity_type|text|The identity type|
|provisioning_state|text|Gets the status of the storage account at the time the operation was called Possible values include: 'Creating', 'ResolvingDNS', 'Succeeded'|
|primary_endpoints_blob|text|Gets the blob endpoint|
|primary_endpoints_queue|text|Gets the queue endpoint|
|primary_endpoints_table|text|Gets the table endpoint|
|primary_endpoints_file|text|Gets the file endpoint|
|primary_endpoints_web|text|Gets the web endpoint|
|primary_endpoints_dfs|text|Gets the dfs endpoint|
|primary_endpoints_microsoft_endpoints_blob|text|Gets the blob endpoint|
|primary_endpoints_microsoft_endpoints_queue|text|Gets the queue endpoint|
|primary_endpoints_microsoft_endpoints_table|text|Gets the table endpoint|
|primary_endpoints_microsoft_endpoints_file|text|Gets the file endpoint|
|primary_endpoints_microsoft_endpoints_web|text|Gets the web endpoint|
|primary_endpoints_microsoft_endpoints_dfs|text|Gets the dfs endpoint|
|primary_endpoints_internet_endpoints_blob|text|Gets the blob endpoint|
|primary_endpoints_internet_endpoints_file|text|Gets the file endpoint|
|primary_endpoints_internet_endpoints_web|text|Gets the web endpoint|
|primary_endpoints_internet_endpoints_dfs|text|Gets the dfs endpoint|
|primary_location|text|Gets the location of the primary data center for the storage account|
|status_of_primary|text|Gets the status indicating whether the primary location of the storage account is available or unavailable Possible values include: 'Available', 'Unavailable'|
|last_geo_failover_time|timestamp without time zone|Gets the timestamp of the most recent instance of a failover to the secondary location Only the most recent timestamp is retained This element is not returned if there has never been a failover instance Only available if the accountType is Standard_GRS or Standard_RAGRS|
|secondary_location|text|Gets the location of the geo-replicated secondary for the storage account Only available if the accountType is Standard_GRS or Standard_RAGRS|
|status_of_secondary|text|Gets the status indicating whether the secondary location of the storage account is available or unavailable Only available if the SKU name is Standard_GRS or Standard_RAGRS Possible values include: 'Available', 'Unavailable'|
|creation_time|timestamp without time zone|Gets the creation date and time of the storage account in UTC|
|custom_domain_name|text|Gets or sets the custom domain name assigned to the storage account Name is the CNAME source|
|custom_domain_use_sub_domain_name|boolean|Indicates whether indirect CName validation is enabled Default value is false This should only be set on updates|
|secondary_endpoints_blob|text|Gets the blob endpoint|
|secondary_endpoints_queue|text|Gets the queue endpoint|
|secondary_endpoints_table|text|Gets the table endpoint|
|secondary_endpoints_file|text|Gets the file endpoint|
|secondary_endpoints_web|text|Gets the web endpoint|
|secondary_endpoints_dfs|text|Gets the dfs endpoint|
|secondary_endpoints_microsoft_endpoints_blob|text|Gets the blob endpoint|
|secondary_endpoints_microsoft_endpoints_queue|text|Gets the queue endpoint|
|secondary_endpoints_microsoft_endpoints_table|text|Gets the table endpoint|
|secondary_endpoints_microsoft_endpoints_file|text|Gets the file endpoint|
|secondary_endpoints_microsoft_endpoints_web|text|Gets the web endpoint|
|secondary_endpoints_microsoft_endpoints_dfs|text|Gets the dfs endpoint|
|secondary_endpoints_internet_endpoints_blob|text|Gets the blob endpoint|
|secondary_endpoints_internet_endpoints_file|text|Gets the file endpoint|
|secondary_endpoints_internet_endpoints_web|text|Gets the web endpoint|
|secondary_endpoints_internet_endpoints_dfs|text|Gets the dfs endpoint|
|encryption_services_blob_enabled|boolean|A boolean indicating whether or not the service encrypts the data as it is stored|
|encryption_services_blob_last_enabled_time|timestamp without time zone|Gets a rough estimate of the date/time when the encryption was last enabled by the user Only returned when encryption is enabled There might be some unencrypted blobs which were written after this time, as it is just a rough estimate|
|encryption_services_blob_key_type|text|Encryption key type to be used for the encryption service 'Account' key type implies that an account-scoped encryption key will be used 'Service' key type implies that a default service key is used Possible values include: 'KeyTypeService', 'KeyTypeAccount'|
|encryption_services_file_enabled|boolean|A boolean indicating whether or not the service encrypts the data as it is stored|
|encryption_services_file_last_enabled_time|timestamp without time zone|Gets a rough estimate of the date/time when the encryption was last enabled by the user Only returned when encryption is enabled There might be some unencrypted blobs which were written after this time, as it is just a rough estimate|
|encryption_services_file_key_type|text|Encryption key type to be used for the encryption service 'Account' key type implies that an account-scoped encryption key will be used 'Service' key type implies that a default service key is used Possible values include: 'KeyTypeService', 'KeyTypeAccount'|
|encryption_services_table_enabled|boolean|A boolean indicating whether or not the service encrypts the data as it is stored|
|encryption_services_table_last_enabled_time|timestamp without time zone|Gets a rough estimate of the date/time when the encryption was last enabled by the user Only returned when encryption is enabled There might be some unencrypted blobs which were written after this time, as it is just a rough estimate|
|encryption_services_table_key_type|text|Encryption key type to be used for the encryption service 'Account' key type implies that an account-scoped encryption key will be used 'Service' key type implies that a default service key is used Possible values include: 'KeyTypeService', 'KeyTypeAccount'|
|encryption_services_queue_enabled|boolean|A boolean indicating whether or not the service encrypts the data as it is stored|
|encryption_services_queue_last_enabled_time|timestamp without time zone|Gets a rough estimate of the date/time when the encryption was last enabled by the user Only returned when encryption is enabled There might be some unencrypted blobs which were written after this time, as it is just a rough estimate|
|encryption_services_queue_key_type|text|Encryption key type to be used for the encryption service 'Account' key type implies that an account-scoped encryption key will be used 'Service' key type implies that a default service key is used Possible values include: 'KeyTypeService', 'KeyTypeAccount'|
|encryption_key_source|text|The encryption keySource (provider) Possible values (case-insensitive):  MicrosoftStorage, MicrosoftKeyvault Possible values include: 'KeySourceMicrosoftStorage', 'KeySourceMicrosoftKeyvault'|
|encryption_require_infrastructure_encryption|boolean|A boolean indicating whether or not the service applies a secondary layer of encryption with platform managed keys for data at rest|
|encryption_key_vault_properties_key_name|text|The name of KeyVault key|
|encryption_key_vault_properties_key_version|text|The version of KeyVault key|
|encryption_key_vault_properties_key_vault_uri|text|The Uri of KeyVault|
|encryption_key_current_versioned_key_identifier|text|The object identifier of the current versioned Key Vault Key in use|
|encryption_key_last_key_rotation_timestamp_time|timestamp without time zone|Timestamp of last rotation of the Key Vault Key|
|access_tier|text|Required for storage accounts where kind = BlobStorage The access tier used for billing Possible values include: 'Hot', 'Cool'|
|files_identity_auth_directory_service_options|text|Indicates the directory service used Possible values include: 'DirectoryServiceOptionsNone', 'DirectoryServiceOptionsAADDS', 'DirectoryServiceOptionsAD'|
|files_identity_auth_ad_properties_domain_name|text|Specifies the primary domain that the AD DNS server is authoritative for|
|files_identity_auth_ad_properties_net_bios_domain_name|text|Specifies the NetBIOS domain name|
|files_identity_auth_ad_properties_forest_name|text|Specifies the Active Directory forest to get|
|files_identity_auth_ad_properties_domain_guid|text|Specifies the domain GUID|
|files_identity_auth_ad_properties_net_bios_domain_sid|text|Specifies the security identifier (SID)|
|files_identity_auth_ad_properties_azure_storage_sid|text|Specifies the security identifier (SID) for Azure Storage|
|enable_https_traffic_only|boolean|Allows https traffic only to storage service if sets to true|
|network_rule_set_bypass|text|Specifies whether traffic is bypassed for Logging/Metrics/AzureServices Possible values are any combination of Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics Possible values include: 'None', 'Logging', 'Metrics', 'AzureServices'|
|network_rule_set_default_action|text|Specifies the default action of allow or deny when no other rules match Possible values include: 'DefaultActionAllow', 'DefaultActionDeny'|
|is_hns_enabled|boolean|Account HierarchicalNamespace enabled if sets to true|
|geo_replication_stats_status|text|The status of the secondary location Possible values are: - Live: Indicates that the secondary location is active and operational - Bootstrap: Indicates initial synchronization from the primary location to the secondary location is in progressThis typically occurs when replication is first enabled - Unavailable: Indicates that the secondary location is temporarily unavailable Possible values include: 'GeoReplicationStatusLive', 'GeoReplicationStatusBootstrap', 'GeoReplicationStatusUnavailable'|
|geo_replication_stats_last_sync_time|timestamp without time zone|All primary writes preceding this UTC date/time value are guaranteed to be available for read operations Primary writes following this point in time may or may not be available for reads Element may be default value if value of LastSyncTime is not available, this can happen if secondary is offline or we are in bootstrap|
|geo_replication_stats_can_failover|boolean|A boolean flag which indicates whether or not account failover is supported for the account|
|failover_in_progress|boolean|If the failover is in progress, the value will be true, otherwise, it will be null|
|large_file_shares_state|text|Allow large file shares if sets to Enabled It cannot be disabled once it is enabled Possible values include: 'LargeFileSharesStateDisabled', 'LargeFileSharesStateEnabled'|
|routing_preference_routing_choice|text|Routing Choice defines the kind of network routing opted by the user Possible values include: 'MicrosoftRouting', 'InternetRouting'|
|routing_preference_publish_microsoft_endpoints|boolean|A boolean flag which indicates whether microsoft routing storage endpoints are to be published|
|routing_preference_publish_internet_endpoints|boolean|A boolean flag which indicates whether internet routing storage endpoints are to be published|
|blob_restore_status|text|The status of blob restore progress Possible values are: - InProgress: Indicates that blob restore is ongoing - Complete: Indicates that blob restore has been completed successfully - Failed: Indicates that blob restore is failed Possible values include: 'InProgress', 'Complete', 'Failed'|
|blob_restore_status_failure_reason|text|Failure reason when blob restore is failed|
|blob_restore_status_restore_id|text|Id for tracking blob restore request|
|blob_restore_status_parameters_time_to_restore_time|timestamp without time zone|Restore blob to the specified time|
|blob_restore_status_parameters_blob_ranges|jsonb|Blob ranges to restore|
|allow_blob_public_access|boolean|Allow or disallow public access to all blobs or containers in the storage account The default interpretation is true for this property|
|minimum_tls_version|text|Set the minimum TLS version to be permitted on requests to storage The default interpretation is TLS 10 for this property Possible values include: 'TLS10', 'TLS11', 'TLS12'|
|tags|jsonb|Resource tags|
|location|text|The geo-location where the resource lives|
|resource_id|text|Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}|
|name|text|The name of the resource|
|type|text|The type of the resource Eg "MicrosoftCompute/virtualMachines" or "MicrosoftStorage/storageAccounts"|
