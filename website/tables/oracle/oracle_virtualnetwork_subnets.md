# Table: oracle_virtualnetwork_subnets

This table shows data for Oracle Virtual Network Subnets.

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
|cidr_block|`utf8`|
|lifecycle_state|`utf8`|
|route_table_id|`utf8`|
|vcn_id|`utf8`|
|virtual_router_ip|`utf8`|
|virtual_router_mac|`utf8`|
|availability_domain|`utf8`|
|defined_tags|`json`|
|dhcp_options_id|`utf8`|
|display_name|`utf8`|
|dns_label|`utf8`|
|freeform_tags|`json`|
|ipv6_cidr_block|`utf8`|
|ipv6_cidr_blocks|`list<item: utf8, nullable>`|
|ipv6_virtual_router_ip|`utf8`|
|prohibit_internet_ingress|`bool`|
|prohibit_public_ip_on_vnic|`bool`|
|security_list_ids|`list<item: utf8, nullable>`|
|subnet_domain_name|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|