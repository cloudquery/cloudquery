# Table: oracle_virtualnetwork_nat_gateways

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
|nat_ip|String|
|time_created|Timestamp|
|vcn_id|String|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|public_ip_id|String|
|route_table_id|String|