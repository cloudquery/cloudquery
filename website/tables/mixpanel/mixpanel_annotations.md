# Table: mixpanel_annotations

This table shows data for Mixpanel Annotations.

https://developer.mixpanel.com/reference/list-all-annotations-for-project

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id|int64|
|date|timestamp[us, tz=UTC]|
|description|utf8|
|id (PK)|int64|