# Table: github_user_keys

This table shows data for Github User Keys.

The composite primary key for this table is (**user_id**, **id**).

## Relations

This table depends on [github_users](github_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|user_id (PK)|`int64`|
|id (PK)|`int64`|
|key|`utf8`|
|url|`utf8`|
|title|`utf8`|
|read_only|`bool`|
|verified|`bool`|
|created_at|`timestamp[us, tz=UTC]`|
|added_by|`utf8`|
|last_used|`timestamp[us, tz=UTC]`|