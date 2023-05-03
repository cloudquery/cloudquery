# Table: aws_securityhub_hubs

This table shows data for AWS Security Hub Hubs.

https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeHub.html

The composite primary key for this table is (**account_id**, **region**, **hub_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|tags|JSON|
|auto_enable_controls|Bool|
|hub_arn (PK)|String|
|subscribed_at|Timestamp|
|result_metadata|JSON|