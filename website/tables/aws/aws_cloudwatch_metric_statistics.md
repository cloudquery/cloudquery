# Table: aws_cloudwatch_metric_statistics

This table shows data for Cloudwatch Metric Statistics.

https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html
To sync this table you must set the 'use_paid_apis' option to 'true' and set the relevant 'table_options' entry in the AWS provider configuration.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_cloudwatch_metrics](aws_cloudwatch_metrics).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|datapoints|`json`|
|label|`utf8`|
|input_json|`json`|