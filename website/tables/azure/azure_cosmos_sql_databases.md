# Table: azure_cosmos_sql_databases

This table shows data for Azure Cosmos SQL Databases.

https://learn.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2022-05-15/sql-resources/list-sql-databases?tabs=HTTP#sqldatabasegetresults

The primary key for this table is **id**.

## Relations

This table depends on [azure_cosmos_database_accounts](azure_cosmos_database_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|location|String|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|type|String|