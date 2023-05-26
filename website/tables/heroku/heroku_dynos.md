# Table: heroku_dynos

This table shows data for Heroku Dynos.

https://devcenter.heroku.com/articles/platform-api-reference#dyno

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|app|json|
|attach_url|utf8|
|command|utf8|
|created_at|timestamp[us, tz=UTC]|
|name|utf8|
|release|json|
|size|utf8|
|state|utf8|
|type|utf8|
|updated_at|timestamp[us, tz=UTC]|