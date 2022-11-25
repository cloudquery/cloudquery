# Table: heroku_pipeline_couplings

https://devcenter.heroku.com/articles/platform-api-reference#pipeline-coupling

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|app|JSON|
|created_at|Timestamp|
|id (PK)|String|
|pipeline|JSON|
|stage|String|
|updated_at|Timestamp|