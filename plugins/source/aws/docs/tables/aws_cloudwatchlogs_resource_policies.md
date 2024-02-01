# Table: aws_cloudwatchlogs_resource_policies

This table shows data for Cloudwatchlogs Resource Policies.

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ResourcePolicy.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **policy_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|policy_name|`utf8`|
|policy_document|`json`|
|last_updated_time|`int64`|