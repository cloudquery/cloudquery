# Table: aws_neptune_clusters

This table shows data for Neptune Clusters.

https://docs.aws.amazon.com/neptune/latest/userguide/api-clusters.html#DescribeDBClusters

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
|allocated_storage|`int64`|
|associated_roles|`json`|
|automatic_restart_time|`timestamp[us, tz=UTC]`|
|availability_zones|`list<item: utf8, nullable>`|
|backup_retention_period|`int64`|
|character_set_name|`utf8`|
|clone_group_id|`utf8`|
|cluster_create_time|`timestamp[us, tz=UTC]`|
|copy_tags_to_snapshot|`bool`|
|cross_account_clone|`bool`|
|db_cluster_arn|`utf8`|
|db_cluster_identifier|`utf8`|
|db_cluster_members|`json`|
|db_cluster_option_group_memberships|`json`|
|db_cluster_parameter_group|`utf8`|
|db_subnet_group|`utf8`|
|database_name|`utf8`|
|db_cluster_resource_id|`utf8`|
|deletion_protection|`bool`|
|earliest_restorable_time|`timestamp[us, tz=UTC]`|
|enabled_cloudwatch_logs_exports|`list<item: utf8, nullable>`|
|endpoint|`utf8`|
|engine|`utf8`|
|engine_version|`utf8`|
|global_cluster_identifier|`utf8`|
|hosted_zone_id|`utf8`|
|iam_database_authentication_enabled|`bool`|
|kms_key_id|`utf8`|
|latest_restorable_time|`timestamp[us, tz=UTC]`|
|master_username|`utf8`|
|multi_az|`bool`|
|pending_modified_values|`json`|
|percent_progress|`utf8`|
|port|`int64`|
|preferred_backup_window|`utf8`|
|preferred_maintenance_window|`utf8`|
|read_replica_identifiers|`list<item: utf8, nullable>`|
|reader_endpoint|`utf8`|
|replication_source_identifier|`utf8`|
|serverless_v2_scaling_configuration|`json`|
|status|`utf8`|
|storage_encrypted|`bool`|
|vpc_security_groups|`json`|