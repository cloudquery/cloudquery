# Table: oracle_virtual_network_remote_peering_connections

This table shows data for Oracle Virtual Network Remote Peering Connections.

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
|display_name|`utf8`|
|drg_id|`utf8`|
|id (PK)|`utf8`|
|is_cross_tenancy_peering|`bool`|
|lifecycle_state|`utf8`|
|peering_status|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|defined_tags|`json`|
|freeform_tags|`json`|
|peer_id|`utf8`|
|peer_region_name|`utf8`|
|peer_tenancy_id|`utf8`|