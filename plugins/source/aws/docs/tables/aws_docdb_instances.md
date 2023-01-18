# Table: aws_docdb_instances

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_docdb_clusters](aws_docdb_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|auto_minor_version_upgrade|Bool|
|availability_zone|String|
|backup_retention_period|Int|
|ca_certificate_identifier|String|
|copy_tags_to_snapshot|Bool|
|db_cluster_identifier|String|
|db_instance_arn|String|
|db_instance_class|String|
|db_instance_identifier|String|
|db_instance_status|String|
|db_subnet_group|JSON|
|dbi_resource_id|String|
|enabled_cloudwatch_logs_exports|StringArray|
|endpoint|JSON|
|engine|String|
|engine_version|String|
|instance_create_time|Timestamp|
|kms_key_id|String|
|latest_restorable_time|Timestamp|
|pending_modified_values|JSON|
|preferred_backup_window|String|
|preferred_maintenance_window|String|
|promotion_tier|Int|
|publicly_accessible|Bool|
|status_infos|JSON|
|storage_encrypted|Bool|
|vpc_security_groups|JSON|