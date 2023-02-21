# Table: oracle_filestorage_file_systems

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
|metered_bytes|Int|
|display_name|String|
|lifecycle_state|String|
|time_created|Timestamp|
|availability_domain|String|
|freeform_tags|JSON|
|defined_tags|JSON|
|kms_key_id|String|
|source_details|JSON|
|is_clone_parent|Bool|
|is_hydrated|Bool|
|lifecycle_details|String|