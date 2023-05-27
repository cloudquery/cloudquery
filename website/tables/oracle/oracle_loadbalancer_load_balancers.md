# Table: oracle_loadbalancer_load_balancers

This table shows data for Oracle Load Balancer Load Balancers.

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
|shape_name|`utf8`|
|ip_addresses|`json`|
|shape_details|`json`|
|is_private|`bool`|
|subnet_ids|`list<item: utf8, nullable>`|
|network_security_group_ids|`list<item: utf8, nullable>`|
|listeners|`json`|
|hostnames|`json`|
|ssl_cipher_suites|`json`|
|certificates|`json`|
|backend_sets|`json`|
|path_route_sets|`json`|
|freeform_tags|`json`|
|defined_tags|`json`|
|system_tags|`json`|
|rule_sets|`json`|
|routing_policies|`json`|