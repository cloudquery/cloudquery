# Table: azure_storage_queues

This table shows data for Azure Storage Queues.

https://learn.microsoft.com/en-us/rest/api/storagerp/queue/list?tabs=HTTP#listqueue

The primary key for this table is **id**.

## Relations

This table depends on [azure_storage_accounts](azure_storage_accounts).

The following tables depend on azure_storage_queues:
  - [azure_storage_queue_acl](azure_storage_queue_acl)

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