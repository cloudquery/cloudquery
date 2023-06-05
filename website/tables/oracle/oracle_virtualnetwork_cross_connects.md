# Table: oracle_virtualnetwork_cross_connects

This table shows data for Oracle Virtual Network Cross Connects.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|cross_connect_group_id|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|location_name|`utf8`|
|port_name|`utf8`|
|port_speed_shape_name|`utf8`|
|customer_reference_name|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|macsec_properties|`json`|
|oci_physical_device_name|`utf8`|
|oci_logical_device_name|`utf8`|