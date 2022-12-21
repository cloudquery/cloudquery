# Table: pagerduty_schedules

https://developer.pagerduty.com/api-reference/846ecf84402bb-list-schedules

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|type|String|
|summary|String|
|self|String|
|html_url|String|
|name|String|
|time_zone|String|
|description|String|
|escalation_policies|JSON|
|users|JSON|
|teams|JSON|
|schedule_layers|JSON|
|override_subschedule|JSON|
|final_schedule|JSON|