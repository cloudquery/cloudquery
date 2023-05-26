# Table: oracle_filestorage_file_systems

This table shows data for Oracle File Storage File Systems.

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
|metered_bytes|`int64`|
|display_name|`utf8`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|availability_domain|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|
|kms_key_id|`utf8`|
|source_details|`json`|
|is_clone_parent|`bool`|
|is_hydrated|`bool`|
|lifecycle_details|`utf8`|