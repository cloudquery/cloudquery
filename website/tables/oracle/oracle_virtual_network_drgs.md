# Table: oracle_virtual_network_drgs

This table shows data for Oracle Virtual Network Dynamic Routing Gateways (DRGs).

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
|lifecycle_state|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|time_created|`timestamp[us, tz=UTC]`|
|default_drg_route_tables|`json`|
|default_export_drg_route_distribution_id|`utf8`|