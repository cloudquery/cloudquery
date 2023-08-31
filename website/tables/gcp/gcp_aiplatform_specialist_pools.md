# Table: gcp_aiplatform_specialist_pools

This table shows data for GCP AI Platform Specialist Pools.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.specialistPools#SpecialistPool

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_specialistpool_locations](gcp_aiplatform_specialistpool_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|specialist_managers_count|`int64`|
|specialist_manager_emails|`list<item: utf8, nullable>`|
|pending_data_labeling_jobs|`list<item: utf8, nullable>`|
|specialist_worker_emails|`list<item: utf8, nullable>`|