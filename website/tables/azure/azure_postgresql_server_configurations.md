# Table: azure_postgresql_server_configurations

https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/configurations/list-by-server?tabs=HTTP#configuration

The primary key for this table is **id**.

## Relations

This table depends on [azure_postgresql_servers](azure_postgresql_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|properties|JSON|
|id (PK)|String|
|name|String|
|type|String|