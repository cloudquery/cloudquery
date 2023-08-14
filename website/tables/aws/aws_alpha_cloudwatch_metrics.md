# Table: aws_alpha_cloudwatch_metrics

This table shows data for Cloudwatch Metrics (Alpha).

https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html
To sync this table you must set the 'use_paid_apis' option to 'true' and set the relevant 'table_options' entry in the AWS provider configuration.

Please note that this table is considered **alpha** (experimental) and may have breaking changes or be removed in the future.

The composite primary key for this table is (**account_id**, **region**, **input_hash**).

## Relations

The following tables depend on aws_alpha_cloudwatch_metrics:
  - [aws_alpha_cloudwatch_metric_statistics](aws_alpha_cloudwatch_metric_statistics)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|input_hash (PK)|`utf8`|
|input_json|`json`|
|dimensions|`json`|
|metric_name|`utf8`|
|namespace|`utf8`|