# Table: aws_cloudwatchlogs_log_group_subscription_filters

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_SubscriptionFilter.html

The composite primary key for this table is (**log_group_arn**, **creation_time**, **filter_name**).

## Relations

This table depends on [aws_cloudwatchlogs_log_groups](aws_cloudwatchlogs_log_groups.md).

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
|creation_time (PK)|Int|
|destination_arn|String|
|distribution|String|
|filter_name (PK)|String|
|filter_pattern|String|
|log_group_name|String|
|role_arn|String|