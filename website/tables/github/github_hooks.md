# Table: github_hooks

This table shows data for Github Hooks.

The composite primary key for this table is (**org**, **id**).

## Relations

The following tables depend on github_hooks:
  - [github_hook_deliveries](github_hook_deliveries)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|url|`utf8`|
|id (PK)|`int64`|
|type|`utf8`|
|name|`utf8`|
|test_url|`utf8`|
|ping_url|`utf8`|
|last_response|`json`|
|config|`json`|
|events|`list<item: utf8, nullable>`|
|active|`bool`|