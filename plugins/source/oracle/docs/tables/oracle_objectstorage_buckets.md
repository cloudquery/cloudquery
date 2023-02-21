# Table: oracle_objectstorage_buckets

The composite primary key for this table is (**region**, **compartment_id**, **namespace**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|compartment_id (PK)|String|
|namespace (PK)|String|
|name (PK)|String|
|created_by|String|
|time_created|Timestamp|
|etag|String|
|freeform_tags|JSON|
|defined_tags|JSON|