# Table: oracle_identity_compartments

This table shows data for Oracle Identity Compartments.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|compartment_id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_state|`utf8`|
|inactive_status|`int64`|
|is_accessible|`bool`|
|freeform_tags|`json`|
|defined_tags|`json`|