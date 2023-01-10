# Table: aws_docdb_cluster_snapshots

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterSnapshot.html

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
|attributes|JSON|
|db_cluster_identifier|String|
|db_cluster_snapshot_identifier|String|
|availability_zones|StringArray|
|cluster_create_time|Timestamp|
|db_cluster_snapshot_arn|String|
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