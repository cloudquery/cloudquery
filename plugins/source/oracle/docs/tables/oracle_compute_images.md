# Table: oracle_compute_images

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
|create_image_allowed|Bool|
|lifecycle_state|String|
|operating_system|String|
|operating_system_version|String|
|time_created|Timestamp|
|base_image_id|String|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|launch_mode|String|
|launch_options|JSON|
|agent_features|JSON|
|listing_type|String|
|size_in_m_bs|Int|
|billable_size_in_g_bs|Int|