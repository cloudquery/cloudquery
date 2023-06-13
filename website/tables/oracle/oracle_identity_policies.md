# Table: oracle_identity_policies

This table shows data for Oracle Identity Policies.

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
|statements|`list<item: utf8, nullable>`|
|description|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_state|`utf8`|
|inactive_status|`int64`|
|version_date|`json`|
|freeform_tags|`json`|
|defined_tags|`json`|