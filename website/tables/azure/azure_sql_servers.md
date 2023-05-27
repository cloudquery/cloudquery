# Table: azure_sql_servers

This table shows data for Azure SQL Servers.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/servers/list?tabs=HTTP#server

The primary key for this table is **id**.

## Relations

The following tables depend on azure_sql_servers:
  - [azure_sql_server_admins](azure_sql_server_admins)
  - [azure_sql_server_advanced_threat_protection_settings](azure_sql_server_advanced_threat_protection_settings)
  - [azure_sql_server_blob_auditing_policies](azure_sql_server_blob_auditing_policies)
  - [azure_sql_server_databases](azure_sql_server_databases)
  - [azure_sql_server_encryption_protectors](azure_sql_server_encryption_protectors)
  - [azure_sql_server_security_alert_policies](azure_sql_server_security_alert_policies)
  - [azure_sql_server_virtual_network_rules](azure_sql_server_virtual_network_rules)
  - [azure_sql_server_vulnerability_assessments](azure_sql_server_vulnerability_assessments)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|kind|`utf8`|
|name|`utf8`|
|type|`utf8`|