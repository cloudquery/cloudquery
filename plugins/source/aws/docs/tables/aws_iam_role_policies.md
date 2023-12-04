# Table: aws_iam_role_policies

This table shows data for IAM Role Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html

The composite primary key for this table is (**account_id**, **role_arn**, **policy_name**).

## Relations

This table depends on [aws_iam_roles](aws_iam_roles.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|role_arn (PK)|`utf8`|
|policy_document|`json`|
|policy_name (PK)|`utf8`|
|role_name|`utf8`|