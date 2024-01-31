# Table: aws_appstream_usage_report_subscriptions

This table shows data for Amazon AppStream Usage Report Subscriptions.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UsageReportSubscription.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **s3_bucket_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|s3_bucket_name|`utf8`|
|last_generated_report_date|`timestamp[us, tz=UTC]`|
|schedule|`utf8`|
|subscription_errors|`json`|