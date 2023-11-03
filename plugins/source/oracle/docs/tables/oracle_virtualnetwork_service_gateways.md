# Table: oracle_virtualnetwork_service_gateways

This table shows data for Oracle Virtual Network Service Gateways.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|block_traffic|`bool`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|services|`json`|
|vcn_id|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|route_table_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|