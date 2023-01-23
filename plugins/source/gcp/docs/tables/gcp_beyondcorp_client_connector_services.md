# Table: gcp_beyondcorp_client_connector_services

https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.clientConnectorServices#ClientConnectorService

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
|display_name|String|
|ingress|JSON|
|egress|JSON|
|state|String|