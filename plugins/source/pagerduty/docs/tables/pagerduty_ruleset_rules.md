# Table: pagerduty_ruleset_rules

https://developer.pagerduty.com/api-reference/c39605f86c5b7-list-event-rules

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_rulesets](pagerduty_rulesets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|self|String|
|position|Int|
|disabled|Bool|
|conditions|JSON|
|actions|JSON|
|ruleset|JSON|
|catch_all|Bool|
|time_frame|JSON|