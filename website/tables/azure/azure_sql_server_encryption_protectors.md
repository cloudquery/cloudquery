# Table: azure_sql_server_encryption_protectors

This table shows data for Azure SQL Server Encryption Protectors.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/encryption-protectors/list-by-server?tabs=HTTP#encryptionprotector

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|subscription_id|utf8|
|properties|json|
|id (PK)|utf8|
|kind|utf8|
|location|utf8|
|name|utf8|
|type|utf8|