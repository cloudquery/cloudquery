# Table: oracle_database_key_stores

This table shows data for Oracle Database Key Stores.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|id (PK)|`utf8`|
|display_name|`utf8`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_details|`utf8`|
|associated_databases|`json`|
|freeform_tags|`json`|
|defined_tags|`json`|