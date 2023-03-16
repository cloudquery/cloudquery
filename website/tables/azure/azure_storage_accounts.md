# Table: azure_storage_accounts

This table shows data for Azure Storage Accounts.

https://learn.microsoft.com/en-us/rest/api/storagerp/storage-accounts/list?tabs=HTTP#storageaccount

The primary key for this table is **id**.

## Relations

The following tables depend on azure_storage_accounts:
  - [azure_storage_blob_services](azure_storage_blob_services)
  - [azure_storage_containers](azure_storage_containers)
  - [azure_storage_tables](azure_storage_tables)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|extended_location|JSON|
|identity|JSON|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|kind|String|
|name|String|
|sku|JSON|
|type|String|