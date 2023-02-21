# Table: gcp_livestream_inputs

https://cloud.google.com/livestream/docs/reference/rest/v1/projects.locations.inputs

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
|type|String|
|tier|String|
|uri|String|
|preprocessing_config|JSON|
|security_rules|JSON|
|input_stream_property|JSON|