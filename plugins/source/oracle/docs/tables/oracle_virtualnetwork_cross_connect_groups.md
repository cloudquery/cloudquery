# Table: oracle_virtualnetwork_cross_connect_groups

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
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|lifecycle_state|String|
|customer_reference_name|String|
|time_created|Timestamp|
|macsec_properties|JSON|
|oci_physical_device_name|String|
|oci_logical_device_name|String|