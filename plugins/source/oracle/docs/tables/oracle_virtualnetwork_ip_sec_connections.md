# Table: oracle_virtualnetwork_ip_sec_connections

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
|cpe_id|String|
|drg_id|String|
|lifecycle_state|String|
|static_routes|StringArray|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|cpe_local_identifier|String|
|cpe_local_identifier_type|String|
|time_created|Timestamp|