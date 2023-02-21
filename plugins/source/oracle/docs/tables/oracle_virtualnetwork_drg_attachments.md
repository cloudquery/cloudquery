# Table: oracle_virtualnetwork_drg_attachments

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
|drg_id|String|
|lifecycle_state|String|
|display_name|String|
|time_created|Timestamp|
|drg_route_table_id|String|
|defined_tags|JSON|
|freeform_tags|JSON|
|route_table_id|String|
|vcn_id|String|
|export_drg_route_distribution_id|String|
|is_cross_tenancy|Bool|