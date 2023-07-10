# Table: pagerduty_ruleset_rules

This table shows data for PagerDuty Ruleset Rules.

https://developer.pagerduty.com/api-reference/c39605f86c5b7-list-event-rules

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_rulesets](pagerduty_rulesets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|self|`utf8`|
|position|`int64`|
|disabled|`bool`|
|conditions|`json`|
|actions|`json`|
|ruleset|`json`|
|catch_all|`bool`|
|time_frame|`json`|