# Table: azure_sql_managed_server_security_alert_policies

This table shows data for Azure SQL Managed Server Security Alert Policies.

https://learn.microsoft.com/en-us/rest/api/sql/managed-server-security-alert-policies/list-by-instance?view=rest-sql-2021-11-01&tabs=HTTP

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_managed_instances](azure_sql_managed_instances.md).

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