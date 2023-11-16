# Table: aws_appstream_users

This table shows data for Amazon AppStream Users.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_User.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|authentication_type|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|enabled|`bool`|
|first_name|`utf8`|
|last_name|`utf8`|
|status|`utf8`|
|user_name|`utf8`|