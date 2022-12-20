# Table: oracle_filestorage_mount_targets

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
|lifecycle_state|String|
|private_ip_ids|StringArray|
|subnet_id|String|
|time_created|Timestamp|
|availability_domain|String|
|export_set_id|String|
|nsg_ids|StringArray|
|freeform_tags|JSON|
|defined_tags|JSON|