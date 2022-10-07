# Table: aws_cloudwatchlogs_log_groups



The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|creation_time|Int|
|kms_key_id|String|
|log_group_name|String|
|metric_filter_count|Int|
|retention_in_days|Int|
|stored_bytes|Int|