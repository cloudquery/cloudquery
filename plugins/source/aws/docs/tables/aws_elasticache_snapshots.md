# Table: aws_elasticache_snapshots

This table shows data for Elasticache Snapshots.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Snapshot.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|auto_minor_version_upgrade|`bool`|
|automatic_failover|`utf8`|
|cache_cluster_create_time|`timestamp[us, tz=UTC]`|
|cache_cluster_id|`utf8`|
|cache_node_type|`utf8`|
|cache_parameter_group_name|`utf8`|
|cache_subnet_group_name|`utf8`|
|data_tiering|`utf8`|
|engine|`utf8`|
|engine_version|`utf8`|
|kms_key_id|`utf8`|
|node_snapshots|`json`|
|num_cache_nodes|`int64`|
|num_node_groups|`int64`|
|port|`int64`|
|preferred_availability_zone|`utf8`|
|preferred_maintenance_window|`utf8`|
|preferred_outpost_arn|`utf8`|
|replication_group_description|`utf8`|
|replication_group_id|`utf8`|
|snapshot_name|`utf8`|
|snapshot_retention_limit|`int64`|
|snapshot_source|`utf8`|
|snapshot_status|`utf8`|
|snapshot_window|`utf8`|
|topic_arn|`utf8`|
|vpc_id|`utf8`|