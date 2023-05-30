# Table: oracle_virtualnetwork_local_peering_gateways

This table shows data for Oracle Virtual Network Local Peering Gateways.

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
|display_name|`utf8`|
|is_cross_tenancy_peering|`bool`|
|lifecycle_state|`utf8`|
|peering_status|`utf8`|
|peer_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|vcn_id|`utf8`|
|defined_tags|`json`|
|freeform_tags|`json`|
|peer_advertised_cidr|`utf8`|
|peer_advertised_cidr_details|`list<item: utf8, nullable>`|
|peering_status_details|`utf8`|
|route_table_id|`utf8`|