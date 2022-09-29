# Table: heroku_dyno_sizes
https://devcenter.heroku.com/articles/platform-api-reference#dyno-size-attributes

The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|compute|Int|
|cost|JSON|
|dedicated|Bool|
|dyno_units|Int|
|id (PK)|String|
|memory|Float|
|name|String|
|private_space_only|Bool|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|