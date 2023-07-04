# Table: gcp_livestream_channels

This table shows data for GCP Livestream Channels.

https://cloud.google.com/livestream/docs/reference/rest/v1/projects.locations.channels

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
|input_attachments|`json`|
|active_input|`utf8`|
|output|`json`|
|elementary_streams|`json`|
|mux_streams|`json`|
|manifests|`json`|
|sprite_sheets|`json`|
|streaming_state|`utf8`|
|streaming_error|`json`|
|log_config|`json`|
|timecode_config|`json`|
|encryptions|`json`|
|input_config|`json`|