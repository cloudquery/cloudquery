# Table: gcp_aiplatform_studies

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.studies#Study

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_vizier_locations](gcp_aiplatform_vizier_locations.md).

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
|study_spec|JSON|
|state|String|
|create_time|Timestamp|
|inactive_reason|String|