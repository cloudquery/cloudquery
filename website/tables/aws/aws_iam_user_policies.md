# Table: aws_iam_user_policies

This table shows data for IAM User Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUserPolicy.html

The composite primary key for this table is (**account_id**, **user_arn**, **policy_name**).

## Relations

This table depends on [aws_iam_users](aws_iam_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|user_arn (PK)|utf8|
|user_id|utf8|
|policy_document|json|
|policy_name (PK)|utf8|
|user_name|utf8|
|result_metadata|json|