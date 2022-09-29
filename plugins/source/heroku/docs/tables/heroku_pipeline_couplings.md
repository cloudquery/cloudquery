# Table: heroku_pipeline_couplings
https://devcenter.heroku.com/articles/platform-api-reference#pipeline-coupling-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|app|JSON|
|created_at|Timestamp|
|id (PK)|String|
|pipeline|JSON|
|stage|String|
|updated_at|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|