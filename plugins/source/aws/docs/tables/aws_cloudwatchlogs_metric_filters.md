# Table: aws_cloudwatchlogs_metric_filters



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
|creation_time|Int|
|filter_name|String|
|filter_pattern|String|
|log_group_name|String|
|metric_transformations|JSON|