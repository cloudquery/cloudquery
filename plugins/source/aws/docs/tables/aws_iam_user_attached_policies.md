# Table: aws_iam_user_attached_policies

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html

The composite primary key for this table is (**account_id**, **user_arn**, **policy_name**).

## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|user_arn (PK)|String|
|policy_name (PK)|String|
|user_id|String|
|policy_arn|String|