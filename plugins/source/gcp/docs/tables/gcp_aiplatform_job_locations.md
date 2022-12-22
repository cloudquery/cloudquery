# Table: gcp_aiplatform_job_locations

https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_aiplatform_job_locations:
  - [gcp_aiplatform_batch_prediction_jobs](gcp_aiplatform_batch_prediction_jobs.md)
  - [gcp_aiplatform_custom_jobs](gcp_aiplatform_custom_jobs.md)
  - [gcp_aiplatform_data_labeling_jobs](gcp_aiplatform_data_labeling_jobs.md)
  - [gcp_aiplatform_hyperparameter_tuning_jobs](gcp_aiplatform_hyperparameter_tuning_jobs.md)
  - [gcp_aiplatform_model_deployment_monitoring_jobs](gcp_aiplatform_model_deployment_monitoring_jobs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|location_id|String|
|display_name|String|
|labels|JSON|
|metadata|JSON|