# Table: azure_storage_accounts

This table shows data for Azure Storage Accounts.

https://learn.microsoft.com/en-us/rest/api/storagerp/storage-accounts/list?tabs=HTTP#storageaccount

The primary key for this table is **id**.

## Relations

The following tables depend on azure_storage_accounts:
  - [azure_storage_blob_services](azure_storage_blob_services)
  - [azure_storage_containers](azure_storage_containers)
  - [azure_storage_file_shares](azure_storage_file_shares)
  - [azure_storage_queue_services](azure_storage_queue_services)
  - [azure_storage_queues](azure_storage_queues)
  - [azure_storage_tables](azure_storage_tables)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|subscription_id|utf8|
|location|utf8|
|extended_location|json|
|identity|json|
|properties|json|
|tags|json|
|id (PK)|utf8|
|kind|utf8|
|name|utf8|
|sku|json|
|type|utf8|