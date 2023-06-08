# Table: oracle_virtual_network_byoip_ranges

This table shows data for Oracle Virtual Network Bring Your Own IP (BYOIP) Ranges.

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
|byoip_range_vcn_ipv6_allocations|`json`|
|cidr_block|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|id (PK)|`utf8`|
|ipv6_cidr_block|`utf8`|
|lifecycle_state|`utf8`|
|lifecycle_details|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|