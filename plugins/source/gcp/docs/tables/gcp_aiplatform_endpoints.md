# Table: gcp_aiplatform_endpoints

https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.endpoints#Endpoint

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_aiplatform_endpoint_locations](gcp_aiplatform_endpoint_locations.md).

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
|deployed_models|JSON|
|traffic_split|JSON|
|etag|String|
|labels|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|encryption_spec|JSON|
|network|String|
|enable_private_service_connect|Bool|
|model_deployment_monitoring_job|String|
|predict_request_response_logging_config|JSON|