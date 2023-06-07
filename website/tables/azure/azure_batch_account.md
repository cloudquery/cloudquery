# Table: azure_batch_account

This table shows data for Azure Batch Account.

https://learn.microsoft.com/en-us/rest/api/batchmanagement/batch-account/list?tabs=HTTP#batchaccount

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|identity|`json`|
|properties|`json`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|tags|`json`|
|type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.


