# Table: heroku_pipeline_couplings

This table shows data for Heroku Pipeline Couplings.

https://devcenter.heroku.com/articles/platform-api-reference#pipeline-coupling

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|app|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|pipeline|`json`|
|stage|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|