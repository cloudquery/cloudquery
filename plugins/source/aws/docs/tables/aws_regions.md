# Table: aws_regions



The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|enabled|Bool|
|partition|String|
|region|String|
|endpoint|String|
|opt_in_status|String|