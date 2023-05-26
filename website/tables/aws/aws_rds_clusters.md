# Table: aws_rds_clusters

This table shows data for Amazon Relational Database Service (RDS) Clusters.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBCluster.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_rds_clusters:
  - [aws_rds_cluster_backtracks](aws_rds_cluster_backtracks)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|tags|json|
|activity_stream_kinesis_stream_name|utf8|
|activity_stream_kms_key_id|utf8|
|activity_stream_mode|utf8|
|activity_stream_status|utf8|
|allocated_storage|int64|
|associated_roles|json|
|auto_minor_version_upgrade|bool|
|automatic_restart_time|timestamp[us, tz=UTC]|
|availability_zones|list<item: utf8, nullable>|
|backtrack_consumed_change_records|int64|
|backtrack_window|int64|
|backup_retention_period|int64|
|capacity|int64|
|character_set_name|utf8|
|clone_group_id|utf8|
|cluster_create_time|timestamp[us, tz=UTC]|
|copy_tags_to_snapshot|bool|
|cross_account_clone|bool|
|custom_endpoints|list<item: utf8, nullable>|
|db_cluster_arn|utf8|
|db_cluster_identifier|utf8|
|db_cluster_instance_class|utf8|
|db_cluster_members|json|
|db_cluster_option_group_memberships|json|
|db_cluster_parameter_group|utf8|
|db_subnet_group|utf8|
|db_system_id|utf8|
|database_name|utf8|
|db_cluster_resource_id|utf8|
|deletion_protection|bool|
|domain_memberships|json|
|earliest_backtrack_time|timestamp[us, tz=UTC]|
|earliest_restorable_time|timestamp[us, tz=UTC]|
|enabled_cloudwatch_logs_exports|list<item: utf8, nullable>|
|endpoint|utf8|
|engine|utf8|
|engine_mode|utf8|
|engine_version|utf8|
|global_write_forwarding_requested|bool|
|global_write_forwarding_status|utf8|
|hosted_zone_id|utf8|
|http_endpoint_enabled|bool|
|iam_database_authentication_enabled|bool|
|iops|int64|
|kms_key_id|utf8|
|latest_restorable_time|timestamp[us, tz=UTC]|
|master_user_secret|json|
|master_username|utf8|
|monitoring_interval|int64|
|monitoring_role_arn|utf8|
|multi_az|bool|
|network_type|utf8|
|pending_modified_values|json|
|percent_progress|utf8|
|performance_insights_enabled|bool|
|performance_insights_kms_key_id|utf8|
|performance_insights_retention_period|int64|
|port|int64|
|preferred_backup_window|utf8|
|preferred_maintenance_window|utf8|
|publicly_accessible|bool|
|read_replica_identifiers|list<item: utf8, nullable>|
|reader_endpoint|utf8|
|replication_source_identifier|utf8|
|scaling_configuration_info|json|
|serverless_v2_scaling_configuration|json|
|status|utf8|
|storage_encrypted|bool|
|storage_type|utf8|
|vpc_security_groups|json|