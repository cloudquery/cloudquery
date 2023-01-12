# Table: oracle_identity_cost_tracking_tags

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|compartment_id|String|
|tag_namespace_id|String|
|tag_namespace_name|String|
|name|String|
|description|String|
|is_retired|Bool|
|time_created|Timestamp|
|freeform_tags|JSON|
|defined_tags|JSON|
|lifecycle_state|String|
|is_cost_tracking|Bool|