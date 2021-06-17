
# Table: azure_mysql_servers
Azure mysql server
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|identity_principal_id|uuid|The Azure Active Directory principal id|
|identity_type|text|The identity type Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource Possible values include: 'SystemAssigned'|
|identity_tenant_id|uuid|The Azure Active Directory tenant id|
|sku_name|text|The name of the sku, typically, tier + family + cores, eg B_Gen4_1, GP_Gen5_8|
|sku_tier|text|The tier of the particular SKU, eg Basic Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'|
|sku_capacity|integer|The scale up/out capacity, representing server's compute units|
|sku_size|text|The size code, to be interpreted by resource as appropriate|
|sku_family|text|The family of hardware|
|administrator_login|text|The administrator's login name of a server Can only be specified when the server is being created (and is required for creation)|
|version|text|Server version Possible values include: 'FiveFullStopSix', 'FiveFullStopSeven', 'EightFullStopZero'|
|ssl_enforcement|text|Enable ssl enforcement or not when connect to server Possible values include: 'SslEnforcementEnumEnabled', 'SslEnforcementEnumDisabled'|
|minimal_tls_version|text|Enforce a minimal Tls version for the server Possible values include: 'TLS10', 'TLS11', 'TLS12', 'TLSEnforcementDisabled'|
|byok_enforcement|text|Status showing whether the server data encryption is enabled with customer-managed keys|
|infrastructure_encryption|text|Status showing whether the server enabled infrastructure encryption Possible values include: 'InfrastructureEncryptionEnabled', 'InfrastructureEncryptionDisabled'|
|user_visible_state|text|A state of a server that is visible to user Possible values include: 'ServerStateReady', 'ServerStateDropping', 'ServerStateDisabled', 'ServerStateInaccessible'|
|fully_qualified_domain_name|text|The fully qualified domain name of a server|
|earliest_restore_date_time|timestamp without time zone|Earliest restore point creation time (ISO8601 format)|
|storage_profile_backup_retention_days|integer|Backup retention days for the server|
|storage_profile_geo_redundant_backup|text|Enable Geo-redundant or not for server backup Possible values include: 'Enabled', 'Disabled'|
|storage_profile_storage_mb|integer|Max storage allowed for a server|
|storage_profile_storage_autogrow|text|Enable Storage Auto Grow Possible values include: 'StorageAutogrowEnabled', 'StorageAutogrowDisabled'|
|replication_role|text|The replication role of the server|
|master_server_id|text|The master server id of a replica server|
|replica_capacity|integer|The maximum number of replicas that a master server can have|
|public_network_access|text|Whether or not public network access is allowed for this server Value is optional but if passed in, must be 'Enabled' or 'Disabled' Possible values include: 'PublicNetworkAccessEnumEnabled', 'PublicNetworkAccessEnumDisabled'|
|tags|jsonb|Resource tags|
|location|text|The geo-location where the resource lives|
|resource_id|text|Fully qualified resource ID for the resource Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}|
|name|text|The name of the resource|
|type|text|The type of the resource Eg "MicrosoftCompute/virtualMachines" or "MicrosoftStorage/storageAccounts"|
