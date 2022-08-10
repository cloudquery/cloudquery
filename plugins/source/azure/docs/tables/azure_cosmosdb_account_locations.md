
# Table: azure_cosmosdb_account_locations
Location a region in which the Azure Cosmos DB database account is deployed.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_cq_id|uuid|Unique CloudQuery ID of azure_cosmosdb_accounts table (FK)|
|id|text|The unique identifier of the region within the database account|
|location_name|text|The name of the region.|
|document_endpoint|text|The connection endpoint for the specific region|
|provisioning_state|text||
|failover_priority|integer|The failover priority of the region|
|is_zone_redundant|boolean|Flag to indicate whether or not this region is an AvailabilityZone region|
