# Table: mixpanel_export_events

The composite primary key for this table is (**project_id**, **time**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|Int|
|time (PK)|Timestamp|
|event|String|
|properties|JSON|