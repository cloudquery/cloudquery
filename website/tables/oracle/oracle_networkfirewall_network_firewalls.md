# Table: oracle_networkfirewall_network_firewalls

This table shows data for Oracle Network Firewall Network Firewalls.

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
|display_name|`utf8`|
|subnet_id|`utf8`|
|network_firewall_policy_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_state|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|
|availability_domain|`utf8`|
|ipv4_address|`utf8`|
|ipv6_address|`utf8`|
|time_updated|`timestamp[us, tz=UTC]`|
|lifecycle_details|`utf8`|
|system_tags|`json`|