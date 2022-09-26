# Table: aws_cloudwatchlogs_metric_filters


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|creation_time|Int|
|filter_name|String|
|filter_pattern|String|
|log_group_name|String|
|metric_transformations|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|