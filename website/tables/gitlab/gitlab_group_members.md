# Table: gitlab_group_members

This table shows data for Gitlab Group Members.

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
|created_at|`timestamp[us, tz=UTC]`|
|expires_at|`timestamp[us, tz=UTC]`|
|access_level|`int64`|
|email|`utf8`|
|group_saml_identity|`json`|