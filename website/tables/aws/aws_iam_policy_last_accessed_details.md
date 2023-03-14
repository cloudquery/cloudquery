# Table: aws_iam_policy_last_accessed_details

This table shows data for AWS IAM Policy Last Accessed Details.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html

The composite primary key for this table is (**account_id**, **arn**, **service_namespace**).

## Relations

This table depends on [aws_iam_policies](aws_iam_policies).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|arn (PK)|String|
|job_id|String|
|service_name|String|
|service_namespace (PK)|String|
|last_authenticated|Timestamp|
|last_authenticated_entity|String|
|last_authenticated_region|String|
|total_authenticated_entities|Int|
|tracked_actions_last_accessed|JSON|