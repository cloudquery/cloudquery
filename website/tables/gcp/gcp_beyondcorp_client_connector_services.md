# Table: gcp_beyondcorp_client_connector_services

This table shows data for GCP Beyondcorp Client Connector Services.

https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.clientConnectorServices#ClientConnectorService

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
|display_name|`utf8`|
|ingress|`json`|
|egress|`json`|
|state|`utf8`|