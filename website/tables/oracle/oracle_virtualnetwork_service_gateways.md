# Table: oracle_virtualnetwork_service_gateways

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
|block_traffic|Bool|
|lifecycle_state|String|
|services|JSON|
|vcn_id|String|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|route_table_id|String|
|time_created|Timestamp|