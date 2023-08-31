# Table: aws_cloudwatchlogs_resource_policies

This table shows data for Cloudwatchlogs Resource Policies.

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ResourcePolicy.html

The composite primary key for this table is (**account_id**, **region**, **policy_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|policy_name (PK)|`utf8`|
|policy_document|`json`|
|last_updated_time|`int64`|