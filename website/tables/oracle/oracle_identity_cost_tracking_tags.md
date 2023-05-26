# Table: oracle_identity_cost_tracking_tags

This table shows data for Oracle Identity Cost Tracking Tags.

The composite primary key for this table is (**region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id|`utf8`|
|tag_namespace_id|`utf8`|
|tag_namespace_name|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|is_retired|`bool`|
|time_created|`timestamp[us, tz=UTC]`|
|freeform_tags|`json`|
|defined_tags|`json`|
|lifecycle_state|`utf8`|
|is_cost_tracking|`bool`|