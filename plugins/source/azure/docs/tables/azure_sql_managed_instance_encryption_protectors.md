# Table: azure_sql_managed_instance_encryption_protectors

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
|id (PK)|String|
|properties|JSON|
|kind|String|
|name|String|
|type|String|