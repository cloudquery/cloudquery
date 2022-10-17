# Table: aws_neptune_cluster_snapshots

https://docs.aws.amazon.com/neptune/latest/userguide/api-snapshots.html#DescribeDBClusterSnapshots

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
|attributes|JSON|
|tags|JSON|
|allocated_storage|Int|
|availability_zones|StringArray|
|cluster_create_time|Timestamp|
|db_cluster_identifier|String|
|db_cluster_snapshot_identifier|String|
|engine|String|
|engine_version|String|
|iam_database_authentication_enabled|Bool|
|kms_key_id|String|
|license_model|String|
|master_username|String|
|percent_progress|Int|
|port|Int|
|snapshot_create_time|Timestamp|
|snapshot_type|String|
|source_db_cluster_snapshot_arn|String|
|status|String|
|storage_encrypted|Bool|
|vpc_id|String|