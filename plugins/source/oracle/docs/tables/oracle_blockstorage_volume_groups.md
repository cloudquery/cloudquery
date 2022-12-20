# Table: oracle_blockstorage_volume_groups

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
|volume_ids|StringArray|
|defined_tags|JSON|
|freeform_tags|JSON|
|size_in_g_bs|Int|
|is_hydrated|Bool|
|volume_group_replicas|JSON|