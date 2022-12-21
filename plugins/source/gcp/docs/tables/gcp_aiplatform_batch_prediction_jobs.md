# Table: gcp_aiplatform_batch_prediction_jobs

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.batchPredictionJobs#BatchPredictionJob

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
|model|String|
|model_version_id|String|
|unmanaged_container_model|JSON|
|input_config|JSON|
|model_parameters|JSON|
|output_config|JSON|
|dedicated_resources|JSON|
|service_account|String|
|manual_batch_tuning_parameters|JSON|
|generate_explanation|Bool|
|explanation_spec|JSON|
|output_info|JSON|
|state|String|
|error|JSON|
|partial_failures|JSON|
|resources_consumed|JSON|
|completion_stats|JSON|
|create_time|Timestamp|
|start_time|Timestamp|
|end_time|Timestamp|
|update_time|Timestamp|
|labels|JSON|
|encryption_spec|JSON|