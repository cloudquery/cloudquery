# Table: azure_mariadb_server_configurations

https://learn.microsoft.com/en-us/rest/api/mariadb/configurations/list-by-server?tabs=HTTP#configuration

The primary key for this table is **id**.

## Relations

This table depends on [azure_mariadb_servers](azure_mariadb_servers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|properties|JSON|
|name|String|
|type|String|