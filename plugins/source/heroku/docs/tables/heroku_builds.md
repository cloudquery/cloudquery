# Table: heroku_builds
https://devcenter.heroku.com/articles/platform-api-reference#build-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|app|JSON|
|buildpacks|JSON|
|created_at|Timestamp|
|id (PK)|String|
|output_stream_url|String|
|release|JSON|
|slug|JSON|
|source_blob|JSON|
|stack|String|
|status|String|
|updated_at|Timestamp|
|user|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|