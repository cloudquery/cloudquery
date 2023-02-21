# Table: oracle_blockstorage_volume_group_backups

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
|type|String|
|volume_backup_ids|StringArray|
|defined_tags|JSON|
|expiration_time|Timestamp|
|freeform_tags|JSON|
|size_in_m_bs|Int|
|size_in_g_bs|Int|
|source_type|String|
|time_request_received|Timestamp|
|unique_size_in_mbs|Int|
|unique_size_in_gbs|Int|
|volume_group_id|String|
|source_volume_group_backup_id|String|