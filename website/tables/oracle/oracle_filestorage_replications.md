# Table: oracle_filestorage_replications

This table shows data for Oracle File Storage Replications.

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
|lifecycle_state|`utf8`|
|display_name|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|availability_domain|`utf8`|
|replication_interval|`int64`|
|freeform_tags|`json`|
|defined_tags|`json`|
|lifecycle_details|`utf8`|
|recovery_point_time|`timestamp[us, tz=UTC]`|