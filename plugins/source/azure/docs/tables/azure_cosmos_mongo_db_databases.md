# Table: azure_cosmos_mongo_db_databases

https://learn.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2022-05-15/mongo-db-resources/list-mongo-db-databases?tabs=HTTP#mongodbdatabasegetresults

The primary key for this table is **id**.

## Relations

This table depends on [azure_cosmos_database_accounts](azure_cosmos_database_accounts.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|location|String|
|properties|JSON|
|tags|JSON|
|name|String|
|type|String|