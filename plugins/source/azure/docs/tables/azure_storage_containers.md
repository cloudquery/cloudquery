# Table: azure_storage_containers

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage#ListContainerItem

The primary key for this table is **id**.

## Relations
This table depends on [azure_storage_accounts](azure_storage_accounts.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|storage_account_id|String|
|version|String|
|deleted|Bool|
|deleted_time|Timestamp|
|remaining_retention_days|Int|
|default_encryption_scope|String|
|deny_encryption_scope_override|Bool|
|public_access|String|
|last_modified_time|Timestamp|
|lease_status|String|
|lease_state|String|
|lease_duration|String|
|metadata|JSON|
|immutability_policy|JSON|
|legal_hold|JSON|
|has_legal_hold|Bool|
|has_immutability_policy|Bool|
|etag|String|
|id (PK)|String|
|name|String|
|type|String|