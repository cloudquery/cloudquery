# Table: oracle_blockstorage_boot_volumes

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
|lifecycle_state|String|
|size_in_m_bs|Int|
|time_created|Timestamp|
|defined_tags|JSON|
|system_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|image_id|String|
|is_hydrated|Bool|
|vpus_per_gb|Int|
|size_in_g_bs|Int|
|volume_group_id|String|
|kms_key_id|String|
|is_auto_tune_enabled|Bool|
|auto_tuned_vpus_per_gb|Int|
|boot_volume_replicas|JSON|
|autotune_policies|JSON|