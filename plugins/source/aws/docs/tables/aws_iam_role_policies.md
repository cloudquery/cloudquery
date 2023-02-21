# Table: aws_iam_role_policies

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html

The composite primary key for this table is (**account_id**, **role_arn**, **policy_name**).

## Relations

This table depends on [aws_iam_roles](aws_iam_roles.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|role_arn (PK)|String|
|policy_name (PK)|String|
|policy_document|JSON|
|role_name|String|
|result_metadata|JSON|