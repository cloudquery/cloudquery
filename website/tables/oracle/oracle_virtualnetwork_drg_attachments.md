# Table: oracle_virtualnetwork_drg_attachments

This table shows data for Oracle Virtual Network Dynamic Routing Gateway (DRG) Attachments.

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
|drg_id|`utf8`|
|lifecycle_state|`utf8`|
|display_name|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|drg_route_table_id|`utf8`|
|defined_tags|`json`|
|freeform_tags|`json`|
|route_table_id|`utf8`|
|vcn_id|`utf8`|
|export_drg_route_distribution_id|`utf8`|
|is_cross_tenancy|`bool`|