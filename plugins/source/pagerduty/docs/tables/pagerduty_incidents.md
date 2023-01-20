# Table: pagerduty_incidents

https://developer.pagerduty.com/api-reference/9d0b4b12e36f9-list-incidents

The primary key for this table is **id**.

## Relations

The following tables depend on pagerduty_incidents:
  - [pagerduty_incident_alerts](pagerduty_incident_alerts.md)
  - [pagerduty_incident_log_entries](pagerduty_incident_log_entries.md)
  - [pagerduty_incident_notes](pagerduty_incident_notes.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|html_url|String|
|created_at|Timestamp|
|last_status_change_at|Timestamp|
|type|String|
|summary|String|
|self|String|
|incident_number|Int|
|title|String|
|description|String|
|pending_actions|JSON|
|incident_key|String|
|service|JSON|
|assignments|JSON|
|acknowledgements|JSON|
|last_status_change_by|JSON|
|first_trigger_log_entry|JSON|
|escalation_policy|JSON|
|teams|JSON|
|priority|JSON|
|urgency|String|
|status|String|
|resolve_reason|JSON|
|alert_counts|JSON|
|body|JSON|
|is_mergeable|Bool|
|conference_bridge|JSON|
|assigned_via|String|
|occurrence|JSON|
|incidents_responders|JSON|
|responder_requests|JSON|