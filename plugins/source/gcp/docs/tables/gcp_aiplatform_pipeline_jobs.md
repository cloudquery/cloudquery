# Table: gcp_aiplatform_pipeline_jobs

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.pipelineJobs#PipelineJob

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_pipeline_locations](gcp_aiplatform_pipeline_locations.md).

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
|create_time|Timestamp|
|start_time|Timestamp|
|end_time|Timestamp|
|update_time|Timestamp|
|pipeline_spec|JSON|
|state|String|
|job_detail|JSON|
|error|JSON|
|labels|JSON|
|runtime_config|JSON|
|encryption_spec|JSON|
|service_account|String|
|network|String|
|template_uri|String|
|template_metadata|JSON|