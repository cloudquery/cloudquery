
# Table: azure_batch_accounts
Account contains information about an Azure Batch account
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|account_endpoint|text|The account endpoint used to interact with the Batch service|
|provisioning_state|text|The provisioned state of the resource|
|pool_allocation_mode|text|Possible values include: 'PoolAllocationModeBatchService', 'PoolAllocationModeUserSubscription'|
|key_vault_reference_id|text|The resource ID of the Azure key vault associated with the Batch account|
|key_vault_reference_url|text|The URL of the Azure key vault associated with the Batch account|
|public_network_access|text|If not specified, the default value is 'enabled'. Possible values include: 'PublicNetworkAccessTypeEnabled', 'PublicNetworkAccessTypeDisabled'|
|auto_storage_last_key_sync_time|timestamp without time zone|The UTC time at which storage keys were last synchronized with the Batch account.|
|auto_storage_storage_account_id|text|The resource ID of the storage account to be used for auto-storage account|
|auto_storage_authentication_mode|text|The authentication mode which the Batch service will use to manage the auto-storage account|
|auto_storage_node_identity_reference_resource_id|text|The ARM resource id of the user assigned identity|
|encryption_key_source|text|Type of the key source. Possible values include: 'KeySourceMicrosoftBatch', 'KeySourceMicrosoftKeyVault'|
|encryption_key_vault_properties_key_identifier|text|Full path to the versioned secret|
|dedicated_core_quota|integer|For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned|
|low_priority_core_quota|integer|For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not returned|
|dedicated_core_quota_per_vm_family|jsonb|A list of the dedicated core quota per Virtual Machine family for the Batch account|
|dedicated_core_quota_per_vm_family_enforced|boolean|Batch is transitioning its core quota system for dedicated cores to be enforced per Virtual Machine family|
|pool_quota|integer|The pool quota for the Batch account.|
|active_job_and_job_schedule_quota|integer|The active job and job schedule quota for the Batch account.|
|allowed_authentication_modes|text[]|List of allowed authentication modes for the Batch account that can be used to authenticate with the data plane|
|identity_principal_id|text|The principal id of the Batch account|
|identity_tenant_id|text|The tenant id associated with the Batch account|
|identity_type|text|The type of identity used for the Batch account|
|identity_user_assigned_identities|jsonb|The list of user identities associated with the Batch account|
|id|text|The ID of the resource|
|name|text|The name of the resource|
|type|text|The type of the resource|
|location|text|The location of the resource|
|tags|jsonb|The tags of the resource|
