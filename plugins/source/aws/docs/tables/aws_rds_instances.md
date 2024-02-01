# Table: aws_rds_instances

This table shows data for Amazon Relational Database Service (RDS) Instances.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBInstance.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|processor_features|`json`|
|tags|`json`|
|activity_stream_engine_native_audit_fields_included|`bool`|
|activity_stream_kinesis_stream_name|`utf8`|
|activity_stream_kms_key_id|`utf8`|
|activity_stream_mode|`utf8`|
|activity_stream_policy_status|`utf8`|
|activity_stream_status|`utf8`|
|allocated_storage|`int64`|
|associated_roles|`json`|
|auto_minor_version_upgrade|`bool`|
|automatic_restart_time|`timestamp[us, tz=UTC]`|
|automation_mode|`utf8`|
|availability_zone|`utf8`|
|aws_backup_recovery_point_arn|`utf8`|
|backup_retention_period|`int64`|
|backup_target|`utf8`|
|ca_certificate_identifier|`utf8`|
|certificate_details|`json`|
|character_set_name|`utf8`|
|copy_tags_to_snapshot|`bool`|
|custom_iam_instance_profile|`utf8`|
|customer_owned_ip_enabled|`bool`|
|db_cluster_identifier|`utf8`|
|db_instance_arn|`utf8`|
|db_instance_automated_backups_replications|`json`|
|db_instance_class|`utf8`|
|db_instance_identifier|`utf8`|
|db_instance_status|`utf8`|
|db_name|`utf8`|
|db_parameter_groups|`json`|
|db_security_groups|`json`|
|db_subnet_group|`json`|
|db_system_id|`utf8`|
|db_instance_port|`int64`|
|dbi_resource_id|`utf8`|
|dedicated_log_volume|`bool`|
|deletion_protection|`bool`|
|domain_memberships|`json`|
|enabled_cloudwatch_logs_exports|`list<item: utf8, nullable>`|
|endpoint|`json`|
|engine|`utf8`|
|engine_version|`utf8`|
|enhanced_monitoring_resource_arn|`utf8`|
|iam_database_authentication_enabled|`bool`|
|instance_create_time|`timestamp[us, tz=UTC]`|
|iops|`int64`|
|is_storage_config_upgrade_available|`bool`|
|kms_key_id|`utf8`|
|latest_restorable_time|`timestamp[us, tz=UTC]`|
|license_model|`utf8`|
|listener_endpoint|`json`|
|master_user_secret|`json`|
|master_username|`utf8`|
|max_allocated_storage|`int64`|
|monitoring_interval|`int64`|
|monitoring_role_arn|`utf8`|
|multi_az|`bool`|
|multi_tenant|`bool`|
|nchar_character_set_name|`utf8`|
|network_type|`utf8`|
|option_group_memberships|`json`|
|pending_modified_values|`json`|
|percent_progress|`utf8`|
|performance_insights_enabled|`bool`|
|performance_insights_kms_key_id|`utf8`|
|performance_insights_retention_period|`int64`|
|preferred_backup_window|`utf8`|
|preferred_maintenance_window|`utf8`|
|promotion_tier|`int64`|
|publicly_accessible|`bool`|
|read_replica_db_cluster_identifiers|`list<item: utf8, nullable>`|
|read_replica_db_instance_identifiers|`list<item: utf8, nullable>`|
|read_replica_source_db_cluster_identifier|`utf8`|
|read_replica_source_db_instance_identifier|`utf8`|
|replica_mode|`utf8`|
|resume_full_automation_mode_time|`timestamp[us, tz=UTC]`|
|secondary_availability_zone|`utf8`|
|status_infos|`json`|
|storage_encrypted|`bool`|
|storage_throughput|`int64`|
|storage_type|`utf8`|
|tde_credential_arn|`utf8`|
|timezone|`utf8`|
|vpc_security_groups|`json`|