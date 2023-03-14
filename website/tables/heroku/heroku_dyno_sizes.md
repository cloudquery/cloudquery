# Table: heroku_dyno_sizes

This table shows data for Heroku Dyno Sizes.

https://devcenter.heroku.com/articles/platform-api-reference#dyno-size

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|compute|Int|
|cost|JSON|
|dedicated|Bool|
|dyno_units|Int|
|memory|Float|
|name|String|
|private_space_only|Bool|