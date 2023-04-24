# Table: aws_cloudwatchlogs_log_group_data_protection_policies

This table shows data for Cloudwatchlogs Log Group Data Protection Policies.

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetDataProtectionPolicy.html

The primary key for this table is **log_group_arn**.

## Relations

This table depends on [aws_cloudwatchlogs_log_groups](aws_cloudwatchlogs_log_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|log_group_arn (PK)|String|
|last_updated_time|Int|
|log_group_identifier|String|
|policy_document|String|