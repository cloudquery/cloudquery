# Table: oracle_virtualnetwork_vtaps

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
|vcn_id|String|
|lifecycle_state|String|
|source_id|String|
|capture_filter_id|String|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|lifecycle_state_details|String|
|time_created|Timestamp|
|target_id|String|
|target_ip|String|
|encapsulation_protocol|String|
|vxlan_network_identifier|Int|
|is_vtap_enabled|Bool|
|source_type|String|
|traffic_mode|String|
|max_packet_size|Int|
|target_type|String|
|source_private_endpoint_ip|String|
|source_private_endpoint_subnet_id|String|