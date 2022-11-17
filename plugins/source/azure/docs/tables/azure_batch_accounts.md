# Table: azure_batch_accounts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch#Account

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|identity|JSON|
|network_profile|JSON|
|public_network_access|String|
|account_endpoint|String|
|active_job_and_job_schedule_quota|Int|
|allowed_authentication_modes|StringArray|
|auto_storage|JSON|
|dedicated_core_quota|Int|
|dedicated_core_quota_per_vm_family|JSON|
|dedicated_core_quota_per_vm_family_enforced|Bool|
|encryption|JSON|
|key_vault_reference|JSON|
|low_priority_core_quota|Int|
|node_management_endpoint|String|
|pool_allocation_mode|String|
|pool_quota|Int|
|private_endpoint_connections|JSON|
|provisioning_state|String|
|id (PK)|String|
|location|String|
|name|String|
|tags|JSON|
|type|String|