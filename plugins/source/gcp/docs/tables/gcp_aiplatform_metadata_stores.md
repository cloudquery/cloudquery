# Table: gcp_aiplatform_metadata_stores

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.metadataStores#MetadataStore

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_metadata_locations](gcp_aiplatform_metadata_locations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|create_time|Timestamp|
|update_time|Timestamp|
|encryption_spec|JSON|
|description|String|
|state|JSON|