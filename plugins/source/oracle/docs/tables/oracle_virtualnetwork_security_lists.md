# Table: oracle_virtualnetwork_security_lists

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
|egress_security_rules|JSON|
|ingress_security_rules|JSON|
|lifecycle_state|String|
|time_created|Timestamp|
|vcn_id|String|
|defined_tags|JSON|
|freeform_tags|JSON|