# Table: gcp_logging_sinks

https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.sinks#LogSink

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|name (PK)|String|
|destination|String|
|filter|String|
|description|String|
|disabled|Bool|
|exclusions|JSON|
|output_version_format|String|
|writer_identity|String|
|include_children|Bool|
|create_time|Timestamp|
|update_time|Timestamp|