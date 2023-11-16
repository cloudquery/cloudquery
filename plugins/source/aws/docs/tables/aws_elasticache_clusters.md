# Table: aws_elasticache_clusters

This table shows data for Elasticache Clusters.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheCluster.html

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
|cache_cluster_create_time|`timestamp[us, tz=UTC]`|
|cache_cluster_id|`utf8`|
|cache_cluster_status|`utf8`|
|cache_node_type|`utf8`|
|cache_nodes|`json`|
|cache_parameter_group|`json`|
|cache_security_groups|`json`|
|cache_subnet_group_name|`utf8`|
|client_download_landing_page|`utf8`|
|configuration_endpoint|`json`|
|engine|`utf8`|
|engine_version|`utf8`|
|ip_discovery|`utf8`|
|log_delivery_configurations|`json`|
|network_type|`utf8`|
|notification_configuration|`json`|
|num_cache_nodes|`int64`|
|pending_modified_values|`json`|
|preferred_availability_zone|`utf8`|
|preferred_maintenance_window|`utf8`|
|preferred_outpost_arn|`utf8`|
|replication_group_id|`utf8`|
|replication_group_log_delivery_enabled|`bool`|
|security_groups|`json`|
|snapshot_retention_limit|`int64`|
|snapshot_window|`utf8`|
|transit_encryption_enabled|`bool`|
|transit_encryption_mode|`utf8`|