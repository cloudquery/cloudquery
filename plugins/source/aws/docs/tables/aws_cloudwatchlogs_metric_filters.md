# Table: aws_cloudwatchlogs_metric_filters

This table shows data for Cloudwatchlogs Metric Filters.

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilter.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**log_group_arn**, **filter_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|log_group_arn|`utf8`|
|creation_time|`int64`|
|filter_name|`utf8`|
|filter_pattern|`utf8`|
|log_group_name|`utf8`|
|metric_transformations|`json`|