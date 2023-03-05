# Table: oracle_blockstorage_boot_volume_backups

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
|boot_volume_id|String|
|defined_tags|JSON|
|system_tags|JSON|
|expiration_time|Timestamp|
|freeform_tags|JSON|
|image_id|String|
|kms_key_id|String|
|size_in_g_bs|Int|
|source_boot_volume_backup_id|String|
|source_type|String|
|time_request_received|Timestamp|
|type|String|
|unique_size_in_g_bs|Int|