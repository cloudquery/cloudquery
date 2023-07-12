# Table: aws_cloudwatchlogs_metric_filters

This table shows data for Cloudwatchlogs Metric Filters.

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilter.html

The composite primary key for this table is (**log_group_arn**, **filter_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|log_group_arn (PK)|`utf8`|
|creation_time|`int64`|
|filter_name (PK)|`utf8`|
|filter_pattern|`utf8`|
|log_group_name|`utf8`|
|metric_transformations|`json`|