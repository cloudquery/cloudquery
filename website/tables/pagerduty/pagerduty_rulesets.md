# Table: pagerduty_rulesets

This table shows data for PagerDuty Rulesets.

https://developer.pagerduty.com/api-reference/633f1ecb6c03b-list-rulesets

The primary key for this table is **id**.

## Relations

The following tables depend on pagerduty_rulesets:
  - [pagerduty_ruleset_rules](pagerduty_ruleset_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|type|`utf8`|
|self|`utf8`|
|routing_keys|`list<item: utf8, nullable>`|
|creator|`json`|
|updated_at|`utf8`|
|updater|`json`|
|team|`json`|