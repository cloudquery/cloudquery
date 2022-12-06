# Table: aws_cloudwatchlogs_log_groups

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogGroup.html

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
|tags|JSON|
|creation_time|Int|
|data_protection_status|String|
|kms_key_id|String|
|log_group_name|String|
|metric_filter_count|Int|
|retention_in_days|Int|
|stored_bytes|Int|