# Table: oracle_virtualnetwork_remote_peering_connections

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
|drg_id|String|
|is_cross_tenancy_peering|Bool|
|lifecycle_state|String|
|peering_status|String|
|time_created|Timestamp|
|defined_tags|JSON|
|freeform_tags|JSON|
|peer_id|String|
|peer_region_name|String|
|peer_tenancy_id|String|