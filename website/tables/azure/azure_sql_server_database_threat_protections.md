# Table: azure_sql_server_database_threat_protections

This table shows data for Azure SQL Server Database Threat Protections.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/database-advanced-threat-protection-settings/list-by-database?tabs=HTTP#databaseadvancedthreatprotection

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_server_databases](azure_sql_server_databases).

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