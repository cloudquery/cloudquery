# Table: github_release_assets

This table shows data for Github Release Assets.

The composite primary key for this table is (**org**, **repository_id**, **id**).

## Relations

This table depends on [github_releases](github_releases).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|id (PK)|`int64`|
|url|`utf8`|
|name|`utf8`|
|label|`utf8`|
|state|`utf8`|
|content_type|`utf8`|
|size|`int64`|
|download_count|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|browser_download_url|`utf8`|
|uploader|`json`|
|node_id|`utf8`|