# Table: aws_iam_role_services_last_accessed

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html

The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_iam_roles](aws_iam_roles.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|resource_arn|String|
|job_id|String|
|service_name|String|
|service_namespace|String|
|last_authenticated|Timestamp|
|last_authenticated_entity|String|
|last_authenticated_region|String|
|total_authenticated_entities|Int|
|tracked_actions_last_accessed|JSON|
|entities|JSON|