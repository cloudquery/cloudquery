# Table: azure_cosmos_database_accounts

This table shows data for Azure Cosmos Database Accounts.

https://learn.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2022-05-15/database-accounts/list?tabs=HTTP#databaseaccountgetresults

The primary key for this table is **id**.

## Relations

The following tables depend on azure_cosmos_database_accounts:
  - [azure_cosmos_mongo_db_databases](azure_cosmos_mongo_db_databases)
  - [azure_cosmos_sql_databases](azure_cosmos_sql_databases)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|identity|`json`|
|kind|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|