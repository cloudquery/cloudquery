# Table: gcp_aiplatform_model_deployment_monitoring_jobs

This table shows data for GCP AI Platform Model Deployment Monitoring Jobs.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.modelDeploymentMonitoringJobs#ModelDeploymentMonitoringJob

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
|endpoint|`utf8`|
|state|`utf8`|
|schedule_state|`utf8`|
|latest_monitoring_pipeline_metadata|`json`|
|model_deployment_monitoring_objective_configs|`json`|
|model_deployment_monitoring_schedule_config|`json`|
|logging_sampling_strategy|`json`|
|model_monitoring_alert_config|`json`|
|predict_instance_schema_uri|`utf8`|
|sample_predict_instance|`json`|
|analysis_instance_schema_uri|`utf8`|
|bigquery_tables|`json`|
|log_ttl|`int64`|
|labels|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|next_schedule_time|`timestamp[us, tz=UTC]`|
|stats_anomalies_base_directory|`json`|
|encryption_spec|`json`|
|enable_monitoring_pipeline_logs|`bool`|
|error|`json`|