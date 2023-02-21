# Table: gcp_videotranscoder_job_templates

https://cloud.google.com/transcoder/docs/reference/rest/v1/projects.locations.jobTemplates

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
|config|JSON|
|labels|JSON|