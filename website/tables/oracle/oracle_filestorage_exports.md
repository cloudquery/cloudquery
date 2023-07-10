# Table: oracle_filestorage_exports

This table shows data for Oracle File Storage Exports.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|export_set_id|`utf8`|
|file_system_id|`utf8`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|path|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|