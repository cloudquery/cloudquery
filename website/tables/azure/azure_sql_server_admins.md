# Table: azure_sql_server_admins

This table shows data for Azure SQL Server Admins.

https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/server-azure-ad-administrators/list-by-server?tabs=HTTP#serverazureadadministrator

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
|name|utf8|
|type|utf8|