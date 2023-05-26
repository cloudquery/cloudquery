# Table: heroku_regions

This table shows data for Heroku Regions.

https://devcenter.heroku.com/articles/platform-api-reference#region

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|country|utf8|
|created_at|timestamp[us, tz=UTC]|
|description|utf8|
|locale|utf8|
|name|utf8|
|private_capable|bool|
|provider|json|
|updated_at|timestamp[us, tz=UTC]|