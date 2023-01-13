# Table: pagerduty_services

https://developer.pagerduty.com/api-reference/e960cca205c0f-list-services

The primary key for this table is **id**.

## Relations

The following tables depend on pagerduty_services:
  - [pagerduty_service_rules](pagerduty_service_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|dependencies|JSON|
|id (PK)|String|
|html_url|String|
|created_at|Timestamp|
|last_incident_timestamp|Timestamp|
|type|String|
|summary|String|
|self|String|
|name|String|
|description|String|
|auto_resolve_timeout|Int|
|acknowledgement_timeout|Int|
|status|String|
|integrations|JSON|
|escalation_policy|JSON|
|teams|JSON|
|incident_urgency_rule|JSON|
|support_hours|JSON|
|scheduled_actions|JSON|
|alert_creation|String|
|alert_grouping|String|
|alert_grouping_timeout|Int|
|alert_grouping_parameters|JSON|
|response_play|JSON|
|addons|JSON|