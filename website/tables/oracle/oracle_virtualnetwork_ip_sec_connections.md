# Table: oracle_virtualnetwork_ip_sec_connections

This table shows data for Oracle Virtual Network IP Sec Connections.

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
|cpe_id|`utf8`|
|drg_id|`utf8`|
|lifecycle_state|`utf8`|
|static_routes|`list<item: utf8, nullable>`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|cpe_local_identifier|`utf8`|
|cpe_local_identifier_type|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|