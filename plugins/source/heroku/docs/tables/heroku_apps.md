# Table: heroku_apps

https://devcenter.heroku.com/articles/platform-api-reference#app

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|acm|Bool|
|archived_at|Timestamp|
|build_stack|JSON|
|buildpack_provided_description|String|
|created_at|Timestamp|
|git_url|String|
|id (PK)|String|
|internal_routing|Bool|
|maintenance|Bool|
|name|String|
|organization|JSON|
|owner|JSON|
|region|JSON|
|released_at|Timestamp|
|repo_size|Int|
|slug_size|Int|
|space|JSON|
|stack|JSON|
|team|JSON|
|updated_at|Timestamp|
|web_url|String|