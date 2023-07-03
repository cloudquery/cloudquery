# Table: gcp_aiplatform_featurestores

This table shows data for GCP AI Platform featurestores.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores#Featurestore

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_featurestore_locations](gcp_aiplatform_featurestore_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|etag|`utf8`|
|labels|`json`|
|online_serving_config|`json`|
|state|`utf8`|
|encryption_spec|`json`|