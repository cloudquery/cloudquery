# Table: oracle_compute_images

This table shows data for Oracle Compute Images.

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
|create_image_allowed|`bool`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|operating_system|`utf8`|
|operating_system_version|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|base_image_id|`utf8`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|launch_mode|`utf8`|
|launch_options|`json`|
|agent_features|`json`|
|listing_type|`utf8`|
|size_in_m_bs|`int64`|
|billable_size_in_g_bs|`int64`|