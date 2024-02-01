# Table: aws_appstream_users

This table shows data for Amazon AppStream Users.

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_User.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|authentication_type|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|enabled|`bool`|
|first_name|`utf8`|
|last_name|`utf8`|
|status|`utf8`|
|user_name|`utf8`|