# Table: heroku_dynos

https://devcenter.heroku.com/articles/platform-api-reference#dyno

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
|attach_url|String|
|command|String|
|created_at|Timestamp|
|name|String|
|release|JSON|
|size|String|
|state|String|
|type|String|
|updated_at|Timestamp|