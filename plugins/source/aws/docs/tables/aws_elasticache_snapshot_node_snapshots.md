
# Table: aws_elasticache_snapshot_node_snapshots
Represents an individual cache node in a snapshot of a cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|snapshot_cq_id|uuid|Unique CloudQuery ID of aws_elasticache_snapshots table (FK)|
|cache_cluster_id|text|A unique identifier for the source cluster.|
|cache_node_create_time|timestamp without time zone|The date and time when the cache node was created in the source cluster.|
|cache_node_id|text|The cache node identifier for the node in the source cluster.|
|cache_size|text|The size of the cache on the source cache node.|
|node_group_configuration_node_group_id|text|Either the ElastiCache for Redis supplied 4-digit id or a user supplied id for the node group these configuration values apply to.|
|node_group_configuration_primary_availability_zone|text|The Availability Zone where the primary node of this node group (shard) is launched.|
|node_group_configuration_primary_outpost_arn|text|The outpost ARN of the primary node.|
|node_group_configuration_replica_availability_zones|text[]|A list of Availability Zones to be used for the read replicas|
|node_group_configuration_replica_count|bigint|The number of read replica nodes in this node group (shard).|
|node_group_configuration_replica_outpost_arns|text[]|The outpost ARN of the node replicas.|
|node_group_configuration_slots|text|A string that specifies the keyspace for a particular node group|
|node_group_id|text|A unique identifier for the source node group (shard).|
|snapshot_create_time|timestamp without time zone|The date and time when the source node's metadata and cache data set was obtained for the snapshot.|
