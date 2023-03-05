# Table: oracle_virtualnetwork_vcns

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
|cidr_blocks|StringArray|
|lifecycle_state|String|
|byoipv6_cidr_blocks|StringArray|
|ipv6_private_cidr_blocks|StringArray|
|default_dhcp_options_id|String|
|default_route_table_id|String|
|default_security_list_id|String|
|defined_tags|JSON|
|display_name|String|
|dns_label|String|
|freeform_tags|JSON|
|ipv6_cidr_blocks|StringArray|
|time_created|Timestamp|
|vcn_domain_name|String|