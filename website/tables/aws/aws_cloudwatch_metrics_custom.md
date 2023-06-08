# Table: aws_cloudwatch_metrics_custom

This table shows data for Cloudwatch Metrics Custom.

https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|dimensions|`json`|
|metric_name|`utf8`|
|namespace|`utf8`|