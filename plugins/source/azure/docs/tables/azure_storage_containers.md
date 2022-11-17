# Table: azure_storage_containers

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage#ListContainerItem

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
|default_encryption_scope|String|
|deny_encryption_scope_override|Bool|
|enable_nfs_v3_all_squash|Bool|
|enable_nfs_v3_root_squash|Bool|
|immutable_storage_with_versioning|JSON|
|metadata|JSON|
|public_access|String|
|deleted|Bool|
|deleted_time|Timestamp|
|has_immutability_policy|Bool|
|has_legal_hold|Bool|
|immutability_policy|JSON|
|last_modified_time|Timestamp|
|lease_duration|String|
|lease_state|String|
|lease_status|String|
|legal_hold|JSON|
|remaining_retention_days|Int|
|version|String|
|etag|String|
|id (PK)|String|
|name|String|
|type|String|
|account_id|String|