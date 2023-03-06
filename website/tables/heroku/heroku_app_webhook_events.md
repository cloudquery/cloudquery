# Table: heroku_app_webhook_events

https://devcenter.heroku.com/articles/platform-api-reference#app-webhook-event

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created_at|Timestamp|
|include|String|
|payload|JSON|
|updated_at|Timestamp|