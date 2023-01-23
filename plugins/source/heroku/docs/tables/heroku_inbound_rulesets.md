# Table: heroku_inbound_rulesets

https://devcenter.heroku.com/articles/platform-api-reference#inbound-ruleset

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created_at|Timestamp|
|created_by|String|
|rules|JSON|
|space|JSON|