# Table: fastly_account_users

This table shows data for Fastly Account Users.

https://developer.fastly.com/reference/api/account/user/

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|customer_id|`utf8`|
|deleted_at|`timestamp[us, tz=UTC]`|
|email_hash|`utf8`|
|limit_services|`bool`|
|locked|`bool`|
|login|`utf8`|
|name|`utf8`|
|require_new_password|`bool`|
|role|`utf8`|
|two_factor_auth_enabled|`bool`|
|two_factor_setup_required|`bool`|
|updated_at|`timestamp[us, tz=UTC]`|