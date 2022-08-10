
# Table: azure_cosmosdb_account_failover_policies
FailoverPolicy the failover policy for a given region of a database account.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_cq_id|uuid|Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)|
|id|text|The unique identifier of the region in which the database account replicates to|
|location_name|text|The name of the region in which the database account exists.|
|failover_priority|integer|The failover priority of the region|
