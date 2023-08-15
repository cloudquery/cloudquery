# Table: gitlab_projects_releases

This table shows data for Gitlab Projects Releases.

The composite primary key for this table is (**base_url**, **project_id**, **created_at**).

## Relations

This table depends on [gitlab_projects](gitlab_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|base_url (PK)|`utf8`|
|project_id (PK)|`int64`|
|tag_name|`utf8`|
|name|`utf8`|
|description|`utf8`|
|description_html|`utf8`|
|created_at (PK)|`timestamp[us, tz=UTC]`|
|released_at|`timestamp[us, tz=UTC]`|
|author|`json`|
|commit|`json`|
|upcoming_release|`bool`|
|commit_path|`utf8`|
|tag_path|`utf8`|
|assets|`json`|
|_links|`json`|