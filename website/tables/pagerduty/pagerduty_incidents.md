# Table: pagerduty_incidents

This table shows data for PagerDuty Incidents.

https://developer.pagerduty.com/api-reference/9d0b4b12e36f9-list-incidents

The primary key for this table is **id**.

## Relations

The following tables depend on pagerduty_incidents:
  - [pagerduty_incident_alerts](pagerduty_incident_alerts)
  - [pagerduty_incident_log_entries](pagerduty_incident_log_entries)
  - [pagerduty_incident_notes](pagerduty_incident_notes)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|html_url|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|last_status_change_at|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|incident_number|`int64`|
|title|`utf8`|
|description|`utf8`|
|pending_actions|`json`|
|incident_key|`utf8`|
|service|`json`|
|assignments|`json`|
|acknowledgements|`json`|
|last_status_change_by|`json`|
|first_trigger_log_entry|`json`|
|escalation_policy|`json`|
|teams|`json`|
|priority|`json`|
|urgency|`utf8`|
|status|`utf8`|
|resolve_reason|`json`|
|alert_counts|`json`|
|body|`json`|
|is_mergeable|`bool`|
|conference_bridge|`json`|
|assigned_via|`utf8`|
|occurrence|`json`|
|incidents_responders|`json`|
|responder_requests|`json`|