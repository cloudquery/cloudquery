# Table: azure_sql_managed_instance_encryption_protectors

This table shows data for Azure SQL Managed Instance Encryption Protectors.

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/managed-instance-encryption-protectors/list-by-instance?tabs=HTTP#managedinstanceencryptionprotector

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_managed_instances](azure_sql_managed_instances).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|kind|`utf8`|
|name|`utf8`|
|type|`utf8`|