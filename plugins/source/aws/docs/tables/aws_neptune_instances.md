# Table: aws_neptune_instances

This table shows data for Neptune Instances.

https://docs.aws.amazon.com/neptune/latest/userguide/api-instances.html#DescribeDBInstances

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
|auto_minor_version_upgrade|`bool`|
|availability_zone|`utf8`|
|backup_retention_period|`int64`|
|ca_certificate_identifier|`utf8`|
|character_set_name|`utf8`|
|copy_tags_to_snapshot|`bool`|
|db_cluster_identifier|`utf8`|
|db_instance_arn|`utf8`|
|db_instance_class|`utf8`|
|db_instance_identifier|`utf8`|
|db_instance_status|`utf8`|
|db_name|`utf8`|
|db_parameter_groups|`json`|
|db_security_groups|`json`|
|db_subnet_group|`json`|
|db_instance_port|`int64`|
|dbi_resource_id|`utf8`|
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
|kms_key_id|`utf8`|
|latest_restorable_time|`timestamp[us, tz=UTC]`|
|license_model|`utf8`|
|master_username|`utf8`|
|monitoring_interval|`int64`|
|monitoring_role_arn|`utf8`|
|multi_az|`bool`|
|option_group_memberships|`json`|
|pending_modified_values|`json`|
|performance_insights_enabled|`bool`|
|performance_insights_kms_key_id|`utf8`|
|preferred_backup_window|`utf8`|
|preferred_maintenance_window|`utf8`|
|promotion_tier|`int64`|
|publicly_accessible|`bool`|
|read_replica_db_cluster_identifiers|`list<item: utf8, nullable>`|
|read_replica_db_instance_identifiers|`list<item: utf8, nullable>`|
|read_replica_source_db_instance_identifier|`utf8`|
|secondary_availability_zone|`utf8`|
|status_infos|`json`|
|storage_encrypted|`bool`|
|storage_type|`utf8`|
|tde_credential_arn|`utf8`|
|timezone|`utf8`|
|vpc_security_groups|`json`|