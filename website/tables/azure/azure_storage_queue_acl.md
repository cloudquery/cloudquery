# Table: azure_storage_queue_acl

This table shows data for Azure Storage Queue ACL.

https://learn.microsoft.com/en-us/rest/api/storageservices/get-queue-acl#response-body

The primary key for this table is **queue_id**.

## Relations

This table depends on [azure_storage_queues](azure_storage_queues).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|queue_id (PK)|`utf8`|
|signed_identifiers|`json`|
|version|`utf8`|