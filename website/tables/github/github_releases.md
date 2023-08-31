# Table: github_releases

This table shows data for Github Releases.

The composite primary key for this table is (**org**, **repository_id**, **id**).

## Relations

This table depends on [github_repositories](github_repositories).

The following tables depend on github_releases:
  - [github_release_assets](github_release_assets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|org (PK)|`utf8`|
|repository_id (PK)|`int64`|
|tag_name|`utf8`|
|target_commitish|`utf8`|
|name|`utf8`|
|body|`utf8`|
|draft|`bool`|
|prerelease|`bool`|
|make_latest|`utf8`|
|discussion_category_name|`utf8`|
|generate_release_notes|`bool`|
|id (PK)|`int64`|
|created_at|`timestamp[us, tz=UTC]`|
|published_at|`timestamp[us, tz=UTC]`|
|url|`utf8`|
|html_url|`utf8`|
|assets_url|`utf8`|
|assets|`json`|
|upload_url|`utf8`|
|zipball_url|`utf8`|
|tarball_url|`utf8`|
|author|`json`|
|node_id|`utf8`|