# Table: heroku_pipelines

This table shows data for Heroku Pipelines.

https://devcenter.heroku.com/articles/platform-api-reference#pipeline

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|created_at|timestamp[us, tz=UTC]|
|name|utf8|
|owner|json|
|updated_at|timestamp[us, tz=UTC]|