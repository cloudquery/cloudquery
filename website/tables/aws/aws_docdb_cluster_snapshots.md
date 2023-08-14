# Table: aws_docdb_cluster_snapshots

This table shows data for Amazon DocumentDB Cluster Snapshots.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterSnapshot.html

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
|attributes|`json`|
|db_cluster_identifier|`utf8`|
|db_cluster_snapshot_identifier|`utf8`|
|availability_zones|`list<item: utf8, nullable>`|
|cluster_create_time|`timestamp[us, tz=UTC]`|
|db_cluster_snapshot_arn|`utf8`|
|engine|`utf8`|
|engine_version|`utf8`|
|kms_key_id|`utf8`|
|master_username|`utf8`|
|percent_progress|`int64`|
|port|`int64`|
|snapshot_create_time|`timestamp[us, tz=UTC]`|
|snapshot_type|`utf8`|
|source_db_cluster_snapshot_arn|`utf8`|
|status|`utf8`|
|storage_encrypted|`bool`|
|vpc_id|`utf8`|