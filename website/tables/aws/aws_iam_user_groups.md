# Table: aws_iam_user_groups

This table shows data for IAM User Groups.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html

The composite primary key for this table is (**account_id**, **user_arn**, **arn**).

## Relations

This table depends on [aws_iam_users](aws_iam_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|user_arn (PK)|String|
|user_id|String|
|arn (PK)|String|
|create_date|Timestamp|
|group_id|String|
|group_name|String|
|path|String|