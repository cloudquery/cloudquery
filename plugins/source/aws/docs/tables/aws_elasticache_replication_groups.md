# Table: aws_elasticache_replication_groups

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReplicationGroup.html

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
|at_rest_encryption_enabled|Bool|
|auth_token_enabled|Bool|
|auth_token_last_modified_date|Timestamp|
|auto_minor_version_upgrade|Bool|
|automatic_failover|String|
|cache_node_type|String|
|cluster_enabled|Bool|
|configuration_endpoint|JSON|
|data_tiering|String|
|description|String|
|global_replication_group_info|JSON|
|ip_discovery|String|
|kms_key_id|String|
|log_delivery_configurations|JSON|
|member_clusters|StringArray|
|member_clusters_outpost_arns|StringArray|
|multi_az|String|
|network_type|String|
|node_groups|JSON|
|pending_modified_values|JSON|
|replication_group_create_time|Timestamp|
|replication_group_id|String|
|snapshot_retention_limit|Int|
|snapshot_window|String|
|snapshotting_cluster_id|String|
|status|String|
|transit_encryption_enabled|Bool|
|user_group_ids|StringArray|