# Table: gcp_aiplatform_model_deployment_monitoring_jobs

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.modelDeploymentMonitoringJobs#ModelDeploymentMonitoringJob

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
|endpoint|String|
|state|String|
|schedule_state|String|
|latest_monitoring_pipeline_metadata|JSON|
|model_deployment_monitoring_objective_configs|JSON|
|model_deployment_monitoring_schedule_config|JSON|
|logging_sampling_strategy|JSON|
|model_monitoring_alert_config|JSON|
|predict_instance_schema_uri|String|
|sample_predict_instance|JSON|
|analysis_instance_schema_uri|String|
|bigquery_tables|JSON|
|log_ttl|Int|
|labels|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|next_schedule_time|Timestamp|
|stats_anomalies_base_directory|JSON|
|encryption_spec|JSON|
|enable_monitoring_pipeline_logs|Bool|
|error|JSON|