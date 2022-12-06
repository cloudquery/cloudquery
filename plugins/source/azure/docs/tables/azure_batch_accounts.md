# Table: azure_batch_accounts

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch#Account

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|account_endpoint|String|
|provisioning_state|String|
|pool_allocation_mode|String|
|key_vault_reference|JSON|
|public_network_access|String|
|private_endpoint_connections|JSON|
|auto_storage|JSON|
|encryption|JSON|
|dedicated_core_quota|Int|
|low_priority_core_quota|Int|
|dedicated_core_quota_per_vm_family|JSON|
|dedicated_core_quota_per_vm_family_enforced|Bool|
|pool_quota|Int|
|active_job_and_job_schedule_quota|Int|
|allowed_authentication_modes|StringArray|
|identity|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|