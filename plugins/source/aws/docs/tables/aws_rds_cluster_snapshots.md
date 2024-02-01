# Table: aws_rds_cluster_snapshots

This table shows data for Amazon Relational Database Service (RDS) Cluster Snapshots.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterSnapshot.html

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
|availability_zones|`list<item: utf8, nullable>`|
|cluster_create_time|`timestamp[us, tz=UTC]`|
|db_cluster_identifier|`utf8`|
|db_cluster_snapshot_arn|`utf8`|
|db_cluster_snapshot_identifier|`utf8`|
|db_system_id|`utf8`|
|db_cluster_resource_id|`utf8`|
|engine|`utf8`|
|engine_mode|`utf8`|
|engine_version|`utf8`|
|iam_database_authentication_enabled|`bool`|
|kms_key_id|`utf8`|
|license_model|`utf8`|
|master_username|`utf8`|
|percent_progress|`int64`|
|port|`int64`|
|snapshot_create_time|`timestamp[us, tz=UTC]`|
|snapshot_type|`utf8`|
|source_db_cluster_snapshot_arn|`utf8`|
|status|`utf8`|
|storage_encrypted|`bool`|
|storage_type|`utf8`|
|vpc_id|`utf8`|