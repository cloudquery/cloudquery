# Table: aws_cloudwatchlogs_metric_filters

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilter.html

The composite primary key for this table is (**account_id**, **region**, **filter_name**, **log_group_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|creation_time|Int|
|filter_name (PK)|String|
|filter_pattern|String|
|log_group_name (PK)|String|
|metric_transformations|JSON|