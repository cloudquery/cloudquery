# Table: oracle_identity_dynamic_groups

This table shows data for Oracle Identity Dynamic Groups.

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
|matching_rule|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_state|`utf8`|
|inactive_status|`int64`|
|freeform_tags|`json`|
|defined_tags|`json`|