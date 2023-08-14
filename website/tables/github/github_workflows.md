# Table: github_workflows

This table shows data for Github Workflows.

The composite primary key for this table is (**org**, **repository_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|contents|`utf8`|
|id (PK)|`int64`|
|node_id|`utf8`|
|name|`utf8`|
|path|`utf8`|
|state|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|url|`utf8`|
|html_url|`utf8`|
|badge_url|`utf8`|