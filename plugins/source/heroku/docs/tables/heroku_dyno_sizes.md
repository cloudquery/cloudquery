# Table: heroku_dyno_sizes

https://devcenter.heroku.com/articles/platform-api-reference#dyno-size

The primary key for this table is **id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|compute|Int|
|cost|JSON|
|dedicated|Bool|
|dyno_units|Int|
|id (PK)|String|
|memory|Float|
|name|String|
|private_space_only|Bool|