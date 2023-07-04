# Table: gcp_aiplatform_tensorboards

This table shows data for GCP AI Platform Tensorboards.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.tensorboards#Tensorboard

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_tensorboard_locations](gcp_aiplatform_tensorboard_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|description|`utf8`|
|encryption_spec|`json`|
|blob_storage_path_prefix|`utf8`|
|run_count|`int64`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|etag|`utf8`|