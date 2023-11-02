# Table: gcp_livestream_inputs

This table shows data for GCP Livestream Inputs.

https://cloud.google.com/livestream/docs/reference/rest/v1/projects.locations.inputs

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
|type|`utf8`|
|tier|`utf8`|
|uri|`utf8`|
|preprocessing_config|`json`|
|security_rules|`json`|
|input_stream_property|`json`|