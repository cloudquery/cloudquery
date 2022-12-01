# Table: aws_rds_clusters

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBCluster.html

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
|activity_stream_kinesis_stream_name|String|
|activity_stream_kms_key_id|String|
|activity_stream_mode|String|
|activity_stream_status|String|
|allocated_storage|Int|
|associated_roles|JSON|
|auto_minor_version_upgrade|Bool|
|automatic_restart_time|Timestamp|
|availability_zones|StringArray|
|backtrack_consumed_change_records|Int|
|backtrack_window|Int|
|backup_retention_period|Int|
|capacity|Int|
|character_set_name|String|
|clone_group_id|String|
|cluster_create_time|Timestamp|
|copy_tags_to_snapshot|Bool|
|cross_account_clone|Bool|
|custom_endpoints|StringArray|
|db_cluster_identifier|String|
|db_cluster_instance_class|String|
|db_cluster_members|JSON|
|db_cluster_option_group_memberships|JSON|
|db_cluster_parameter_group|String|
|db_subnet_group|String|
|db_system_id|String|
|database_name|String|
|db_cluster_resource_id|String|
|deletion_protection|Bool|
|domain_memberships|JSON|
|earliest_backtrack_time|Timestamp|
|earliest_restorable_time|Timestamp|
|enabled_cloudwatch_logs_exports|StringArray|
|endpoint|String|
|engine|String|
|engine_mode|String|
|engine_version|String|
|global_write_forwarding_requested|Bool|
|global_write_forwarding_status|String|
|hosted_zone_id|String|
|http_endpoint_enabled|Bool|
|iam_database_authentication_enabled|Bool|
|iops|Int|
|kms_key_id|String|
|latest_restorable_time|Timestamp|
|master_username|String|
|monitoring_interval|Int|
|monitoring_role_arn|String|
|multi_az|Bool|
|network_type|String|
|pending_modified_values|JSON|
|percent_progress|String|
|performance_insights_enabled|Bool|
|performance_insights_kms_key_id|String|
|performance_insights_retention_period|Int|
|port|Int|
|preferred_backup_window|String|
|preferred_maintenance_window|String|
|publicly_accessible|Bool|
|read_replica_identifiers|StringArray|
|reader_endpoint|String|
|replication_source_identifier|String|
|scaling_configuration_info|JSON|
|serverless_v2_scaling_configuration|JSON|
|status|String|
|storage_encrypted|Bool|
|storage_type|String|
|vpc_security_groups|JSON|