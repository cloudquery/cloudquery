# Table: oracle_identity_compartments

This table shows data for Oracle Identity Compartments.

The composite primary key for this table is (**region**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
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