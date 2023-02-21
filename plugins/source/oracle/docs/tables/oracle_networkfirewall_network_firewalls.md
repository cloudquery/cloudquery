# Table: oracle_networkfirewall_network_firewalls

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
|subnet_id|String|
|network_firewall_policy_id|String|
|time_created|Timestamp|
|lifecycle_state|String|
|freeform_tags|JSON|
|defined_tags|JSON|
|availability_domain|String|
|ipv4_address|String|
|ipv6_address|String|
|time_updated|Timestamp|
|lifecycle_details|String|
|system_tags|JSON|