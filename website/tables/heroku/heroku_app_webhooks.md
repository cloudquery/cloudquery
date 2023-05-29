# Table: heroku_app_webhooks

This table shows data for Heroku App Webhooks.

https://devcenter.heroku.com/articles/platform-api-reference#app-webhook

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|app|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|include|`list<item: utf8, nullable>`|
|level|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|url|`utf8`|