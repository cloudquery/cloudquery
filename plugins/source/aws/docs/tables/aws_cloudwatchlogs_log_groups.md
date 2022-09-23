# Table: aws_cloudwatchlogs_log_groups


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|