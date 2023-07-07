# Table: azure_storage_queue_services

This table shows data for Azure Storage Queue Services.

https://learn.microsoft.com/en-us/rest/api/storagerp/queue-services/list?tabs=HTTP#queueserviceproperties

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