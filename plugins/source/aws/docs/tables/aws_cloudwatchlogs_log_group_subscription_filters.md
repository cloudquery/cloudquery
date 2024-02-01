# Table: aws_cloudwatchlogs_log_group_subscription_filters

This table shows data for Cloudwatchlogs Log Group Subscription Filters.

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_SubscriptionFilter.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**log_group_arn**, **creation_time**, **filter_name**).
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
|creation_time|`int64`|
|destination_arn|`utf8`|
|distribution|`utf8`|
|filter_name|`utf8`|
|filter_pattern|`utf8`|
|log_group_name|`utf8`|
|role_arn|`utf8`|