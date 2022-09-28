# Table: heroku_releases
https://devcenter.heroku.com/articles/platform-api-reference#release-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
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
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|