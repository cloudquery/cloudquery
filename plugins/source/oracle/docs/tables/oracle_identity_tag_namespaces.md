# Table: oracle_identity_tag_namespaces

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
|name|String|
|description|String|
|freeform_tags|JSON|
|defined_tags|JSON|
|is_retired|Bool|
|lifecycle_state|String|
|time_created|Timestamp|
|locks|JSON|