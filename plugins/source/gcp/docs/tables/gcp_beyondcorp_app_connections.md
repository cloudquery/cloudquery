# Table: gcp_beyondcorp_app_connections

https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#AppConnection

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
|uid|String|
|type|String|
|application_endpoint|JSON|
|connectors|StringArray|
|state|String|
|gateway|JSON|