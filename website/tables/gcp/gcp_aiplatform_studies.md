# Table: gcp_aiplatform_studies

This table shows data for GCP AI Platform Studies.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.studies#Study

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_vizier_locations](gcp_aiplatform_vizier_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|study_spec|`json`|
|state|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|inactive_reason|`utf8`|