# Table: oracle_compute_instance_console_connections

This table shows data for Oracle Compute Instance Console Connections.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|id (PK)|`utf8`|
|connection_string|`utf8`|
|defined_tags|`json`|
|fingerprint|`utf8`|
|freeform_tags|`json`|
|instance_id|`utf8`|
|lifecycle_state|`utf8`|
|service_host_key_fingerprint|`utf8`|
|vnc_connection_string|`utf8`|