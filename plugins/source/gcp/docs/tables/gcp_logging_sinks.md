# Table: gcp_logging_sinks


The primary key for this table is **name**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|name (PK)|String|
|destination|String|
|filter|String|
|description|String|
|disabled|Bool|
|exclusions|JSON|
|output_version_format|Int|
|writer_identity|String|
|include_children|Bool|
|create_time|Timestamp|
|update_time|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|