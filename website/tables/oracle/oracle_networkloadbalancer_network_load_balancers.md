# Table: oracle_networkloadbalancer_network_load_balancers

This table shows data for Oracle Networkloadbalancer Network Load Balancers.

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
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|ip_addresses|`json`|
|subnet_id|`utf8`|
|lifecycle_details|`utf8`|
|nlb_ip_version|`utf8`|
|time_updated|`timestamp[us, tz=UTC]`|
|is_private|`bool`|
|is_preserve_source_destination|`bool`|
|network_security_group_ids|`list<item: utf8, nullable>`|
|listeners|`json`|
|backend_sets|`json`|
|freeform_tags|`json`|
|defined_tags|`json`|
|system_tags|`json`|