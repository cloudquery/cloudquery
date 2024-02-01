# Table: aws_iam_user_groups

This table shows data for IAM User Groups.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **user_arn**, **arn**).
## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|user_arn|`utf8`|
|user_id|`utf8`|
|arn|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|group_id|`utf8`|
|group_name|`utf8`|
|path|`utf8`|