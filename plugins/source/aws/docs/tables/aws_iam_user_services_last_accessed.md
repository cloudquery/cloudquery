# Table: aws_iam_user_services_last_accessed

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServiceLastAccessed.html

The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|service_name|String|
|service_namespace|String|
|last_authenticated|Timestamp|
|last_authenticated_entity|String|
|last_authenticated_region|String|
|total_authenticated_entities|Int|
|tracked_actions_last_accessed|JSON|
|entities|JSON|