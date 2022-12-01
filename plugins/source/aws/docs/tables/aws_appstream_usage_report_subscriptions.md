# Table: aws_appstream_usage_report_subscriptions

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_UsageReportSubscription.html

The composite primary key for this table is (**account_id**, **region**, **s3_bucket_name**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|s3_bucket_name (PK)|String|
|last_generated_report_date|Timestamp|
|schedule|String|
|subscription_errors|JSON|