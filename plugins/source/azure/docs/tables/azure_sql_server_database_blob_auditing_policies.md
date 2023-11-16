# Table: azure_sql_server_database_blob_auditing_policies

This table shows data for Azure SQL Server Database Blob Auditing Policies.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/database-blob-auditing-policies/list-by-database?tabs=HTTP#databaseblobauditingpolicy

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_server_databases](azure_sql_server_databases.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|kind|`utf8`|
|name|`utf8`|
|type|`utf8`|