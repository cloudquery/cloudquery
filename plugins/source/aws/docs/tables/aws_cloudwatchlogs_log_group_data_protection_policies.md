# Table: aws_cloudwatchlogs_log_group_data_protection_policies

This table shows data for Cloudwatchlogs Log Group Data Protection Policies.

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetDataProtectionPolicy.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **log_group_arn**.
## Relations

This table depends on [aws_cloudwatchlogs_log_groups](aws_cloudwatchlogs_log_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|log_group_arn|`utf8`|
|last_updated_time|`int64`|
|log_group_identifier|`utf8`|
|policy_document|`utf8`|