# Table: azure_sql_server_security_alert_policies

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/server-security-alert-policies/list-by-server?tabs=HTTP#serversecurityalertpolicy

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_servers](azure_sql_servers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|properties|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|