
# Table: aws_elasticache_replication_group_node_groups
Represents a collection of cache nodes in a replication group
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|replication_group_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_replication_groups table (FK)|
|node_group_id|text|The identifier for the node group (shard)|
|primary_endpoint_address|text|The DNS hostname of the cache node.|
|primary_endpoint_port|bigint|The port number that the cache engine is listening on.|
|reader_endpoint_address|text|The DNS hostname of the cache node.|
|reader_endpoint_port|bigint|The port number that the cache engine is listening on.|
|slots|text|The keyspace for this node group (shard).|
|status|text|The current state of this replication group - creating, available, modifying, deleting.|
