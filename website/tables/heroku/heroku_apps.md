# Table: heroku_apps

This table shows data for Heroku Apps.

https://devcenter.heroku.com/articles/platform-api-reference#app

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|acm|bool|
|archived_at|timestamp[us, tz=UTC]|
|build_stack|json|
|buildpack_provided_description|utf8|
|created_at|timestamp[us, tz=UTC]|
|git_url|utf8|
|internal_routing|bool|
|maintenance|bool|
|name|utf8|
|organization|json|
|owner|json|
|region|json|
|released_at|timestamp[us, tz=UTC]|
|repo_size|int64|
|slug_size|int64|
|space|json|
|stack|json|
|team|json|
|updated_at|timestamp[us, tz=UTC]|
|web_url|utf8|