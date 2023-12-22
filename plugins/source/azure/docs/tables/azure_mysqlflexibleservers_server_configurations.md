# Table: azure_mysqlflexibleservers_server_configurations

This table shows data for Azure Mysqlflexibleservers Server Configurations.

https://learn.microsoft.com/en-us/rest/api/mysql/flexibleserver/configurations/list-by-server

The primary key for this table is **id**.

## Relations

This table depends on [azure_mysqlflexibleservers_servers](azure_mysqlflexibleservers_servers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|