# Table: gitlab_project_branches

This table shows data for Gitlab Project Branches.

The composite primary key for this table is (**base_url**, **project_id**, **name**).

## Relations

This table depends on [gitlab_projects](gitlab_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|base_url (PK)|`utf8`|
|project_id (PK)|`int64`|
|commit|`json`|
|name (PK)|`utf8`|
|protected|`bool`|
|merged|`bool`|
|default|`bool`|
|can_push|`bool`|
|developers_can_push|`bool`|
|developers_can_merge|`bool`|
|web_url|`utf8`|