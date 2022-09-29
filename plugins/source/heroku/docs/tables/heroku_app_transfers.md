# Table: heroku_app_transfers
https://devcenter.heroku.com/articles/platform-api-reference#app-transfer-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|app|JSON|
|created_at|Timestamp|
|id (PK)|String|
|owner|JSON|
|recipient|JSON|
|state|String|
|updated_at|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|