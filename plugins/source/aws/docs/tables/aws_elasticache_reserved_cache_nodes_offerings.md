
# Table: aws_elasticache_reserved_cache_nodes_offerings
Describes all of the attributes of a reserved cache node offering.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|cache_node_type|text|The cache node type for the reserved cache node|
|duration|bigint|The duration of the offering|
|fixed_price|float|The fixed price charged for this offering.|
|offering_type|text|The offering type.|
|product_description|text|The cache engine used by the offering.|
|reserved_cache_nodes_offering_id|text|A unique identifier for the reserved cache node offering.|
|usage_price|float|The hourly price charged for this offering.|
