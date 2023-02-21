# Table: oracle_blockstorage_boot_volume_replicas

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
|size_in_g_bs|Int|
|time_created|Timestamp|
|time_last_synced|Timestamp|
|boot_volume_id|String|
|defined_tags|JSON|
|freeform_tags|JSON|
|image_id|String|
|total_data_transferred_in_g_bs|Int|
|volume_group_replica_id|String|