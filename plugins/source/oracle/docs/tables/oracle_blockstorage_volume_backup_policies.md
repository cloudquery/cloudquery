# Table: oracle_blockstorage_volume_backup_policies

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
|schedules|JSON|
|time_created|Timestamp|
|destination_region|String|
|defined_tags|JSON|
|freeform_tags|JSON|