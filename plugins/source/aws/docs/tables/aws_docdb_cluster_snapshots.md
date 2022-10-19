# Table: aws_docdb_cluster_snapshots



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_docdb_clusters](aws_docdb_clusters.md).
The following tables depend on aws_docdb_cluster_snapshots:
  - [aws_docdb_cluster_snapshot_attributes](aws_docdb_cluster_snapshot_attributes.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|availability_zones|StringArray|
|cluster_create_time|Timestamp|
|db_cluster_identifier|String|
|db_cluster_snapshot_arn|String|
|db_cluster_snapshot_identifier|String|
|engine|String|
|engine_version|String|
|kms_key_id|String|
|master_username|String|
|percent_progress|Int|
|port|Int|
|snapshot_create_time|Timestamp|
|snapshot_type|String|
|source_db_cluster_snapshot_arn|String|
|status|String|
|storage_encrypted|Bool|
|vpc_id|String|