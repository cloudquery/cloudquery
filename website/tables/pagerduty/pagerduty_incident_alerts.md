# Table: pagerduty_incident_alerts

https://developer.pagerduty.com/api-reference/4bc42e7ac0c59-list-alerts-for-an-incident

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_incidents](pagerduty_incidents).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created_at|Timestamp|
|html_url|String|
|type|String|
|summary|String|
|self|String|
|status|String|
|alert_key|String|
|service|JSON|
|body|JSON|
|incident|JSON|
|suppressed|Bool|
|severity|String|
|integration|JSON|