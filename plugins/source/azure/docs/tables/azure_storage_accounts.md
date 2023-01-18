# Table: azure_storage_accounts

https://learn.microsoft.com/en-us/rest/api/storagerp/storage-accounts/list?tabs=HTTP#storageaccount

The primary key for this table is **id**.

## Relations

The following tables depend on azure_storage_accounts:
  - [azure_storage_blob_services](azure_storage_blob_services.md)
  - [azure_storage_containers](azure_storage_containers.md)
  - [azure_storage_tables](azure_storage_tables.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|properties|JSON|
|tags|JSON|
|kind|String|
|name|String|
|sku|JSON|
|type|String|