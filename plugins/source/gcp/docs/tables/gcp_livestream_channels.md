# Table: gcp_livestream_channels

https://cloud.google.com/livestream/docs/reference/rest/v1/projects.locations.channels

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
|input_attachments|JSON|
|active_input|String|
|output|JSON|
|elementary_streams|JSON|
|mux_streams|JSON|
|manifests|JSON|
|sprite_sheets|JSON|
|streaming_state|String|
|streaming_error|JSON|
|log_config|JSON|