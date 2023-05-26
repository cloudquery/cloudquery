# Table: heroku_credits

This table shows data for Heroku Credits.

https://devcenter.heroku.com/articles/platform-api-reference#credit

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|amount|float64|
|balance|float64|
|created_at|timestamp[us, tz=UTC]|
|expires_at|timestamp[us, tz=UTC]|
|title|utf8|
|updated_at|timestamp[us, tz=UTC]|