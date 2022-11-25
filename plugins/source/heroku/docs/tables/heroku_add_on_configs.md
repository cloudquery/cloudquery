# Table: heroku_add_on_configs

https://devcenter.heroku.com/articles/platform-api-reference#add-on-config

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|name|String|
|value|String|