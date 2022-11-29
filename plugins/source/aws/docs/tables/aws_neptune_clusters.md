# Table: aws_neptune_clusters

https://docs.aws.amazon.com/neptune/latest/userguide/api-clusters.html#DescribeDBClusters

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
|tags|JSON|
|allocated_storage|Int|
|associated_roles|JSON|
|automatic_restart_time|Timestamp|
|availability_zones|StringArray|
|backup_retention_period|Int|
|character_set_name|String|
|clone_group_id|String|
|cluster_create_time|Timestamp|
|copy_tags_to_snapshot|Bool|
|cross_account_clone|Bool|
|db_cluster_identifier|String|
|db_cluster_members|JSON|
|db_cluster_option_group_memberships|JSON|
|db_cluster_parameter_group|String|
|db_subnet_group|String|
|database_name|String|
|db_cluster_resource_id|String|
|deletion_protection|Bool|
|earliest_restorable_time|Timestamp|
|enabled_cloudwatch_logs_exports|StringArray|
|endpoint|String|
|engine|String|
|engine_version|String|
|hosted_zone_id|String|
|iam_database_authentication_enabled|Bool|
|kms_key_id|String|
|latest_restorable_time|Timestamp|
|master_username|String|
|multi_az|Bool|
|percent_progress|String|
|port|Int|
|preferred_backup_window|String|
|preferred_maintenance_window|String|
|read_replica_identifiers|StringArray|
|reader_endpoint|String|
|replication_source_identifier|String|
|serverless_v2_scaling_configuration|JSON|
|status|String|
|storage_encrypted|Bool|
|vpc_security_groups|JSON|