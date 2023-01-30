# Table: aws_iam_user_groups

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html

The composite primary key for this table is (**user_id**, **arn**).

## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|user_arn|String|
|user_id (PK)|String|
|arn (PK)|String|
|create_date|Timestamp|
|group_id|String|
|group_name|String|
|path|String|