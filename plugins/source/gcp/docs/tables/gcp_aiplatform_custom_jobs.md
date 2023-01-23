# Table: gcp_aiplatform_custom_jobs

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.customJobs#CustomJob

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_job_locations](gcp_aiplatform_job_locations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|display_name|String|
|job_spec|JSON|
|state|String|
|create_time|Timestamp|
|start_time|Timestamp|
|end_time|Timestamp|
|update_time|Timestamp|
|error|JSON|
|labels|JSON|
|encryption_spec|JSON|
|web_access_uris|JSON|