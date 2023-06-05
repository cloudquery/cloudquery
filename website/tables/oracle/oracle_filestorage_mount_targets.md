# Table: oracle_filestorage_mount_targets

This table shows data for Oracle File Storage Mount Targets.

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
|private_ip_ids|`list<item: utf8, nullable>`|
|subnet_id|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|export_set_id|`utf8`|
|nsg_ids|`list<item: utf8, nullable>`|
|freeform_tags|`json`|
|defined_tags|`json`|