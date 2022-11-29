# Table: aws_iam_roles

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_iam_roles:
  - [aws_iam_role_policies](aws_iam_role_policies.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|policies|JSON|
|id (PK)|String|
|assume_role_policy_document|JSON|
|arn|String|
|create_date|Timestamp|
|path|String|
|role_name|String|
|description|String|
|max_session_duration|Int|
|permissions_boundary|JSON|
|role_last_used|JSON|
|tags|JSON|