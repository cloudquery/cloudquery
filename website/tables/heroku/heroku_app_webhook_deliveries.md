# Table: heroku_app_webhook_deliveries

This table shows data for Heroku App Webhook Deliveries.

https://devcenter.heroku.com/articles/platform-api-reference#app-webhook-delivery

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
|event|`json`|
|last_attempt|`json`|
|next_attempt_at|`timestamp[us, tz=UTC]`|
|num_attempts|`int64`|
|status|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|
|webhook|`json`|