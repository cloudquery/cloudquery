# Table: pagerduty_incident_notes

This table shows data for PagerDuty Incident Notes.

https://developer.pagerduty.com/api-reference/a1ac30885eb7a-list-notes-for-an-incident

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
|user|`json`|
|content|`utf8`|