# Table: gitlab_project_members

This table shows data for Gitlab Project Members.

The composite primary key for this table is (**base_url**, **project_id**, **id**).

## Relations

This table depends on [gitlab_projects](gitlab_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|base_url (PK)|`utf8`|
|project_id (PK)|`int64`|
|id (PK)|`int64`|
|username|`utf8`|
|email|`utf8`|
|name|`utf8`|
|state|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|expires_at|`timestamp[us, tz=UTC]`|
|access_level|`int64`|
|web_url|`utf8`|
|avatar_url|`utf8`|