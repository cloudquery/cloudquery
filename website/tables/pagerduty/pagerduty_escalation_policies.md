# Table: pagerduty_escalation_policies

This table shows data for PagerDuty Escalation Policies.

https://developer.pagerduty.com/api-reference/51b21014a4f5a-list-escalation-policies

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|html_url|`utf8`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|name|`utf8`|
|escalation_rules|`json`|
|services|`json`|
|num_loops|`int64`|
|teams|`json`|
|description|`utf8`|
|on_call_handoff_notifications|`utf8`|