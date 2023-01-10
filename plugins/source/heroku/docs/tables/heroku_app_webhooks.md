# Table: heroku_app_webhooks

https://devcenter.heroku.com/articles/platform-api-reference#app-webhook

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|app|JSON|
|created_at|Timestamp|
|include|StringArray|
|level|String|
|updated_at|Timestamp|
|url|String|