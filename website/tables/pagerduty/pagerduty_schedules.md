# Table: pagerduty_schedules

This table shows data for PagerDuty Schedules.

https://developer.pagerduty.com/api-reference/846ecf84402bb-list-schedules

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
|time_zone|`utf8`|
|description|`utf8`|
|escalation_policies|`json`|
|users|`json`|
|teams|`json`|
|schedule_layers|`json`|
|override_subschedule|`json`|
|final_schedule|`json`|