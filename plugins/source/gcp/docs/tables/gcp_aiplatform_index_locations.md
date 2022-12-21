# Table: gcp_aiplatform_index_locations

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.indexes#Index

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_aiplatform_index_locations:
  - [gcp_aiplatform_indexes](gcp_aiplatform_indexes.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|location_id|String|
|display_name|String|
|labels|JSON|
|metadata|JSON|