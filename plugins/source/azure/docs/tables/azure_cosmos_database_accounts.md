# Table: azure_cosmos_database_accounts



The primary key for this table is **id**.

## Relations

The following tables depend on azure_cosmos_database_accounts:
  - [azure_cosmos_mongo_db_databases](azure_cosmos_mongo_db_databases.md)
  - [azure_cosmos_sql_databases](azure_cosmos_sql_databases.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|identity|JSON|
|kind|String|
|location|String|
|properties|JSON|
|tags|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|