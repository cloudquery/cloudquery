# Table: aws_iam_user_groups

This table shows data for IAM User Groups.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html

The composite primary key for this table is (**account_id**, **user_arn**, **arn**).

## Relations

This table depends on [aws_iam_users](aws_iam_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|user_arn (PK)|`utf8`|
|user_id|`utf8`|
|arn (PK)|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|group_id|`utf8`|
|group_name|`utf8`|
|path|`utf8`|