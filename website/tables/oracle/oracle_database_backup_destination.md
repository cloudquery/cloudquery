# Table: oracle_database_backup_destination

This table shows data for Oracle Database Backup Destination.

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
|display_name|`utf8`|
|type|`utf8`|
|associated_databases|`json`|
|connection_string|`utf8`|
|vpc_users|`list<item: utf8, nullable>`|
|local_mount_point_path|`utf8`|
|nfs_mount_type|`utf8`|
|nfs_server|`list<item: utf8, nullable>`|
|nfs_server_export|`utf8`|
|lifecycle_state|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_details|`utf8`|
|freeform_tags|`json`|
|defined_tags|`json`|