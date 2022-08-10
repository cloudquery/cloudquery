
# Table: azure_cosmosdb_sql_databases
Azure Cosmos DB SQL database.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|database_id|text|Name of the Cosmos DB SQL database|
|database_rid|text|A system generated property|
|database_ts|float|A system generated property that denotes the last updated timestamp of the resource.|
|database_etag|text|A system generated property representing the resource etag required for optimistic concurrency control.|
|database_colls|text|A system generated property that specified the addressable path of the collections resource.|
|database_users|text|A system generated property that specifies the addressable path of the users resource.|
|sql_database_get_properties_throughput|integer|Value of the Cosmos DB resource throughput or autoscaleSettings|
|autoscale_settings_max_throughput|integer|Represents maximum throughput, the resource can scale up to.|
|id|text|The unique resource identifier of the ARM resource.|
|name|text|The name of the ARM resource.|
|type|text|The type of Azure resource.|
|location|text|The location of the resource group to which the resource belongs.|
|tags|jsonb|Resource tags.|
