# Table: pagerduty_rulesets

https://developer.pagerduty.com/api-reference/633f1ecb6c03b-list-rulesets

The primary key for this table is **id**.

## Relations

The following tables depend on pagerduty_rulesets:
  - [pagerduty_ruleset_rules](pagerduty_ruleset_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|type|String|
|self|String|
|routing_keys|StringArray|
|created_at|Timestamp|
|creator|JSON|
|updated_at|String|
|updater|JSON|
|team|JSON|