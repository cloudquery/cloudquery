# Table: gitlab_group_billable_members

This table shows data for Gitlab Group Billable Members.

The composite primary key for this table is (**base_url**, **group_id**, **id**).

## Relations

This table depends on [gitlab_groups](gitlab_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|base_url (PK)|`utf8`|
|group_id (PK)|`int64`|
|id (PK)|`int64`|
|username|`utf8`|
|name|`utf8`|
|state|`utf8`|
|avatar_url|`utf8`|
|web_url|`utf8`|
|email|`utf8`|
|last_activity_on|`timestamp[us, tz=UTC]`|
|membership_type|`utf8`|
|removable|`bool`|
|created_at|`timestamp[us, tz=UTC]`|
|is_last_owner|`bool`|
|last_login_at|`timestamp[us, tz=UTC]`|