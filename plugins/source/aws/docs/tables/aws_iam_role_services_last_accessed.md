# Table: aws_iam_role_services_last_accessed

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html

The composite primary key for this table is (**job_id**, **resource_arn**, **service_name**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|job_id (PK)|String|
|resource_arn (PK)|String|
|service_name (PK)|String|
|service_namespace|String|
|last_authenticated|Timestamp|
|last_authenticated_entity|String|
|last_authenticated_region|String|
|total_authenticated_entities|Int|
|tracked_actions_last_accessed|JSON|
|entities|JSON|