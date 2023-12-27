# Table: azure_postgresqlflexibleservers_server_configurations

This table shows data for Azure PostgreSQL Flexible Servers Server Configurations.

https://learn.microsoft.com/en-us/rest/api/postgresql/flexibleserver/configurations/list-by-server?view=rest-postgresql-flexibleserver-2022-12-01&tabs=HTTP

The primary key for this table is **id**.

## Relations

This table depends on [azure_postgresqlflexibleservers_servers](azure_postgresqlflexibleservers_servers.md).

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