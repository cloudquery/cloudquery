# Table: aws_docdb_cluster_snapshot_attributes

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBClusterSnapshotAttributesResult.html

The primary key for this table is **db_cluster_snapshot_arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|db_cluster_snapshot_arn (PK)|String|
|db_cluster_snapshot_attributes|JSON|
|db_cluster_snapshot_identifier|String|