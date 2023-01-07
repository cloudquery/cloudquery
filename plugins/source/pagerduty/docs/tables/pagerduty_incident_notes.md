# Table: pagerduty_incident_notes

https://developer.pagerduty.com/api-reference/a1ac30885eb7a-list-notes-for-an-incident

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
|created_at|Timestamp|
|user|JSON|
|content|String|