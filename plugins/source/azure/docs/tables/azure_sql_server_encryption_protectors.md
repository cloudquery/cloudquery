# Table: azure_sql_server_encryption_protectors

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/encryption-protectors/list-by-server?tabs=HTTP#encryptionprotector

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers.md).

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
|location|String|
|name|String|
|type|String|