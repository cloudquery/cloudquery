# Table: gcp_aiplatform_hyperparameter_tuning_jobs

This table shows data for GCP AI Platform Hyperparameter Tuning Jobs.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.hyperparameterTuningJobs#HyperparameterTuningJob

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
|study_spec|`json`|
|max_trial_count|`int64`|
|parallel_trial_count|`int64`|
|max_failed_trial_count|`int64`|
|trial_job_spec|`json`|
|trials|`json`|
|state|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|start_time|`timestamp[us, tz=UTC]`|
|end_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|error|`json`|
|labels|`json`|
|encryption_spec|`json`|