# Table: aws_regions

Describes a Region.

The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|enabled|Bool|
|endpoint|String|
|opt_in_status|String|
|region|String|
|partition|String|