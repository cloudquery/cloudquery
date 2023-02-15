# Table: github_releases

The composite primary key for this table is (**org**, **repository_id**, **id**).

## Relations

This table depends on [github_repositories](github_repositories.md).

The following tables depend on github_releases:
  - [github_release_assets](github_release_assets.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|repository_id (PK)|Int|
|tag_name|String|
|target_commitish|String|
|name|String|
|body|String|
|draft|Bool|
|prerelease|Bool|
|make_latest|String|
|discussion_category_name|String|
|generate_release_notes|Bool|
|id (PK)|Int|
|created_at|Timestamp|
|published_at|Timestamp|
|url|String|
|html_url|String|
|assets_url|String|
|assets|JSON|
|upload_url|String|
|zipball_url|String|
|tarball_url|String|
|author|JSON|
|node_id|String|