# Table: azure_storage_blob_services

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage#BlobServiceProperties

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
|cors|JSON|
|default_service_version|String|
|delete_retention_policy|JSON|
|is_versioning_enabled|Bool|
|automatic_snapshot_policy_enabled|Bool|
|change_feed|JSON|
|restore_policy|JSON|
|container_delete_retention_policy|JSON|
|last_access_time_tracking_policy|JSON|
|sku|JSON|
|id (PK)|String|
|name|String|
|type|String|