# Table: pagerduty_oncalls

This table shows data for PagerDuty Oncalls.

https://developer.pagerduty.com/api-reference/3a6b910f11050-list-all-of-the-on-calls

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|user|`json`|
|schedule|`json`|
|escalation_policy|`json`|
|escalation_level|`int64`|
|start|`utf8`|
|end|`utf8`|