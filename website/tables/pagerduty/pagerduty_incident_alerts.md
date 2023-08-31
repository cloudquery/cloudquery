# Table: pagerduty_incident_alerts

This table shows data for PagerDuty Incident Alerts.

https://developer.pagerduty.com/api-reference/4bc42e7ac0c59-list-alerts-for-an-incident

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_incidents](pagerduty_incidents).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|html_url|`utf8`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|status|`utf8`|
|alert_key|`utf8`|
|service|`json`|
|body|`json`|
|incident|`json`|
|suppressed|`bool`|
|severity|`utf8`|
|integration|`json`|