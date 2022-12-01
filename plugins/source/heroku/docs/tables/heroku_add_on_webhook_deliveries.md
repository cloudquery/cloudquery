# Table: heroku_add_on_webhook_deliveries

https://devcenter.heroku.com/articles/platform-api-reference#add-on-webhook-delivery

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|created_at|Timestamp|
|event|JSON|
|id (PK)|String|
|last_attempt|JSON|
|next_attempt_at|Timestamp|
|num_attempts|Int|
|status|String|
|updated_at|Timestamp|
|webhook|JSON|