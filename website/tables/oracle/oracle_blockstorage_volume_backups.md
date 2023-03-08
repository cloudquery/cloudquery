# Table: oracle_blockstorage_volume_backups

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
|defined_tags|JSON|
|system_tags|JSON|
|expiration_time|Timestamp|
|freeform_tags|JSON|
|kms_key_id|String|
|size_in_g_bs|Int|
|size_in_m_bs|Int|
|source_type|String|
|source_volume_backup_id|String|
|time_request_received|Timestamp|
|unique_size_in_g_bs|Int|
|unique_size_in_mbs|Int|
|volume_id|String|