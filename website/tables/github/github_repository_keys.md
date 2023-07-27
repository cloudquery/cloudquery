# Table: github_repository_keys

This table shows data for Github Repository Keys.

The composite primary key for this table is (**org**, **repository_id**, **id**).

## Relations

This table depends on [github_repositories](github_repositories).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|id (PK)|`int64`|
|key|`utf8`|
|url|`utf8`|
|title|`utf8`|
|read_only|`bool`|
|verified|`bool`|
|created_at|`timestamp[us, tz=UTC]`|