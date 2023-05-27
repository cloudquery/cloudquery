# Table: heroku_add_on_webhook_events

This table shows data for Heroku Add On Webhook Events.

https://devcenter.heroku.com/articles/platform-api-reference#add-on-webhook-event

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|include|`utf8`|
|payload|`json`|
|updated_at|`timestamp[us, tz=UTC]`|