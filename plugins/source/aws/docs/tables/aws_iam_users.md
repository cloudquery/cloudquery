# Table: aws_iam_users

https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html

The composite primary key for this table is (**id**, **account_id**).

## Relations

The following tables depend on aws_iam_users:
  - [aws_iam_user_access_keys](aws_iam_user_access_keys.md)
  - [aws_iam_user_groups](aws_iam_user_groups.md)
  - [aws_iam_user_attached_policies](aws_iam_user_attached_policies.md)
  - [aws_iam_user_policies](aws_iam_user_policies.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|arn|String|
|id (PK)|String|
|account_id (PK)|String|
|create_date|Timestamp|
|path|String|
|user_name|String|
|password_last_used|Timestamp|
|permissions_boundary|JSON|
|tags|JSON|