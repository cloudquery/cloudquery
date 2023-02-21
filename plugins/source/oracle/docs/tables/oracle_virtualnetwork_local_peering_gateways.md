# Table: oracle_virtualnetwork_local_peering_gateways

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
|display_name|String|
|is_cross_tenancy_peering|Bool|
|lifecycle_state|String|
|peering_status|String|
|peer_id|String|
|time_created|Timestamp|
|vcn_id|String|
|defined_tags|JSON|
|freeform_tags|JSON|
|peer_advertised_cidr|String|
|peer_advertised_cidr_details|StringArray|
|peering_status_details|String|
|route_table_id|String|