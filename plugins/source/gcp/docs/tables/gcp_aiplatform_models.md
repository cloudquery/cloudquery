# Table: gcp_aiplatform_models

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.models#Model

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_model_locations](gcp_aiplatform_model_locations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|version_id|String|
|version_aliases|StringArray|
|version_create_time|Timestamp|
|version_update_time|Timestamp|
|display_name|String|
|description|String|
|version_description|String|
|predict_schemata|JSON|
|metadata_schema_uri|String|
|supported_export_formats|JSON|
|training_pipeline|String|
|container_spec|JSON|
|artifact_uri|String|
|supported_deployment_resources_types|IntArray|
|supported_input_storage_formats|StringArray|
|supported_output_storage_formats|StringArray|
|create_time|Timestamp|
|update_time|Timestamp|
|deployed_models|JSON|
|explanation_spec|JSON|
|etag|String|
|labels|JSON|
|encryption_spec|JSON|
|model_source_info|JSON|
|metadata_artifact|String|