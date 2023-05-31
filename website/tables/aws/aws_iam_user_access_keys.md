# Table: aws_iam_user_access_keys

This table shows data for IAM User Access Keys.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AccessKeyMetadata.html

The composite primary key for this table is (**account_id**, **user_arn**, **access_key_id**).

## Relations

This table depends on [aws_iam_users](aws_iam_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|user_arn (PK)|`utf8`|
|access_key_id (PK)|`utf8`|
|user_id|`utf8`|
|last_used|`timestamp[us, tz=UTC]`|
|last_used_service_name|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|user_name|`utf8`|
|last_rotated|`timestamp[us, tz=UTC]`|