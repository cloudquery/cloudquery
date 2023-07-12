# Table: aws_alpha_cloudwatch_metric_statistics

This table shows data for Cloudwatch Metric Statistics (Alpha).

https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html
To sync this table you must set the 'use_paid_apis' option to 'true' and set the relevant 'table_options' entry in the AWS provider configuration.

Please note that this table is considered **alpha** (experimental) and may have breaking changes or be removed in the future.

The composite primary key for this table is (**account_id**, **region**, **parent_input_hash**, **input_hash**, **timestamp**, **label**).

## Relations

This table depends on [aws_alpha_cloudwatch_metrics](aws_alpha_cloudwatch_metrics).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|parent_input_hash (PK)|`utf8`|
|input_hash (PK)|`utf8`|
|input_json|`json`|
|average|`float64`|
|extended_statistics|`json`|
|maximum|`float64`|
|minimum|`float64`|
|sample_count|`float64`|
|sum|`float64`|
|timestamp (PK)|`timestamp[us, tz=UTC]`|
|unit|`utf8`|
|label (PK)|`utf8`|