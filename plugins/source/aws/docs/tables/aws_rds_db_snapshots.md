# Table: aws_rds_db_snapshots

This table shows data for Amazon Relational Database Service (RDS) DB Snapshots.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSnapshot.html

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
|tags|`json`|
|attributes|`json`|
|allocated_storage|`int64`|
|availability_zone|`utf8`|
|db_instance_identifier|`utf8`|
|db_snapshot_arn|`utf8`|
|db_snapshot_identifier|`utf8`|
|db_system_id|`utf8`|
|dbi_resource_id|`utf8`|
|dedicated_log_volume|`bool`|
|encrypted|`bool`|
|engine|`utf8`|
|engine_version|`utf8`|
|iam_database_authentication_enabled|`bool`|
|instance_create_time|`timestamp[us, tz=UTC]`|
|iops|`int64`|
|kms_key_id|`utf8`|
|license_model|`utf8`|
|master_username|`utf8`|
|multi_tenant|`bool`|
|option_group_name|`utf8`|
|original_snapshot_create_time|`timestamp[us, tz=UTC]`|
|percent_progress|`int64`|
|port|`int64`|
|processor_features|`json`|
|snapshot_create_time|`timestamp[us, tz=UTC]`|
|snapshot_database_time|`timestamp[us, tz=UTC]`|
|snapshot_target|`utf8`|
|snapshot_type|`utf8`|
|source_db_snapshot_identifier|`utf8`|
|source_region|`utf8`|
|status|`utf8`|
|storage_throughput|`int64`|
|storage_type|`utf8`|
|tde_credential_arn|`utf8`|
|timezone|`utf8`|
|vpc_id|`utf8`|