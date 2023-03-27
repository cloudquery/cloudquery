# Table: azure_storage_queue_services

This table shows data for Azure Storage Queue Services.

https://learn.microsoft.com/en-us/rest/api/storagerp/queue-services/list?tabs=HTTP#queueserviceproperties

The primary key for this table is **id**.

## Relations

This table depends on [azure_storage_accounts](azure_storage_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|