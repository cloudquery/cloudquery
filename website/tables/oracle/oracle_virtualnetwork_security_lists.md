# Table: oracle_virtualnetwork_security_lists

This table shows data for Oracle Virtual Network Security Lists.

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
|display_name|`utf8`|
|egress_security_rules|`json`|
|id (PK)|`utf8`|
|ingress_security_rules|`json`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|vcn_id|`utf8`|
|defined_tags|`json`|
|freeform_tags|`json`|