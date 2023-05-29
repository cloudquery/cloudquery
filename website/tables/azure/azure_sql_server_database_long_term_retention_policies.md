# Table: azure_sql_server_database_long_term_retention_policies

This table shows data for Azure SQL Server Database Long Term Retention Policies.

https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/long-term-retention-policies/list-by-database?tabs=HTTP#longtermretentionpolicy

The primary key for this table is **id**.

## Relations

This table depends on [azure_sql_server_databases](azure_sql_server_databases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|