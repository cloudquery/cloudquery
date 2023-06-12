# Table: aws_cloudwatch_metrics

This table shows data for Cloudwatch Metrics.

https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html
To sync this table you must set the 'use_paid_apis' option to 'true' and set the relevant 'table_options' entry in the AWS provider configuration.

The composite primary key for this table is (**account_id**, **region**, **input_hash**).

## Relations

The following tables depend on aws_cloudwatch_metrics:
  - [aws_cloudwatch_metric_statistics](aws_cloudwatch_metric_statistics)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|input_hash (PK)|`utf8`|
|input_json|`json`|
|dimensions|`json`|
|metric_name|`utf8`|
|namespace|`utf8`|