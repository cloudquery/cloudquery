# Table: aws_cloudwatchlogs_resource_policies

https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ResourcePolicy.html

The composite primary key for this table is (**account_id**, **region**, **policy_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|policy_name (PK)|String|
|last_updated_time|Int|
|policy_document|String|