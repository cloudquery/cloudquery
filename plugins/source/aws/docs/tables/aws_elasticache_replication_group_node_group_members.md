
# Table: aws_elasticache_replication_group_node_group_members
Represents a single node within a node group (shard).
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|replication_group_node_group_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_replication_group_node_groups table (FK)|
|cache_cluster_id|text|The ID of the cluster to which the node belongs.|
|cache_node_id|text|The ID of the node within its cluster|
|current_role|text|The role that is currently assigned to the node - primary or replica|
|preferred_availability_zone|text|The name of the Availability Zone in which the node is located.|
|preferred_outpost_arn|text|The outpost ARN of the node group member.|
|read_endpoint_address|text|The DNS hostname of the cache node.|
|read_endpoint_port|bigint|The port number that the cache engine is listening on.|
