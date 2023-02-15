# Table: azure_postgresql_databases

https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/databases/list-by-server?tabs=HTTP#database

The primary key for this table is **id**.

## Relations

This table depends on [azure_postgresql_servers](azure_postgresql_servers.md).

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
|name|String|
|type|String|