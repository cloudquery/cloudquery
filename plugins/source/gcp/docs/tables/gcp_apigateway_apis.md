# Table: gcp_apigateway_apis

This table shows data for GCP API Gateway APIs.

https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations.apis#Api

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|display_name|`utf8`|
|managed_service|`utf8`|
|state|`utf8`|