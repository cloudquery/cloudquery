
# Table: aws_elasticache_cluster_cache_nodes
Represents an individual cache node within a cluster
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_clusters table (FK)|
|create_time|timestamp without time zone|The date and time when the cache node was created.|
|id|text|The cache node identifier|
|status|text|The current state of this cache node, one of the following values: available, creating, rebooting, or deleting.|
|customer_availability_zone|text|The Availability Zone where this node was created and now resides.|
|customer_outpost_arn|text|The customer outpost ARN of the cache node.|
|endpoint_address|text|The DNS hostname of the cache node.|
|endpoint_port|bigint|The port number that the cache engine is listening on.|
|parameter_group_status|text|The status of the parameter group applied to this cache node.|
|source_cache_node_id|text|The ID of the primary node to which this read replica node is synchronized|
