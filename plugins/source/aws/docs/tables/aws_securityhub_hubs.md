# Table: aws_securityhub_hubs

This table shows data for AWS Security Hub Hubs.

https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeHub.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **hub_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|auto_enable_controls|`bool`|
|control_finding_generator|`utf8`|
|hub_arn|`utf8`|
|subscribed_at|`timestamp[us, tz=UTC]`|