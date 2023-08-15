# Table: pagerduty_incident_log_entries

This table shows data for PagerDuty Incident Log Entries.

https://developer.pagerduty.com/api-reference/367602cbc1c28-list-log-entries-for-an-incident

The primary key for this table is **id**.

## Relations

This table depends on [pagerduty_incidents](pagerduty_incidents).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|html_url|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|agent|`json`|
|channel|`json`|
|teams|`json`|
|contexts|`json`|
|acknowledgement_timeout|`int64`|
|event_details|`json`|
|assignees|`json`|
|incident|`json`|
|service|`json`|
|user|`json`|