# Table: aws_elasticache_replication_groups

This table shows data for Elasticache Replication Groups.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReplicationGroup.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|at_rest_encryption_enabled|`bool`|
|auth_token_enabled|`bool`|
|auth_token_last_modified_date|`timestamp[us, tz=UTC]`|
|auto_minor_version_upgrade|`bool`|
|automatic_failover|`utf8`|
|cache_node_type|`utf8`|
|cluster_enabled|`bool`|
|cluster_mode|`utf8`|
|configuration_endpoint|`json`|
|data_tiering|`utf8`|
|description|`utf8`|
|global_replication_group_info|`json`|
|ip_discovery|`utf8`|
|kms_key_id|`utf8`|
|log_delivery_configurations|`json`|
|member_clusters|`list<item: utf8, nullable>`|
|member_clusters_outpost_arns|`list<item: utf8, nullable>`|
|multi_az|`utf8`|
|network_type|`utf8`|
|node_groups|`json`|
|pending_modified_values|`json`|
|replication_group_create_time|`timestamp[us, tz=UTC]`|
|replication_group_id|`utf8`|
|snapshot_retention_limit|`int64`|
|snapshot_window|`utf8`|
|snapshotting_cluster_id|`utf8`|
|status|`utf8`|
|transit_encryption_enabled|`bool`|
|transit_encryption_mode|`utf8`|
|user_group_ids|`list<item: utf8, nullable>`|