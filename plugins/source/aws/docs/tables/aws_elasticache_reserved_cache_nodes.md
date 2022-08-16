
# Table: aws_elasticache_reserved_cache_nodes
Reserved Elasticache Cache Nodes
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|cache_node_count|bigint|The number of cache nodes that have been reserved.|
|cache_node_type|text|The cache node type for the reserved cache nodes|
|duration|bigint|The duration of the reservation in seconds.|
|fixed_price|float|The fixed price charged for this reserved cache node.|
|offering_type|text|The offering type of this reserved cache node.|
|product_description|text|The description of the reserved cache node.|
|reservation_arn|text|The Amazon Resource Name (ARN) of the reserved cache node|
|reserved_cache_node_id|text|The unique identifier for the reservation.|
|reserved_cache_nodes_offering_id|text|The offering identifier.|
|start_time|timestamp without time zone|The time the reservation started.|
|state|text|The state of the reserved cache node.|
|usage_price|float|The hourly price charged for this reserved cache node.|
