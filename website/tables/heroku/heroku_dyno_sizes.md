# Table: heroku_dyno_sizes

This table shows data for Heroku Dyno Sizes.

https://devcenter.heroku.com/articles/platform-api-reference#dyno-size

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|compute|int64|
|cost|json|
|dedicated|bool|
|dyno_units|int64|
|memory|float64|
|name|utf8|
|private_space_only|bool|