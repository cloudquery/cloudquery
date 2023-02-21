# Table: oracle_virtualnetwork_subnets

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
|cidr_block|String|
|lifecycle_state|String|
|route_table_id|String|
|vcn_id|String|
|virtual_router_ip|String|
|virtual_router_mac|String|
|availability_domain|String|
|defined_tags|JSON|
|dhcp_options_id|String|
|display_name|String|
|dns_label|String|
|freeform_tags|JSON|
|ipv6_cidr_block|String|
|ipv6_cidr_blocks|StringArray|
|ipv6_virtual_router_ip|String|
|prohibit_internet_ingress|Bool|
|prohibit_public_ip_on_vnic|Bool|
|security_list_ids|StringArray|
|subnet_domain_name|String|
|time_created|Timestamp|