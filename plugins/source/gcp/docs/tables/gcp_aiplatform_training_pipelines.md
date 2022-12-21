# Table: gcp_aiplatform_training_pipelines

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.trainingPipelines#TrainingPipeline

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
|input_data_config|JSON|
|training_task_definition|String|
|training_task_inputs|JSON|
|training_task_metadata|JSON|
|model_to_upload|JSON|
|model_id|String|
|parent_model|String|
|state|String|
|error|JSON|
|create_time|Timestamp|
|start_time|Timestamp|
|end_time|Timestamp|
|update_time|Timestamp|
|labels|JSON|
|encryption_spec|JSON|