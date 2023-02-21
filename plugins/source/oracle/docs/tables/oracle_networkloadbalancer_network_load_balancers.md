# Table: oracle_networkloadbalancer_network_load_balancers

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
|ip_addresses|JSON|
|subnet_id|String|
|lifecycle_details|String|
|nlb_ip_version|String|
|time_updated|Timestamp|
|is_private|Bool|
|is_preserve_source_destination|Bool|
|network_security_group_ids|StringArray|
|listeners|JSON|
|backend_sets|JSON|
|freeform_tags|JSON|
|defined_tags|JSON|
|system_tags|JSON|