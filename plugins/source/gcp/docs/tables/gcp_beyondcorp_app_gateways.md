# Table: gcp_beyondcorp_app_gateways

This table shows data for GCP Beyondcorp App Gateways.

https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appGateways#AppGateway

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
|uid|`utf8`|
|type|`utf8`|
|state|`utf8`|
|uri|`utf8`|
|allocated_connections|`json`|
|host_type|`utf8`|