# Table: azure_mysql_servers

This table shows data for Azure MySQL Servers.

https://learn.microsoft.com/en-us/rest/api/mysql/singleserver/servers(2017-12-01)/list?tabs=HTTP#server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_mysql_servers:
  - [azure_mysql_server_configurations](azure_mysql_server_configurations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|subscription_id|utf8|
|location|utf8|
|identity|json|
|properties|json|
|sku|json|
|tags|json|
|id (PK)|utf8|
|name|utf8|
|type|utf8|