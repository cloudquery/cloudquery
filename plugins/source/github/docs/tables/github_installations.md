# Table: github_installations

This table shows data for Github Installations.

The composite primary key for this table is (**org**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|id (PK)|`int64`|
|node_id|`utf8`|
|app_id|`int64`|
|app_slug|`utf8`|
|target_id|`int64`|
|account|`json`|
|access_tokens_url|`utf8`|
|repositories_url|`utf8`|
|html_url|`utf8`|
|target_type|`utf8`|
|single_file_name|`utf8`|
|repository_selection|`utf8`|
|events|`list<item: utf8, nullable>`|
|single_file_paths|`list<item: utf8, nullable>`|
|permissions|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|has_multiple_single_files|`bool`|
|suspended_by|`json`|
|suspended_at|`timestamp[us, tz=UTC]`|