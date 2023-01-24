# Table: azure_sql_server_database_threat_protections

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/database-advanced-threat-protection-settings/list-by-database?tabs=HTTP#databaseadvancedthreatprotection

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_server_databases](azure_sql_server_databases.md).

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
|system_data|JSON|
|name|String|
|type|String|