# Table: aws_iam_user_access_keys

This table shows data for IAM User Access Keys.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AccessKeyMetadata.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **user_arn**, **access_key_id**).
## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|user_arn|`utf8`|
|access_key_id|`utf8`|
|user_id|`utf8`|
|last_used|`timestamp[us, tz=UTC]`|
|last_used_service_name|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|user_name|`utf8`|
|last_rotated|`timestamp[us, tz=UTC]`|