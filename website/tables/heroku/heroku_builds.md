# Table: heroku_builds

This table shows data for Heroku Builds.

https://devcenter.heroku.com/articles/platform-api-reference#build

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|app|JSON|
|buildpacks|JSON|
|created_at|Timestamp|
|output_stream_url|String|
|release|JSON|
|slug|JSON|
|source_blob|JSON|
|stack|String|
|status|String|
|updated_at|Timestamp|
|user|JSON|