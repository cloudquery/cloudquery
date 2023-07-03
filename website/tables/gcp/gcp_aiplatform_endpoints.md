# Table: gcp_aiplatform_endpoints

This table shows data for GCP AI Platform Endpoints.

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.endpoints#Endpoint

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_endpoint_locations](gcp_aiplatform_endpoint_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|display_name|`utf8`|
|description|`utf8`|
|deployed_models|`json`|
|traffic_split|`json`|
|etag|`utf8`|
|labels|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|encryption_spec|`json`|
|network|`utf8`|
|enable_private_service_connect|`bool`|
|model_deployment_monitoring_job|`utf8`|
|predict_request_response_logging_config|`json`|