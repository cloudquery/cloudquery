# Table: azure_sql_managed_instance_encryption_protectors

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/managed-instance-encryption-protectors/list-by-instance?tabs=HTTP#managedinstanceencryptionprotector

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_managed_instances](azure_sql_managed_instances.md).

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
|kind|String|
|name|String|
|type|String|