# Table: oracle_identity_policies

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
|statements|StringArray|
|description|String|
|time_created|Timestamp|
|lifecycle_state|String|
|inactive_status|Int|
|version_date|JSON|
|freeform_tags|JSON|
|defined_tags|JSON|