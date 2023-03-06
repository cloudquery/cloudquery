# Table: oracle_virtualnetwork_capture_filters

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
|lifecycle_state|String|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|filter_type|String|
|time_created|Timestamp|
|vtap_capture_filter_rules|JSON|