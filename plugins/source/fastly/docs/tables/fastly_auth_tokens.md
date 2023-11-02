# Table: fastly_auth_tokens

This table shows data for Fastly Auth Tokens.

https://developer.fastly.com/reference/api/auth-tokens/

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|ip|`utf8`|
|access_token|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|expires_at|`timestamp[us, tz=UTC]`|
|last_used_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|scope|`utf8`|
|services|`list<item: utf8, nullable>`|
|user_id|`utf8`|