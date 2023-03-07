# Table: gcp_apigateway_gateways

https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations.gateways#Gateway

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|create_time|Timestamp|
|update_time|Timestamp|
|labels|JSON|
|display_name|String|
|api_config|String|
|state|String|
|default_hostname|String|