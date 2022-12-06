# Table: aws_rds_db_snapshots

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSnapshot.html

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
|attributes|JSON|
|allocated_storage|Int|
|availability_zone|String|
|db_instance_identifier|String|
|db_snapshot_identifier|String|
|dbi_resource_id|String|
|encrypted|Bool|
|engine|String|
|engine_version|String|
|iam_database_authentication_enabled|Bool|
|instance_create_time|Timestamp|
|iops|Int|
|kms_key_id|String|
|license_model|String|
|master_username|String|
|option_group_name|String|
|original_snapshot_create_time|Timestamp|
|percent_progress|Int|
|port|Int|
|processor_features|JSON|
|snapshot_create_time|Timestamp|
|snapshot_database_time|Timestamp|
|snapshot_target|String|
|snapshot_type|String|
|source_db_snapshot_identifier|String|
|source_region|String|
|status|String|
|storage_throughput|Int|
|storage_type|String|
|tde_credential_arn|String|
|timezone|String|
|vpc_id|String|