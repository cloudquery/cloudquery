# Table: oracle_filestorage_export_sets

This table shows data for Oracle File Storage Export Sets.

The composite primary key for this table is (**region**, **compartment_id**, **availability_domain**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|availability_domain (PK)|`utf8`|
|display_name|`utf8`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|vcn_id|`utf8`|