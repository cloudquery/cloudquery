# Table: gcp_aiplatform_training_pipelines

This table shows data for GCP AI Platform Training Pipelines.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.trainingPipelines#TrainingPipeline

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
|input_data_config|`json`|
|training_task_definition|`utf8`|
|training_task_inputs|`json`|
|training_task_metadata|`json`|
|model_to_upload|`json`|
|model_id|`utf8`|
|parent_model|`utf8`|
|state|`utf8`|
|error|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|start_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|encryption_spec|`json`|