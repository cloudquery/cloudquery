# Table: gcp_aiplatform_models

This table shows data for GCP AI Platform Models.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.models#Model

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_model_locations](gcp_aiplatform_model_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|version_id|`utf8`|
|version_aliases|`list<item: utf8, nullable>`|
|version_create_time|`timestamp[us, tz=UTC]`|
|version_update_time|`timestamp[us, tz=UTC]`|
|display_name|`utf8`|
|description|`utf8`|
|version_description|`utf8`|
|predict_schemata|`json`|
|metadata_schema_uri|`utf8`|
|supported_export_formats|`json`|
|training_pipeline|`utf8`|
|container_spec|`json`|
|artifact_uri|`utf8`|
|supported_deployment_resources_types|`list<item: int64, nullable>`|
|supported_input_storage_formats|`list<item: utf8, nullable>`|
|supported_output_storage_formats|`list<item: utf8, nullable>`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|deployed_models|`json`|
|explanation_spec|`json`|
|etag|`utf8`|
|labels|`json`|
|encryption_spec|`json`|
|model_source_info|`json`|
|original_model_info|`json`|
|metadata_artifact|`utf8`|