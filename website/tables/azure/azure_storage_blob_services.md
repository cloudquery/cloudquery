# Table: azure_storage_blob_services

This table shows data for Azure Storage Blob Services.

https://learn.microsoft.com/en-us/rest/api/storagerp/blob-services/list?tabs=HTTP#blobserviceproperties

The primary key for this table is **id**.

## Relations

This table depends on [azure_storage_accounts](azure_storage_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|sku|`json`|
|type|`utf8`|