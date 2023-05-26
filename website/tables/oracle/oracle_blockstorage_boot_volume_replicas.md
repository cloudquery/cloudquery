# Table: oracle_blockstorage_boot_volume_replicas

This table shows data for Oracle Block Storage Boot Volume Replicas.

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
|availability_domain|`utf8`|
|display_name|`utf8`|
|lifecycle_state|`utf8`|
|size_in_g_bs|`int64`|
|time_created|`timestamp[us, tz=UTC]`|
|time_last_synced|`timestamp[us, tz=UTC]`|
|boot_volume_id|`utf8`|
|defined_tags|`json`|
|freeform_tags|`json`|
|image_id|`utf8`|
|total_data_transferred_in_g_bs|`int64`|
|volume_group_replica_id|`utf8`|