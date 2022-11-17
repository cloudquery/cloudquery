# Table: azure_storage_blob_services

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage#BlobServiceProperties

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
|automatic_snapshot_policy_enabled|Bool|
|change_feed|JSON|
|container_delete_retention_policy|JSON|
|cors|JSON|
|default_service_version|String|
|delete_retention_policy|JSON|
|is_versioning_enabled|Bool|
|last_access_time_tracking_policy|JSON|
|restore_policy|JSON|
|id (PK)|String|
|name|String|
|sku|JSON|
|type|String|
|account_id|String|