# Table: gcp_videotranscoder_jobs

https://cloud.google.com/transcoder/docs/reference/rest/v1/projects.locations.jobs

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|input_uri|String|
|output_uri|String|
|state|String|
|create_time|Timestamp|
|start_time|Timestamp|
|end_time|Timestamp|
|ttl_after_completion_days|Int|
|labels|JSON|
|error|JSON|