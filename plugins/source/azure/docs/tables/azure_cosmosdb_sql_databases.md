# Table: azure_cosmosdb_sql_databases


The primary key for this table is **id**.

## Relations
This table depends on [`azure_cosmosdb_accounts`](azure_cosmosdb_accounts.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|cosmosdb_account_id|UUID|
|resource|JSON|
|options|JSON|
|id (PK)|String|
|name|String|
|type|String|
|location|String|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|