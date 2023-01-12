# Table: oracle_compute_instance_console_connections

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|compartment_id (PK)|String|
|id (PK)|String|
|connection_string|String|
|defined_tags|JSON|
|fingerprint|String|
|freeform_tags|JSON|
|instance_id|String|
|lifecycle_state|String|
|service_host_key_fingerprint|String|
|vnc_connection_string|String|