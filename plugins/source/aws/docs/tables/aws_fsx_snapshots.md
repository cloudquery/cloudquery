# Table: aws_fsx_snapshots

https://docs.aws.amazon.com/fsx/latest/APIReference/API_Snapshot.html

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
|administrative_actions|JSON|
|creation_time|Timestamp|
|lifecycle|String|
|lifecycle_transition_reason|JSON|
|name|String|
|snapshot_id|String|
|tags|JSON|
|volume_id|String|