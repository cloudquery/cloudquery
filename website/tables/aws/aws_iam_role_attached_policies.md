# Table: aws_iam_role_attached_policies

This table shows data for IAM Role Attached Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html

The composite primary key for this table is (**role_arn**, **policy_arn**).

## Relations

This table depends on [aws_iam_roles](aws_iam_roles).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|role_arn (PK)|String|
|policy_arn (PK)|String|
|policy_name|String|