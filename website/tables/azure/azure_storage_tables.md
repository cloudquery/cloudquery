# Table: azure_storage_tables

This table shows data for Azure Storage Tables.

https://learn.microsoft.com/en-us/rest/api/storagerp/table/list?tabs=HTTP#table

The primary key for this table is **id**.

## Relations

This table depends on [azure_storage_accounts](azure_storage_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|