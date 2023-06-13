# Table: oracle_virtualnetwork_nat_gateways

This table shows data for Oracle Virtual Network Nat Gateways.

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
|block_traffic|`bool`|
|lifecycle_state|`utf8`|
|nat_ip|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|vcn_id|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|public_ip_id|`utf8`|
|route_table_id|`utf8`|