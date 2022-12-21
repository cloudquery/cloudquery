# Table: pagerduty_maintenance_windows

https://developer.pagerduty.com/api-reference/4c0936c241cbb-list-maintenance-windows

The primary key for this table is **id**.

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
|sequence_number|Int|
|start_time|Timestamp|
|end_time|Timestamp|
|description|String|
|services|JSON|
|teams|JSON|
|created_by|JSON|