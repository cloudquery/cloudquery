# Table: gcp_aiplatform_datasets

This table shows data for GCP AI Platform Datasets.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.datasets#Dataset

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_dataset_locations](gcp_aiplatform_dataset_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|description|`utf8`|
|metadata_schema_uri|`utf8`|
|metadata|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|etag|`utf8`|
|labels|`json`|
|saved_queries|`json`|
|encryption_spec|`json`|
|metadata_artifact|`utf8`|