# Table: gcp_aiplatform_data_labeling_jobs

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.dataLabelingJobs#DataLabelingJob

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
|datasets|StringArray|
|annotation_labels|JSON|
|labeler_count|Int|
|instruction_uri|String|
|inputs_schema_uri|String|
|inputs|JSON|
|state|String|
|labeling_progress|Int|
|current_spend|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|error|JSON|
|labels|JSON|
|specialist_pools|StringArray|
|encryption_spec|JSON|
|active_learning_config|JSON|