# Table: aws_cloudtrail_channels

This table shows data for AWS CloudTrail Channels.

https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_GetChannel.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|channel_arn|`utf8`|
|destinations|`json`|
|ingestion_status|`json`|
|name|`utf8`|
|source|`utf8`|
|source_config|`json`|