# Table: azure_mysql_server_configurations

This table shows data for Azure MySQL Server Configurations.

https://learn.microsoft.com/en-us/rest/api/mysql/singleserver/configurations/list-by-server?tabs=HTTP#configuration

The primary key for this table is **id**.

## Relations

This table depends on [azure_mysql_servers](azure_mysql_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|