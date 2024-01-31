# Table: aws_iam_role_policies

This table shows data for IAM Role Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetRolePolicy.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **role_arn**, **policy_name**).
## Relations

This table depends on [aws_iam_roles](aws_iam_roles.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|role_arn|`utf8`|
|policy_document|`json`|
|policy_name|`utf8`|
|role_name|`utf8`|