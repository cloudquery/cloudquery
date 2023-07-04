# Table: gcp_aiplatform_indexes

This table shows data for GCP AI Platform Indexes.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.indexes#Index

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_index_locations](gcp_aiplatform_index_locations).

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
|deployed_indexes|`json`|
|etag|`utf8`|
|labels|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|index_stats|`json`|
|index_update_method|`utf8`|