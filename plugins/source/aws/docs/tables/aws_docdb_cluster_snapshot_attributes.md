# Table: aws_docdb_cluster_snapshot_attributes



The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|db_cluster_snapshot_attributes|JSON|
|db_cluster_snapshot_identifier|String|