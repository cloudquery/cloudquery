# Table: oracle_filestorage_replications

This table shows data for Oracle File Storage Replications.

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
|lifecycle_state|String|
|display_name|String|
|time_created|Timestamp|
|availability_domain|String|
|replication_interval|Int|
|freeform_tags|JSON|
|defined_tags|JSON|
|lifecycle_details|String|
|recovery_point_time|Timestamp|