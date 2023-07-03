# Table: gcp_aiplatform_batch_prediction_jobs

This table shows data for GCP AI Platform Batch Prediction Jobs.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.batchPredictionJobs#BatchPredictionJob

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_job_locations](gcp_aiplatform_job_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|model|`utf8`|
|model_version_id|`utf8`|
|unmanaged_container_model|`json`|
|input_config|`json`|
|instance_config|`json`|
|model_parameters|`json`|
|output_config|`json`|
|dedicated_resources|`json`|
|service_account|`utf8`|
|manual_batch_tuning_parameters|`json`|
|generate_explanation|`bool`|
|explanation_spec|`json`|
|output_info|`json`|
|state|`utf8`|
|error|`json`|
|partial_failures|`json`|
|resources_consumed|`json`|
|completion_stats|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|start_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|encryption_spec|`json`|
|disable_container_logging|`bool`|