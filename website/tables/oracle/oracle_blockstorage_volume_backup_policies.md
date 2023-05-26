# Table: oracle_blockstorage_volume_backup_policies

This table shows data for Oracle Block Storage Volume Backup Policies.

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
|id (PK)|`utf8`|
|display_name|`utf8`|
|schedules|`json`|
|time_created|`timestamp[us, tz=UTC]`|
|destination_region|`utf8`|
|defined_tags|`json`|
|freeform_tags|`json`|