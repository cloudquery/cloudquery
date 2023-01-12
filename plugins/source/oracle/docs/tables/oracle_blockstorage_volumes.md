# Table: oracle_blockstorage_volumes

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
|availability_domain|String|
|display_name|String|
|lifecycle_state|String|
|size_in_m_bs|Int|
|time_created|Timestamp|
|defined_tags|JSON|
|freeform_tags|JSON|
|system_tags|JSON|
|is_hydrated|Bool|
|kms_key_id|String|
|vpus_per_gb|Int|
|size_in_g_bs|Int|
|volume_group_id|String|
|is_auto_tune_enabled|Bool|
|auto_tuned_vpus_per_gb|Int|
|block_volume_replicas|JSON|
|autotune_policies|JSON|