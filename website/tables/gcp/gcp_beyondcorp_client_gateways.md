# Table: gcp_beyondcorp_client_gateways

This table shows data for GCP Beyondcorp Client Gateways.

https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.clientGateways#ClientGateway

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
|state|`utf8`|
|id|`utf8`|
|client_connector_service|`utf8`|