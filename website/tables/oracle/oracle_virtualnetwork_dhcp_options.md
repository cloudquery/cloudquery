# Table: oracle_virtualnetwork_dhcp_options

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
|lifecycle_state|String|
|options|JSON|
|time_created|Timestamp|
|vcn_id|String|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|domain_name_type|String|