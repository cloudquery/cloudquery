# Table: pagerduty_maintenance_windows

This table shows data for PagerDuty Maintenance Windows.

https://developer.pagerduty.com/api-reference/4c0936c241cbb-list-maintenance-windows

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|html_url|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|summary|`utf8`|
|self|`utf8`|
|sequence_number|`int64`|
|description|`utf8`|
|services|`json`|
|teams|`json`|
|created_by|`json`|