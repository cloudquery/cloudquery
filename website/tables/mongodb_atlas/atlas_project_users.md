# Table: atlas_project_users

This table shows data for Atlas Project Users.

The composite primary key for this table is (**project_id**, **id**).

## Relations

This table depends on [atlas_projects](atlas_projects.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|country|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|email_address|`utf8`|
|first_name|`utf8`|
|id (PK)|`utf8`|
|last_auth|`timestamp[us, tz=UTC]`|
|last_name|`utf8`|
|links|`json`|
|mobile_number|`utf8`|
|password|`utf8`|
|roles|`json`|
|team_ids|`list<item: utf8, nullable>`|
|username|`utf8`|