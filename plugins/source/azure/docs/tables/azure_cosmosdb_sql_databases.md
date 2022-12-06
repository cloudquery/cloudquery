# Table: azure_cosmosdb_sql_databases

https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb#SQLDatabaseGetResults

The primary key for this table is **id**.

## Relations
This table depends on [azure_cosmosdb_accounts](azure_cosmosdb_accounts.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|cosmosdb_account_id|String|
|resource|JSON|
|options|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|