# Table: gcp_aiplatform_datasets

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.datasets#Dataset

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_dataset_locations](gcp_aiplatform_dataset_locations.md).

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
|description|String|
|metadata_schema_uri|String|
|metadata|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|etag|String|
|labels|JSON|
|encryption_spec|JSON|
|metadata_artifact|String|