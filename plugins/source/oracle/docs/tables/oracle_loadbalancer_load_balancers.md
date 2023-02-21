# Table: oracle_loadbalancer_load_balancers

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
|lifecycle_state|String|
|time_created|Timestamp|
|shape_name|String|
|ip_addresses|JSON|
|shape_details|JSON|
|is_private|Bool|
|subnet_ids|StringArray|
|network_security_group_ids|StringArray|
|listeners|JSON|
|hostnames|JSON|
|ssl_cipher_suites|JSON|
|certificates|JSON|
|backend_sets|JSON|
|path_route_sets|JSON|
|freeform_tags|JSON|
|defined_tags|JSON|
|system_tags|JSON|
|rule_sets|JSON|
|routing_policies|JSON|