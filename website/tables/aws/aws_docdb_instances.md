# Table: aws_docdb_instances

This table shows data for Amazon DocumentDB Instances.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_docdb_clusters](aws_docdb_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|auto_minor_version_upgrade|`bool`|
|availability_zone|`utf8`|
|backup_retention_period|`int64`|
|ca_certificate_identifier|`utf8`|
|copy_tags_to_snapshot|`bool`|
|db_cluster_identifier|`utf8`|
|db_instance_arn|`utf8`|
|db_instance_class|`utf8`|
|db_instance_identifier|`utf8`|
|db_instance_status|`utf8`|
|db_subnet_group|`json`|
|dbi_resource_id|`utf8`|
|enabled_cloudwatch_logs_exports|`list<item: utf8, nullable>`|
|endpoint|`json`|
|engine|`utf8`|
|engine_version|`utf8`|
|instance_create_time|`timestamp[us, tz=UTC]`|
|kms_key_id|`utf8`|
|latest_restorable_time|`timestamp[us, tz=UTC]`|
|pending_modified_values|`json`|
|preferred_backup_window|`utf8`|
|preferred_maintenance_window|`utf8`|
|promotion_tier|`int64`|
|publicly_accessible|`bool`|
|status_infos|`json`|
|storage_encrypted|`bool`|
|vpc_security_groups|`json`|