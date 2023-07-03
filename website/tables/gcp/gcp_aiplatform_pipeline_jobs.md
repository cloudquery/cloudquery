# Table: gcp_aiplatform_pipeline_jobs

This table shows data for GCP AI Platform Pipeline Jobs.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.pipelineJobs#PipelineJob

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_pipeline_locations](gcp_aiplatform_pipeline_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|start_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|pipeline_spec|`json`|
|state|`utf8`|
|job_detail|`json`|
|error|`json`|
|labels|`json`|
|runtime_config|`json`|
|encryption_spec|`json`|
|service_account|`utf8`|
|network|`utf8`|
|template_uri|`utf8`|
|template_metadata|`json`|