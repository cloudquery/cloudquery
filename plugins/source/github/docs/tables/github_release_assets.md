# Table: github_release_assets

The composite primary key for this table is (**org**, **repository_id**, **id**).

## Relations

This table depends on [github_releases](github_releases.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|repository_id (PK)|Int|
|id (PK)|Int|
|url|String|
|name|String|
|label|String|
|state|String|
|content_type|String|
|size|Int|
|download_count|Int|
|created_at|Timestamp|
|updated_at|Timestamp|
|browser_download_url|String|
|uploader|JSON|
|node_id|String|