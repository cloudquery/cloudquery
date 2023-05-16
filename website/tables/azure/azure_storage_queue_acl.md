# Table: azure_storage_queue_acl

This table shows data for Azure Storage Queue ACL.

https://learn.microsoft.com/en-us/rest/api/storageservices/get-queue-acl#response-body

The primary key for this table is **queue_id**.

## Relations

This table depends on [azure_storage_queues](azure_storage_queues).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|queue_id (PK)|String|
|signed_identifiers|JSON|
|version|String|