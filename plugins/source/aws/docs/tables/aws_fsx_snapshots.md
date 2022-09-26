# Table: aws_fsx_snapshots


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|administrative_actions|JSON|
|creation_time|Timestamp|
|lifecycle|String|
|lifecycle_transition_reason|JSON|
|name|String|
|snapshot_id|String|
|volume_id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|