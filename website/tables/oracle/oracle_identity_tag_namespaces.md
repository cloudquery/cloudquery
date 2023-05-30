# Table: oracle_identity_tag_namespaces

This table shows data for Oracle Identity Tag Namespaces.

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
|name|`utf8`|
|description|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|
|is_retired|`bool`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|locks|`json`|