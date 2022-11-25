# Table: aws_elasticache_clusters

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheCluster.html

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
|cache_cluster_create_time|Timestamp|
|cache_cluster_id|String|
|cache_cluster_status|String|
|cache_node_type|String|
|cache_nodes|JSON|
|cache_parameter_group|JSON|
|cache_security_groups|JSON|
|cache_subnet_group_name|String|
|client_download_landing_page|String|
|configuration_endpoint|JSON|
|engine|String|
|engine_version|String|
|ip_discovery|String|
|log_delivery_configurations|JSON|
|network_type|String|
|notification_configuration|JSON|
|num_cache_nodes|Int|
|pending_modified_values|JSON|
|preferred_availability_zone|String|
|preferred_maintenance_window|String|
|preferred_outpost_arn|String|
|replication_group_id|String|
|replication_group_log_delivery_enabled|Bool|
|security_groups|JSON|
|snapshot_retention_limit|Int|
|snapshot_window|String|
|transit_encryption_enabled|Bool|