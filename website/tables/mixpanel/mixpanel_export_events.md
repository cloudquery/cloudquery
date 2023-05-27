# Table: mixpanel_export_events

This table shows data for Mixpanel Export Events.

https://developer.mixpanel.com/reference/raw-event-export

The composite primary key for this table is (**project_id**, **time**, **distinct_id**, **event**).
It supports incremental syncs based on the **time** column.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id (PK)|int64|
|time (PK) (Incremental Key)|timestamp[us, tz=UTC]|
|distinct_id (PK)|utf8|
|event (PK)|utf8|
|properties|json|