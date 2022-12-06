# Table: aws_elasticache_snapshots

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Snapshot.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|auto_minor_version_upgrade|Bool|
|automatic_failover|String|
|cache_cluster_create_time|Timestamp|
|cache_cluster_id|String|
|cache_node_type|String|
|cache_parameter_group_name|String|
|cache_subnet_group_name|String|
|data_tiering|String|
|engine|String|
|engine_version|String|
|kms_key_id|String|
|node_snapshots|JSON|
|num_cache_nodes|Int|
|num_node_groups|Int|
|port|Int|
|preferred_availability_zone|String|
|preferred_maintenance_window|String|
|preferred_outpost_arn|String|
|replication_group_description|String|
|replication_group_id|String|
|snapshot_name|String|
|snapshot_retention_limit|Int|
|snapshot_source|String|
|snapshot_status|String|
|snapshot_window|String|
|topic_arn|String|
|vpc_id|String|