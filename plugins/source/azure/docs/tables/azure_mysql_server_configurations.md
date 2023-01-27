# Table: azure_mysql_server_configurations

https://learn.microsoft.com/en-us/rest/api/mysql/singleserver/configurations/list-by-server?tabs=HTTP#configuration

The primary key for this table is **id**.

## Relations

This table depends on [azure_mysql_servers](azure_mysql_servers.md).

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