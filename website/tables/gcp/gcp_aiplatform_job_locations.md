# Table: gcp_aiplatform_job_locations

This table shows data for GCP AI Platform Job Locations.

https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_aiplatform_job_locations:
  - [gcp_aiplatform_batch_prediction_jobs](gcp_aiplatform_batch_prediction_jobs)
  - [gcp_aiplatform_custom_jobs](gcp_aiplatform_custom_jobs)
  - [gcp_aiplatform_data_labeling_jobs](gcp_aiplatform_data_labeling_jobs)
  - [gcp_aiplatform_hyperparameter_tuning_jobs](gcp_aiplatform_hyperparameter_tuning_jobs)
  - [gcp_aiplatform_model_deployment_monitoring_jobs](gcp_aiplatform_model_deployment_monitoring_jobs)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|location_id|`utf8`|
|display_name|`utf8`|
|labels|`json`|
|metadata|`json`|