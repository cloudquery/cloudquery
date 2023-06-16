# Table: oracle_virtualnetwork_vcns

This table shows data for Oracle Virtual Network Vcns.

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
|cidr_block|`utf8`|
|cidr_blocks|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|byoipv6_cidr_blocks|`list<item: utf8, nullable>`|
|ipv6_private_cidr_blocks|`list<item: utf8, nullable>`|
|default_dhcp_options_id|`utf8`|
|default_route_table_id|`utf8`|
|default_security_list_id|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|dns_label|`utf8`|
|freeform_tags|`json`|
|ipv6_cidr_blocks|`list<item: utf8, nullable>`|
|time_created|`timestamp[us, tz=UTC]`|
|vcn_domain_name|`utf8`|