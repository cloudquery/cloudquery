
# Table: aws_elasticache_snapshots
Represents a copy of an entire Redis cluster as of the time when the snapshot was taken.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The ARN (Amazon Resource Name) of the snapshot.|
|auto_minor_version_upgrade|boolean|Auto minor version upgrade|
|automatic_failover|text|Indicates the status of automatic failover for the source Redis replication group.|
|cache_cluster_create_time|timestamp without time zone|The date and time when the source cluster was created.|
|cache_cluster_id|text|The user-supplied identifier of the source cluster.|
|cache_node_type|text|The name of the compute and memory capacity node type for the source cluster. The following node types are supported by ElastiCache|
|cache_parameter_group_name|text|The cache parameter group that is associated with the source cluster.|
|cache_subnet_group_name|text|The name of the cache subnet group associated with the source cluster.|
|data_tiering|text|Data tiering|
|engine|text|The name of the cache engine (memcached or redis) used by the source cluster.|
|engine_version|text|The version of the cache engine version that is used by the source cluster.|
|kms_key_id|text|The ID of the KMS key used to encrypt the snapshot.|
|num_cache_nodes|bigint|The number of cache nodes in the source cluster|
|num_node_groups|bigint|The number of node groups (shards) in this snapshot|
|port|bigint|The port number used by each cache nodes in the source cluster.|
|preferred_availability_zone|text|The name of the Availability Zone in which the source cluster is located.|
|preferred_maintenance_window|text|Specifies the weekly time range during which maintenance on the cluster is performed|
|preferred_outpost_arn|text|The ARN (Amazon Resource Name) of the preferred outpost.|
|replication_group_description|text|A description of the source replication group.|
|replication_group_id|text|The unique identifier of the source replication group.|
|snapshot_name|text|The name of a snapshot|
|snapshot_retention_limit|bigint|For an automatic snapshot, the number of days for which ElastiCache retains the snapshot before deleting it|
|snapshot_source|text|Indicates whether the snapshot is from an automatic backup (automated) or was created manually (manual).|
|snapshot_status|text|The status of the snapshot|
|snapshot_window|text|The daily time range during which ElastiCache takes daily snapshots of the source cluster.|
|topic_arn|text|The Amazon Resource Name (ARN) for the topic used by the source cluster for publishing notifications.|
|vpc_id|text|The Amazon Virtual Private Cloud identifier (VPC ID) of the cache subnet group for the source cluster.|
