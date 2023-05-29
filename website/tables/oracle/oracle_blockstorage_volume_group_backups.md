# Table: oracle_blockstorage_volume_group_backups

This table shows data for Oracle Block Storage Volume Group Backups.

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
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|volume_backup_ids|`list<item: utf8, nullable>`|
|defined_tags|`json`|
|expiration_time|`timestamp[us, tz=UTC]`|
|freeform_tags|`json`|
|size_in_m_bs|`int64`|
|size_in_g_bs|`int64`|
|source_type|`utf8`|
|time_request_received|`timestamp[us, tz=UTC]`|
|unique_size_in_mbs|`int64`|
|unique_size_in_gbs|`int64`|
|volume_group_id|`utf8`|
|source_volume_group_backup_id|`utf8`|