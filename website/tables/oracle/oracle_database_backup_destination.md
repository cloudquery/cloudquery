# Table: oracle_database_backup_destination

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|compartment_id (PK)|String|
|id (PK)|String|
|display_name|String|
|type|String|
|associated_databases|JSON|
|connection_string|String|
|vpc_users|StringArray|
|local_mount_point_path|String|
|nfs_mount_type|String|
|nfs_server|StringArray|
|nfs_server_export|String|
|lifecycle_state|String|
|time_created|Timestamp|
|lifecycle_details|String|
|freeform_tags|JSON|
|defined_tags|JSON|