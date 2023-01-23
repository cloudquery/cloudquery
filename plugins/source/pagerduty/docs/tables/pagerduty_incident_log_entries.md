# Table: pagerduty_incident_log_entries

https://developer.pagerduty.com/api-reference/367602cbc1c28-list-log-entries-for-an-incident

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_incidents](pagerduty_incidents.md).

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
|created_at|Timestamp|
|agent|JSON|
|channel|JSON|
|teams|JSON|
|contexts|JSON|
|acknowledgement_timeout|Int|
|event_details|JSON|
|assignees|JSON|
|incident|JSON|
|service|JSON|
|user|JSON|