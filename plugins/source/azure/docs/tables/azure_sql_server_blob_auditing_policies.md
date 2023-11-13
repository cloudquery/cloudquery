# Table: azure_sql_server_blob_auditing_policies

This table shows data for Azure SQL Server Blob Auditing Policies.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/server-blob-auditing-policies/list-by-server?tabs=HTTP#serverblobauditingpolicy

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers.md).

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