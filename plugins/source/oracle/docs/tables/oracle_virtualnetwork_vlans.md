# Table: oracle_virtualnetwork_vlans

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
|cidr_block|String|
|lifecycle_state|String|
|vcn_id|String|
|availability_domain|String|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|nsg_ids|StringArray|
|vlan_tag|Int|
|route_table_id|String|
|time_created|Timestamp|