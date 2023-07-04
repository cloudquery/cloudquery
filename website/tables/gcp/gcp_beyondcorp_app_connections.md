# Table: gcp_beyondcorp_app_connections

This table shows data for GCP Beyondcorp App Connections.

https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#AppConnection

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
|application_endpoint|`json`|
|connectors|`list<item: utf8, nullable>`|
|state|`utf8`|
|gateway|`json`|