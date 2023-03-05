# Table: oracle_virtualnetwork_byoip_ranges

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
|byoip_range_vcn_ipv6_allocations|JSON|
|cidr_block|String|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|ipv6_cidr_block|String|
|lifecycle_state|String|
|lifecycle_details|String|
|time_created|Timestamp|