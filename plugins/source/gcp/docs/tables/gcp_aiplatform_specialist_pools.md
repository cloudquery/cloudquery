# Table: gcp_aiplatform_specialist_pools

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.specialistPools#SpecialistPool

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_specialistpool_locations](gcp_aiplatform_specialistpool_locations.md).

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
|specialist_managers_count|Int|
|specialist_manager_emails|StringArray|
|pending_data_labeling_jobs|StringArray|
|specialist_worker_emails|StringArray|