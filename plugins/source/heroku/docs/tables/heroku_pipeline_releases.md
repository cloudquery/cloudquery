# Table: heroku_pipeline_releases

https://devcenter.heroku.com/articles/platform-api-reference#pipeline-release

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|addon_plan_names|StringArray|
|app|JSON|
|created_at|Timestamp|
|current|Bool|
|description|String|
|id (PK)|String|
|output_stream_url|String|
|slug|JSON|
|status|String|
|updated_at|Timestamp|
|user|JSON|
|version|Int|