# Table: pagerduty_services

This table shows data for PagerDuty Services.

https://developer.pagerduty.com/api-reference/e960cca205c0f-list-services

The primary key for this table is **id**.

## Relations

The following tables depend on pagerduty_services:
  - [pagerduty_service_rules](pagerduty_service_rules)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|dependencies|`json`|
|id (PK)|`utf8`|
|html_url|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|last_incident_timestamp|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|name|`utf8`|
|description|`utf8`|
|auto_resolve_timeout|`int64`|
|acknowledgement_timeout|`int64`|
|status|`utf8`|
|integrations|`json`|
|escalation_policy|`json`|
|teams|`json`|
|incident_urgency_rule|`json`|
|support_hours|`json`|
|scheduled_actions|`json`|
|alert_creation|`utf8`|
|alert_grouping|`utf8`|
|alert_grouping_timeout|`int64`|
|alert_grouping_parameters|`json`|
|response_play|`json`|
|addons|`json`|