# Table: oracle_database_key_stores

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
|lifecycle_details|String|
|associated_databases|JSON|
|freeform_tags|JSON|
|defined_tags|JSON|