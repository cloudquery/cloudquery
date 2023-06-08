# Table: oracle_object_storage_buckets

This table shows data for Oracle Object Storage Buckets.

The composite primary key for this table is (**region**, **compartment_id**, **namespace**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|namespace (PK)|`utf8`|
|name (PK)|`utf8`|
|created_by|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|etag|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|