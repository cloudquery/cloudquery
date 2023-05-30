# Table: oracle_virtualnetwork_vtaps

This table shows data for Oracle Virtual Network Virtual Tunnel Access Points (VTAPs).

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
|vcn_id|`utf8`|
|lifecycle_state|`utf8`|
|source_id|`utf8`|
|capture_filter_id|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|lifecycle_state_details|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|target_id|`utf8`|
|target_ip|`utf8`|
|encapsulation_protocol|`utf8`|
|vxlan_network_identifier|`int64`|
|is_vtap_enabled|`bool`|
|source_type|`utf8`|
|traffic_mode|`utf8`|
|max_packet_size|`int64`|
|target_type|`utf8`|
|source_private_endpoint_ip|`utf8`|
|source_private_endpoint_subnet_id|`utf8`|