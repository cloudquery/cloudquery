# Table: oracle_blockstorage_boot_volume_backups

This table shows data for Oracle Block Storage Boot Volume Backups.

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
|display_name|`utf8`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|boot_volume_id|`utf8`|
|defined_tags|`json`|
|system_tags|`json`|
|expiration_time|`timestamp[us, tz=UTC]`|
|freeform_tags|`json`|
|image_id|`utf8`|
|kms_key_id|`utf8`|
|size_in_g_bs|`int64`|
|source_boot_volume_backup_id|`utf8`|
|source_type|`utf8`|
|time_request_received|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|unique_size_in_g_bs|`int64`|