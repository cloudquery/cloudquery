# Table: azure_mariadb_server_configurations

This table shows data for Azure MariaDB Server Configurations.

https://learn.microsoft.com/en-us/rest/api/mariadb/configurations/list-by-server?tabs=HTTP#configuration

The primary key for this table is **id**.

## Relations

This table depends on [azure_mariadb_servers](azure_mariadb_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|