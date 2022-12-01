# Table: aws_neptune_instances

https://docs.aws.amazon.com/neptune/latest/userguide/api-instances.html#DescribeDBInstances

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
|auto_minor_version_upgrade|Bool|
|availability_zone|String|
|backup_retention_period|Int|
|ca_certificate_identifier|String|
|character_set_name|String|
|copy_tags_to_snapshot|Bool|
|db_cluster_identifier|String|
|db_instance_class|String|
|db_instance_identifier|String|
|db_instance_status|String|
|db_name|String|
|db_parameter_groups|JSON|
|db_security_groups|JSON|
|db_subnet_group|JSON|
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
|master_username|String|
|monitoring_interval|Int|
|monitoring_role_arn|String|
|multi_az|Bool|
|option_group_memberships|JSON|
|pending_modified_values|JSON|
|performance_insights_enabled|Bool|
|performance_insights_kms_key_id|String|
|preferred_backup_window|String|
|preferred_maintenance_window|String|
|promotion_tier|Int|
|publicly_accessible|Bool|
|read_replica_db_cluster_identifiers|StringArray|
|read_replica_db_instance_identifiers|StringArray|
|read_replica_source_db_instance_identifier|String|
|secondary_availability_zone|String|
|status_infos|JSON|
|storage_encrypted|Bool|
|storage_type|String|
|tde_credential_arn|String|
|timezone|String|
|vpc_security_groups|JSON|