# Table: azure_storage_containers

https://learn.microsoft.com/en-us/rest/api/storagerp/blob-containers/list?tabs=HTTP#listcontaineritem

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
|id (PK)|String|
|properties|JSON|
|etag|String|
|name|String|
|type|String|