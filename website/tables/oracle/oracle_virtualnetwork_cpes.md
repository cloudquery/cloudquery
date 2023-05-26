# Table: oracle_virtualnetwork_cpes

This table shows data for Oracle Virtual Network Customer Premises Equipment (CPEs).

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
|id (PK)|`utf8`|
|ip_address|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|cpe_device_shape_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|