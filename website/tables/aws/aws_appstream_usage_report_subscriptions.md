# Table: aws_appstream_usage_report_subscriptions

This table shows data for Amazon AppStream Usage Report Subscriptions.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UsageReportSubscription.html

The composite primary key for this table is (**account_id**, **region**, **s3_bucket_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|s3_bucket_name (PK)|`utf8`|
|last_generated_report_date|`timestamp[us, tz=UTC]`|
|schedule|`utf8`|
|subscription_errors|`json`|