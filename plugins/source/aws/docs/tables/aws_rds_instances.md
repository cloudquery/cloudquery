# Table: aws_rds_instances

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBInstance.html

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
|processor_features|JSON|
|tags|JSON|
|activity_stream_engine_native_audit_fields_included|Bool|
|activity_stream_kinesis_stream_name|String|
|activity_stream_kms_key_id|String|
|activity_stream_mode|String|
|activity_stream_policy_status|String|
|activity_stream_status|String|
|allocated_storage|Int|
|associated_roles|JSON|
|auto_minor_version_upgrade|Bool|
|automatic_restart_time|Timestamp|
|automation_mode|String|
|availability_zone|String|
|aws_backup_recovery_point_arn|String|
|backup_retention_period|Int|
|backup_target|String|
|ca_certificate_identifier|String|
|character_set_name|String|
|copy_tags_to_snapshot|Bool|
|custom_iam_instance_profile|String|
|customer_owned_ip_enabled|Bool|
|db_cluster_identifier|String|
|db_instance_automated_backups_replications|JSON|
|db_instance_class|String|
|db_instance_identifier|String|
|db_instance_status|String|
|db_name|String|
|db_parameter_groups|JSON|
|db_security_groups|JSON|
|db_subnet_group|JSON|
|db_system_id|String|
|db_instance_port|Int|
|dbi_resource_id|String|
|deletion_protection|Bool|
|domain_memberships|JSON|
|enabled_cloudwatch_logs_exports|StringArray|
|endpoint|JSON|
|engine|String|
|engine_version|String|
|enhanced_monitoring_resource_arn|String|
|iam_database_authentication_enabled|Bool|
|instance_create_time|Timestamp|
|iops|Int|
|kms_key_id|String|
|latest_restorable_time|Timestamp|
|license_model|String|
|listener_endpoint|JSON|
|master_username|String|
|max_allocated_storage|Int|
|monitoring_interval|Int|
|monitoring_role_arn|String|
|multi_az|Bool|
|nchar_character_set_name|String|
|network_type|String|
|option_group_memberships|JSON|
|pending_modified_values|JSON|
|performance_insights_enabled|Bool|
|performance_insights_kms_key_id|String|
|performance_insights_retention_period|Int|
|preferred_backup_window|String|
|preferred_maintenance_window|String|
|promotion_tier|Int|
|publicly_accessible|Bool|
|read_replica_db_cluster_identifiers|StringArray|
|read_replica_db_instance_identifiers|StringArray|
|read_replica_source_db_instance_identifier|String|
|replica_mode|String|
|resume_full_automation_mode_time|Timestamp|
|secondary_availability_zone|String|
|status_infos|JSON|
|storage_encrypted|Bool|
|storage_throughput|Int|
|storage_type|String|
|tde_credential_arn|String|
|timezone|String|
|vpc_security_groups|JSON|