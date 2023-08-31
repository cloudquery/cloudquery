# Table: gcp_aiplatform_metadata_stores

This table shows data for GCP AI Platform Metadata Stores.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.metadataStores#MetadataStore

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_metadata_locations](gcp_aiplatform_metadata_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|encryption_spec|`json`|
|description|`utf8`|
|state|`json`|