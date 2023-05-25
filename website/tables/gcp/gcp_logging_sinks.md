# Table: gcp_logging_sinks

This table shows data for GCP Logging Sinks.

https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.sinks#LogSink

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id (PK)|utf8|
|name (PK)|utf8|
|destination|utf8|
|filter|utf8|
|description|utf8|
|disabled|bool|
|exclusions|json|
|output_version_format|utf8|
|writer_identity|utf8|
|include_children|bool|
|create_time|timestamp[us, tz=UTC]|
|update_time|timestamp[us, tz=UTC]|