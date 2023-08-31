# Table: aws_cloudwatchlogs_log_groups

This table shows data for Cloudwatchlogs Log Groups.

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_cloudwatchlogs_log_groups:
  - [aws_cloudwatchlogs_log_group_data_protection_policies](aws_cloudwatchlogs_log_group_data_protection_policies)
  - [aws_cloudwatchlogs_log_group_subscription_filters](aws_cloudwatchlogs_log_group_subscription_filters)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|creation_time|`int64`|
|data_protection_status|`utf8`|
|inherited_properties|`list<item: utf8, nullable>`|
|kms_key_id|`utf8`|
|log_group_name|`utf8`|
|metric_filter_count|`int64`|
|retention_in_days|`int64`|
|stored_bytes|`int64`|