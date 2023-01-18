# Table: gcp_aiplatform_tensorboards

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.tensorboards#Tensorboard

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_tensorboard_locations](gcp_aiplatform_tensorboard_locations.md).

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
|encryption_spec|JSON|
|blob_storage_path_prefix|String|
|run_count|Int|
|create_time|Timestamp|
|update_time|Timestamp|
|labels|JSON|
|etag|String|