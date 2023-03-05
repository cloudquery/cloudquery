# Table: mixpanel_export_events

https://developer.mixpanel.com/reference/raw-event-export

The composite primary key for this table is (**project_id**, **time**, **distinct_id**, **event**).
It supports incremental syncs based on the **time** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|Int|
|time (PK) (Incremental Key)|Timestamp|
|distinct_id (PK)|String|
|event (PK)|String|
|properties|JSON|