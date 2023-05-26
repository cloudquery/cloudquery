# Table: heroku_inbound_rulesets

This table shows data for Heroku Inbound Rulesets.

https://devcenter.heroku.com/articles/platform-api-reference#inbound-ruleset

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|created_at|timestamp[us, tz=UTC]|
|created_by|utf8|
|rules|json|
|space|json|