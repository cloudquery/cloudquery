# Table: oracle_blockstorage_volume_groups

This table shows data for Oracle Block Storage Volume Groups.

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
|size_in_m_bs|`int64`|
|time_created|`timestamp[us, tz=UTC]`|
|volume_ids|`list<item: utf8, nullable>`|
|defined_tags|`json`|
|freeform_tags|`json`|
|size_in_g_bs|`int64`|
|is_hydrated|`bool`|
|volume_group_replicas|`json`|