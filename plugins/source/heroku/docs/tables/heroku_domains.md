# Table: heroku_domains

https://devcenter.heroku.com/articles/platform-api-reference#domain

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|acm_status|String|
|acm_status_reason|String|
|app|JSON|
|cname|String|
|created_at|Timestamp|
|hostname|String|
|kind|String|
|sni_endpoint|JSON|
|status|String|
|updated_at|Timestamp|