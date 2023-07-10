# Table: pagerduty_service_rules

This table shows data for PagerDuty Service Rules.

https://developer.pagerduty.com/api-reference/d69ad7f1ec900-list-service-s-event-rules

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_services](pagerduty_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|self|`utf8`|
|disabled|`bool`|
|conditions|`json`|
|time_frame|`json`|
|position|`int64`|
|actions|`json`|