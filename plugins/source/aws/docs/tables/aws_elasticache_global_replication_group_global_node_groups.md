
# Table: aws_elasticache_global_replication_group_global_node_groups
Indicates the slot configuration and global identifier for a slice group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|global_replication_group_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_global_replication_groups table (FK)|
|global_node_group_id|text|The name of the global node group|
|slots|text|The keyspace for this node group|
