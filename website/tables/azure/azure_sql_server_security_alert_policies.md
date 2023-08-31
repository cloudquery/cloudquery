# Table: azure_sql_server_security_alert_policies

This table shows data for Azure SQL Server Security Alert Policies.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/server-security-alert-policies/list-by-server?tabs=HTTP#serversecurityalertpolicy

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|