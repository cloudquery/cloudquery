# Table: gcp_aiplatform_data_labeling_jobs

This table shows data for GCP AI Platform Data Labeling Jobs.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.dataLabelingJobs#DataLabelingJob

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
|datasets|`list<item: utf8, nullable>`|
|annotation_labels|`json`|
|labeler_count|`int64`|
|instruction_uri|`utf8`|
|inputs_schema_uri|`utf8`|
|inputs|`json`|
|state|`utf8`|
|labeling_progress|`int64`|
|current_spend|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|error|`json`|
|labels|`json`|
|specialist_pools|`list<item: utf8, nullable>`|
|encryption_spec|`json`|
|active_learning_config|`json`|