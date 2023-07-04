# Table: gcp_aiplatform_index_endpoints

This table shows data for GCP AI Platform Index Endpoints.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.indexEndpoints#IndexEndpoint

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_indexendpoint_locations](gcp_aiplatform_indexendpoint_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|description|`utf8`|
|deployed_indexes|`json`|
|etag|`utf8`|
|labels|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|network|`utf8`|
|enable_private_service_connect|`bool`|
|private_service_connect_config|`json`|
|public_endpoint_enabled|`bool`|
|public_endpoint_domain_name|`utf8`|