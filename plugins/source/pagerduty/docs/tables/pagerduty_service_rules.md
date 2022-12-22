# Table: pagerduty_service_rules

https://developer.pagerduty.com/api-reference/d69ad7f1ec900-list-service-s-event-rules

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_services](pagerduty_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|self|String|
|disabled|Bool|
|conditions|JSON|
|time_frame|JSON|
|position|Int|
|actions|JSON|