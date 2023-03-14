# Table: aws_iam_roles

This table shows data for AWS IAM Roles.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_iam_roles:
  - [aws_iam_role_last_accessed_details](aws_iam_role_last_accessed_details)
  - [aws_iam_role_policies](aws_iam_role_policies)

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
|tags|JSON|
|arn|String|
|create_date|Timestamp|
|path|String|
|role_id|String|
|role_name|String|
|description|String|
|max_session_duration|Int|
|permissions_boundary|JSON|
|role_last_used|JSON|